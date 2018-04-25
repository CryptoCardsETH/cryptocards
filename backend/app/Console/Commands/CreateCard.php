<?php

namespace App\Console\Commands;

use App\Helpers\EthereumHelper;
use App\Models\Contract;
use Illuminate\Console\Command;
use RpcServer\CreateCardRequest;

class CreateCard extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    //default is public key of ganache-cli user #1
    protected $signature = 'contract:createcard {ownerAddress=0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Command description';

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
        $ownerAddress = $this->argument('ownerAddress');
        $this->info($ownerAddress);

        $ca = Contract::getRpcCoreContractAddressMessage();

        $client = EthereumHelper::getRpcClient();

        //make a card, for testing
        $makeCardMsg = new CreateCardRequest();
        $makeCardMsg->setCoreAddress($ca);
        $makeCardMsg->setOwnerAddress($ownerAddress);
        list($reply, $status) = $client->CreateCard($makeCardMsg)->wait();
        $this->info('created Card with txn: '.$reply->getMessage().' for user '.$ownerAddress);
    }
}
