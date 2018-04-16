<?php

namespace App\Http\Controllers;

use App\Models\Contract;
use Illuminate\Support\Facades\Request;
use Log;
use App\Jobs\IngestTransactionFromHash;
use App\Models\User;
use App\Models\BattleGroup;
use App\Models\BattleGroupCard;
use App\Models\Card;

class ContractController extends Controller
{
    public function contractAddressIngest()
    {
        if (env('APP_DEBUG') != 'true') {
            //this feature is only really for the development workflow, so no auth needed.
            return response()->build(self::RESPONSE_MESSAGE_ERROR_UNAUTHORIZED);
        }

        $data = json_decode(Request::getContent(), true);
        Log::info($data);

        foreach ($data as $contractName => $data) {
            $address = $data['address'];
            Log::info($contractName.' @ '.$address);
            Contract::updateAddress($contractName, $address, $data['transactionHash']);
        }

        return response()->build(self::RESPONSE_MESSAGE_SUCCESS);
    }

    public function getContractAddresses()
    {
        $c = Contract::all()->keyBy(Contract::FIELD_NAME);

        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, $c);
    }
    public function startWatchingTransaction()
    {
        $data = json_decode(Request::getContent(), true);
        $txHash = $data['txHash'];
        IngestTransactionFromHash::dispatch($txHash);
        return response()->build(self::RESPONSE_MESSAGE_SUCCESS);

    }
    public static function processNewBattleGroupEvent($ownerAddress, $battleGroupId, $cards) {
        $user = User::getByAddress($ownerAddress);
        $bg = BattleGroup::create(['user_id'=>$user->id, 'token_id'=> $battleGroupId]);
        foreach($cards as $c) {
            $card = Card::getByTokenId($c);
            BattleGroupCard::create(['group_id'=>$bg->id, 'card_id'=>$card->id]);
        }
    }
}
