<?php

namespace Tests\Feature;

use App\Http\Controllers\Controller;
use App\Mail\WelcomeEmail;
use App\Models\Card;
use App\Models\User;
use Illuminate\Support\Facades\Mail;
use Tests\TestCase;

class UserTest extends TestCase
{
    public function testGetUpdateUser()
    {
        //get /me on a user
        $user = factory(User::class)->create();
        $response = $this->authenticatedJSON('GET', '/v1/me', [], $user->getToken());
        $response
            ->assertStatus(200)
            ->assertJson(['success' => true]);
        $data = $response->json()['data'];
        $address = $user->address;
        $this->assertEquals($address, $data['address']);

        //change nickname, ensure it persists
        $faker = \Faker\Factory::create();
        $newNickName = $faker->firstName;
        $data[User::FIELD_NICKNAME] = $newNickName;
        $response = $this->authenticatedJSON('PUT', '/v1/me', $data, $user->getToken());

        $response
            ->assertStatus(200)
            ->assertJson(['success' => true]);

        $this->assertDatabaseHas('users', [
            'id'                => $user->id,
            User::FIELD_NICKNAME=> $newNickName,
        ]);

        //test that changing email triggers welcome email
        Mail::fake();

        $data[User::FIELD_EMAIL] = $faker->unique()->email;
        $reponse = $this->authenticatedJSON('PUT', '/v1/me', $data, $user->getToken());
        $response
            ->assertStatus(200)
            ->assertJson(['success' => true]);

        Mail::assertSent(WelcomeEmail::class, function ($mail) use ($user) {
            return $mail->user->id === $user->id;
        });
    }

    public function testGetUpdateUserDuplicateNickname()
    {
        $faker = \Faker\Factory::create();
        $conflictingNickname = $faker->unique()->lastName;

        $conflictingUser = factory(User::class)->create();
        $conflictingUser->nickname = $conflictingNickname;
        $conflictingUser->save();

        //get /me on a user
        $user = factory(User::class)->create();
        $response = $this->authenticatedJSON('GET', '/v1/me', [], $user->getToken());
        $response
            ->assertStatus(200)
            ->assertJson(['success' => true]);
        $data = $response->json()['data'];
        $address = $user->address;
        $originalNickname = $user->nickname;
        $this->assertEquals($address, $data['address']);

        //change nickname, ensure it doesn't persist
        $data[User::FIELD_NICKNAME] = $conflictingNickname;
        $response = $this->authenticatedJSON('PUT', '/v1/me', $data, $user->getToken());
        $response
            ->assertStatus(500)
            ->assertJson(['success' => false, 'response_id' => Controller::RESPONSE_MESSAGE_ERROR_DUPLICATE[0]]);

        $this->assertDatabaseHas('users', [
            'id'                => $user->id,
            User::FIELD_NICKNAME=> $originalNickname,
        ]);
    }

    public function testGetUpdateUserDuplicateEmail()
    {
        $faker = \Faker\Factory::create();
        $confictingEmail = $faker->unique()->email;

        $conflictingUser = factory(User::class)->create();
        $conflictingUser->email = $confictingEmail;
        $conflictingUser->save();

        //get /me on a user
        $user = factory(User::class)->create();
        $response = $this->authenticatedJSON('GET', '/v1/me', [], $user->getToken());
        $response
            ->assertStatus(200)
            ->assertJson(['success' => true]);
        $data = $response->json()['data'];
        $originalEmail = $user->email;

        //change email, ensure it doesn't persist
        $data[User::FIELD_EMAIL] = $confictingEmail;
        $response = $this->authenticatedJSON('PUT', '/v1/me', $data, $user->getToken());
        $response
            ->assertStatus(500)
            ->assertJson(['success' => false, 'response_id' => Controller::RESPONSE_MESSAGE_ERROR_DUPLICATE[0]]);

        $this->assertDatabaseHas('users', [
            'id'                => $user->id,
            User::FIELD_EMAIL   => $originalEmail,
        ]);
    }

    public function testGetUpdateCardHiddenFlag()
    {
        $card = factory(Card::class)->create();
        $user = factory(User::class)->create();
        $card->user_id = $user->id;
        $card->save();

        //ensure initial state is visible
        $this->assertDatabaseHas('cards', [
            'id'                      => $card->id,
            CARD::FIELD_HIDDEN_TOGGLE => false,
        ]);

        //get the current state of the card
        $response = $this->authenticatedJSON('GET', '/v1/cards/'.$card->id, [], $user->getToken());
        $response->assertStatus(200)->assertJson(['success' => true]);
        $data = $response->json()['data'];

        //update the hidden flag to true
        $data[Card::FIELD_HIDDEN_TOGGLE] = true;
        $response = $this->authenticatedJSON('PUT', '/v1/cards/'.$card->id, $data, $user->getToken());
        $response->assertStatus(200)->assertJson(['success' => true]);

        //ensure update was persisted
        $this->assertDatabaseHas('cards', [
            'id'                      => $card->id,
            CARD::FIELD_HIDDEN_TOGGLE => true,
        ]);

        //now let's try to have an unauthorized user change it back to visible
        $userNotOwner = factory(User::class)->create();
        $data[Card::FIELD_HIDDEN_TOGGLE] = false;
        $response = $this->authenticatedJSON('PUT', '/v1/cards/'.$card->id, $data, $userNotOwner->getToken());
        $response->assertStatus(401)->assertJson(['success' => false]);

        //ensure update was persisted
        $this->assertDatabaseHas('cards', [
            'id'                      => $card->id,
            CARD::FIELD_HIDDEN_TOGGLE => true,
        ]);
    }

    public function testGetUserById()
    {
        //make a user with 2 cards, one of which is hidden.
        $user = factory(User::class)->create();

        $card1 = factory(Card::class)->create();
        $card1->user_id = $user->id;
        $card1->save();

        $card2 = factory(Card::class)->create();
        $card2->user_id = $user->id;
        $card2->hidden = true;
        $card2->save();

        //getting me as myself should show hidden cards (so 2 total)
        $response = $this->authenticatedJSON('GET', '/v1/users/'.$user->id, [], $user->getToken());
        $response->assertStatus(200)->assertJson(['success' => true]);
        $data = $response->json()['data'];
        self::assertEquals(2, count($data['cards']));

        //another user getting me should only see my 1 public card
        $user2 = factory(User::class)->create();
        $response = $this->authenticatedJSON('GET', '/v1/users/'.$user->id, [], $user2->getToken());
        $response->assertStatus(200)->assertJson(['success' => true]);
        $data = $response->json()['data'];
        self::assertEquals(1, count($data['cards']));

        //an unauthorized user getting me should only see my 1 public card
        $response = $this->json('GET', '/v1/users/'.$user->id);
        $response->assertStatus(200)->assertJson(['success' => true]);
        $data = $response->json()['data'];
        self::assertEquals(1, count($data['cards']));
    }
}
