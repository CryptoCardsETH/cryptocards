<?php

namespace Tests\Unit;

use App\Http\Controllers\ContractController;
use App\Models\BattleGroup;
use Tests\TestCase;
use App\Jobs\IngestTransactionFromHash;
use Illuminate\Support\Facades\Queue;
class ContractIngestionTest extends TestCase
{
    public function testIngestBattleGroupEvent()
    {
        $nextBattleGroupId = BattleGroup::max('token_id') + 1;
        $ownerAddress = 'aaaa';
        ContractController::processNewBattleGroupEvent($ownerAddress, $nextBattleGroupId, [1, 2, 3]);

        $this->assertDatabaseHas('battle_groups', [
            BattleGroup::FIELD_TOKEN_ID  => $nextBattleGroupId,
        ]);
    }
    public function testStartWatchingTransaction() {
        Queue::fake();

        $hash = "0x123";

        $response = $this->json('PUT', 'v1/contracts/watchTransaction', ['txHash'=>$hash]);
        $response
            ->assertStatus(200)
            ->assertJson(['success' => true]);
        Queue::assertPushed(IngestTransactionFromHash::class, function ($job) use ($hash) {
            return $job->getHash() === $hash;
        });
    }
}
