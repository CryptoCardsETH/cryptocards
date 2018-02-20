<?php

namespace Tests\Feature;

use App\Http\Controllers\Controller;
use App\Models\Card;
use Tests\TestCase;

class ExampleTest extends TestCase
{
    /**
     * A basic test example.
     *
     * @return void
     */
    public function testBasicTest()
    {
        $response = $this->get('/');

        $response->assertStatus(200);
    }

    public function testAuthAndGetMe()
    {
        $faker = \Faker\Factory::create();

        $address = '0xfakex'.$faker->sha1;
        $response = $this->json('POST', '/v1/auth', ['address'=>$address, 'signed'=>'aa', 'plaintext'=>'CryptoCards'])
            ->assertJsonStructure(['data'=>['token']]);
        $token = json_decode($response->getContent(), true)['data']['token'];

        $response = $this->authenticatedJSON('GET', '/v1/me', [], $token);
        $me = json_decode($response->getContent(), true)['data'];
        $this->assertEquals($me['address'], $address);
        $this->assertEquals($me['email'], null);
    }

    public function testGetMeUnAuthed()
    {
        $response = $this->json('GET', '/v1/me');
        $response_decoded = json_decode($response->getContent(), true);
        $this->assertEquals($response_decoded['data'], null);
        $this->assertEquals($response_decoded['message'], Controller::RESPONSE_MESSAGE_ERROR_JWT_ERROR[2]);
    }

    public function testGetCards()
    {
        $response = $this->json('GET', '/v1/cards');
        $data = json_decode($response->getContent(), true)['data'];
        $this->assertEquals(count($data), Card::all()->count());
    }

    public function testGetCardDetails()
    {
        $card = factory(Card::class)->create();
        $response = $this->json('GET', '/v1/cards/'.$card->id);
        $data = json_decode($response->getContent(), true)['data'];
        $this->assertEquals($card->name, $data['name']);
    }
}
