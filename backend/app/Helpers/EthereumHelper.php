<?php

namespace App\Helpers;
use RpcServer\GreeterClient;
class EthereumHelper 
{
    public static function getRpcClient() {
        return new GreeterClient(env('RPC_SERVER_ADDRESS'), [
            'credentials' => \Grpc\ChannelCredentials::createInsecure(),
        ]); 
    }
}
