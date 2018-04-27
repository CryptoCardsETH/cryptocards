<?php

namespace App\Console\Commands;

use App\Helpers\EthereumHelper;
use App\Models\Contract;
use Illuminate\Console\Command;

class BlockchainPlayground extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'contract:playground';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'BlockchainPlayground';

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
        list($reply, $status) = $client->TestThings($ca)->wait();
    }
}
