<?php

namespace App\Console\Commands;

use App\Helpers\EthereumHelper;
use App\Http\Controllers\ContractController;
use App\Models\Contract;
use Illuminate\Console\Command;
use RpcServer\CardInfoRequest;

class IngestCardsFromBlockchain extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'contract:ingestcards';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Ingest Cards from contract events';

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

        $msg = new CardInfoRequest();
        $msg->setContract($ca);

        //ask ethereum_proxy for the CardInfo, providing it with the CoreContract contract address.

        list($reply, $status) = $client->RequestCardInfo($msg)->wait();
        foreach ($reply->getItems() as $item) {
            ContractController::processCardInfoRpc($item);
        }
    }
}
