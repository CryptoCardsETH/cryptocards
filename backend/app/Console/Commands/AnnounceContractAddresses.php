<?php

namespace App\Console\Commands;

use App\Http\Controllers\ContractController;
use Illuminate\Console\Command;

class AnnounceContractAddresses extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'contract:announce';

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
        ContractController::announceContractAddressesToProxy();
        $this->info('called announce function:');
    }
}
