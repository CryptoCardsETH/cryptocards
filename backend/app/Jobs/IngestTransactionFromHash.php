<?php

namespace App\Jobs;

use App\Http\Controllers\ContractController;
use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Foundation\Bus\Dispatchable;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Queue\SerializesModels;
use Log;

class IngestTransactionFromHash implements ShouldQueue
{
    use Dispatchable, InteractsWithQueue, Queueable, SerializesModels;

    protected $hash;

    /**
     * Create a new job instance.
     *
     * @return void
     */
    public function __construct($hash)
    {
        $this->hash = $hash;
    }
    public function getHash()
    {
        return $this->hash;
    }

    /**
     * Execute the job.
     *
     * @return void
     */
    public function handle()
    {
        Log::info('IngestTransactionFromHash: '.$this->hash);
        //TODO: call ethereum_proxy, persist information to DB
        //TOOD: re-enqueue itself if txn isn't ready
        //eg: ContractController::processNewBatteGroupEvent("0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1", 10, [1,0,0,0,3]);
    }
}
