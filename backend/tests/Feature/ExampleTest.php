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
        $response = $this->get('/v1');

        $response->assertStatus(200);
    }

    public function testAuthAndGetMe()
    {
        $faker = \Faker\Factory::create();

        $address = '0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1';
        $signedMessage = '0x856359f869e2c5cd56670c0a6af4cccf82bb2299b73116077a8f88254006f75a6e8530e8b19e25c4bab99f7a06536166a7312cb4d10342cea252458ef39375551b';
        $response = $this->json('POST', '/v1/auth', ['address'=>$address, 'signed'=>$signedMessage, 'plaintext'=>'CryptoCards'])
                  ->assertJsonStructure(['data'=>['token']])->assertStatus(200);
                  
        $token = json_decode($response->getContent(), true)['data']['token'];

        $response = $this->authenticatedJSON('GET', '/v1/me', [], $token);
        $me = json_decode($response->getContent(), true)['data'];
        $this->assertEquals($me['address'], $address);
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
