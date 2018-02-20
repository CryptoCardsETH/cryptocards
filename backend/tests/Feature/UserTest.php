<?php

namespace Tests\Feature;

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
        $this->authenticatedJSON('PUT', '/v1/me', $data, $user->getToken());

        Mail::assertSent(WelcomeEmail::class, function ($mail) use ($user) {
            return $mail->user->id === $user->id;
        });
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
}
