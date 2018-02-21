<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use RpcServer\BlankRequest;
use RpcServer\GreeterClient;

class ProtoTest extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'test:proto';

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
        $msg = new BlankRequest();
        $msg->setName('test');
        $this->info('sending:'.$msg->getName());

        $client = new GreeterClient(env('RPC_SERVER_ADDRESS'), [
            'credentials' => \Grpc\ChannelCredentials::createInsecure(),
        ]);
        list($reply, $status) = $client->GetBlank($msg)->wait();
        $this->info('received:'.$reply->getMessage());
    }
}
