<?php

namespace App\Console\Commands;

use App\Helpers\EthereumHelper;
use App\Http\Controllers\ContractController;
use App\Models\Contract;
use Illuminate\Console\Command;
use RpcServer\BattleInfoRequest;

class IngestBattlesFromBlockchain extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'contract:ingestbattles';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Ingest Battles from contract events';

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
        $ca = Contract::getRpcCoreContractAddressMessage();
        $client = EthereumHelper::getRpcClient();

        $msg = new BattleInfoRequest();
        $msg->setCoreAddress($ca);

        list($reply, $status) = $client->RequestBattleInfo($msg)->wait();
        foreach ($reply->getBattles() as $item) {
            ContractController::processBattleInfoRpc($item);
        }
    }
}
