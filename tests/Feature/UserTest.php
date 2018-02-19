<?php

namespace Tests\Feature;

use App\Mail\WelcomeEmail;
use App\Models\User;
use Illuminate\Support\Facades\Mail;
use Tests\TestCase;
use Illuminate\Foundation\Testing\RefreshDatabase;

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
            'id'=>$user->id,
            User::FIELD_NICKNAME=>$newNickName,
        ]);


        //test that changing email triggers welcome email
        Mail::fake();

        $data[User::FIELD_EMAIL] = $faker->unique()->email;
        $this->authenticatedJSON('PUT', '/v1/me', $data, $user->getToken());

        Mail::assertSent(WelcomeEmail::class, function ($mail) use ($user) {
            return $mail->user->id === $user->id;
        });

    }
}
