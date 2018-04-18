<?php

namespace App\Helpers;

use RpcServer\GreeterClient;
use RpcServer\ECRecoverRequest;
use RpcServer\ECRecoverReply;
use Log;

class EthereumHelper
{
    public static function getRpcClient()
    {
        return new GreeterClient(env('RPC_SERVER_ADDRESS'), [
            'credentials' => \Grpc\ChannelCredentials::createInsecure(),
        ]);
    }
    /**
     * TODO:
     * This needs to verify that the message was properly signed.
     * Just like web3js does: http://web3js.readthedocs.io/en/1.0/web3-eth-accounts.html#recover.
     *
     * @param $signerAddress
     * @param $signedMessage
     * @param $plainTextMessage
     *
     * @return bool
     */
    public static function didAddressReallySignMessage($signerAddress, $signedMessage, $plainTextMessage)
    {
        $client = self::getRpcClient();



        $msg = new ECRecoverRequest();
        $msg->setAddress($signerAddress);
        $msg->setSigned($signedMessage);
        $msg->setPlaintext($plainTextMessage);

        //ask ethereum_proxy for the BattleGroupInfo, providing it with the BattleGroup contract address.

        list($reply, $status) = $client->PerformECRecover($msg)->wait();

        //to dat ECrecover
        return $reply;

    }
    
}
