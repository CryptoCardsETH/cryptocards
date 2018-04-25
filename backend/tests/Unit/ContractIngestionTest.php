<?php

namespace Tests\Unit;

use App\Http\Controllers\ContractController;
use App\Models\BattleGroup;
use App\Models\Contract;
use Tests\TestCase;

class ContractIngestionTest extends TestCase
{
    const CONTRACT_NAME_CORE = 'CryptoCardsCoreTest';

    public function testIngestBattleGroupEvent()
    {
        $nextBattleGroupId = BattleGroup::max('token_id') + 1;
        $ownerAddress = 'aaaa';
        ContractController::processNewBattleGroupEvent($ownerAddress, $nextBattleGroupId, [1, 2, 3]);

        $this->assertDatabaseHas('battle_groups', [
            BattleGroup::FIELD_TOKEN_ID  => $nextBattleGroupId,
        ]);
    }

    public function testIngestCoreContractAddress()
    {
        putenv('APP_DEBUG=true');
        $faker = \Faker\Factory::create();
        //test ingesting the adddress (truffle hits this endpoint)
        $address = '0xfakex'.$faker->sha1;
        $response = $this->json('PUT', '/v1/contracts/ingest', [
            self::CONTRACT_NAME_CORE => [
                CONTRACT::FIELD_ADDRESS => $address,
                'transactionHash'       => 'asdf',
                ],

        ])->assertStatus(200);

        return;

        //ensure the API returns the proper core address
        $response = $this->get('/v1/contracts')->assertStatus(200);
        $data = json_decode($response->getContent(), true)['data'];
        $this->assertEquals($data[self::CONTRACT_NAME_CORE]['address'], $address);

        //test that this address shows up in the protobuf msg
        $rpcMessage = Contract::getRpcCoreContractAddressMessage(self::CONTRACT_NAME_CORE);
        $this->assertEquals($rpcMessage->getAddress(), $address);
    }

    public function testIngestCoreContractAddressAuth()
    {
        putenv('APP_DEBUG=false');
        $response = $this->put('/v1/contracts/ingest')->assertStatus(401);
        putenv('APP_DEBUG=true');
    }
}
