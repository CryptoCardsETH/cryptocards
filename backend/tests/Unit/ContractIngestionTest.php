<?php

namespace Tests\Unit;

use App\Http\Controllers\ContractController;
use App\Models\BattleGroup;
use Illuminate\Support\Facades\Queue;
use Tests\TestCase;

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
}
