<?php

namespace App\Console\Commands;

use App\Helpers\EthereumHelper;
use App\Http\Controllers\ContractController;
use App\Models\Contract;
use Illuminate\Console\Command;
use RpcServer\BattleGroupInfoRequest;

class IngestBattleGroupsFromBlockchain extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'contract:ingestbattlegroups';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Ingest BattleGroups from contract events';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $contract = Contract::where('name', 'BattleGroups')->first();
        $ca = $contract->getRpcContractAddressMessage();

        $client = EthereumHelper::getRpcClient();

        $msg = new BattleGroupInfoRequest();
        $msg->setContract($ca);

        //ask ethereum_proxy for the BattleGroupInfo, providing it with the BattleGroup contract address.

        list($reply, $status) = $client->RequestBattleGroupInfo($msg)->wait();
        foreach ($reply->getItems() as $item) {
            $cardIds = [];
            foreach ($item->getCards() as $id) {
                array_push($cardIds, $id);
            }
            ContractController::processNewBattleGroupEvent($item->getOwnerAddress(), $item->getId(), $cardIds);
        }
    }
}
