<?php

namespace App\Http\Controllers;

use App\Jobs\IngestTransactionFromHash;
use App\Models\BattleGroup;
use App\Models\BattleGroupCard;
use App\Models\Card;
use RpcServer\ContractAddresses;
use App\Helpers\EthereumHelper;
use App\Models\Contract;
use App\Models\User;
use Illuminate\Support\Facades\Request;
use Log;

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

        self::announceContractAddressesToProxy();

        return response()->build(self::RESPONSE_MESSAGE_SUCCESS);
    }
    /*
     * Announce contract addresses to the RPC server
     */
    public static function announceContractAddressesToProxy() {
        $addresses = [];
        foreach(Contract::all() as $c) {
            array_push($addresses,$c->getRpcContractAddressMessage());
        }

        $msg = new ContractAddresses();
        $msg->setItems($addresses);

        $client = EthereumHelper::getRpcClient();

        list($reply, $status) = $client->AnnounceContractAddresses($msg)->wait();
        Log::info('announce received:'.$reply->getMessage());

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

        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, $txHash);
    }

    /*
     * Save the results of a BattleGroup contract Event into the DB.
     */
    public static function processNewBattleGroupEvent($ownerAddress, $battleGroupId, $cards)
    {
        $user = User::getByAddress($ownerAddress);
        $bgData = ['user_id'=>$user->id, 'token_id'=> $battleGroupId];
        $bg = BattleGroup::firstOrCreate($bgData);
        foreach ($cards as $c) {
            $card = Card::getByTokenId($c);
            BattleGroupCard::firstOrCreate(['group_id'=>$bg->id, 'card_id'=>$card->id]);
        }
        if($bg->wasRecentlyCreated) {
            Log::info("ingested new BattleGroup: user_id",["data"=>$bgData, "cardIds"=>$cards]);
        }
    }
}
