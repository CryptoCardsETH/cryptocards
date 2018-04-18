<?php

namespace App\Helpers;

use RpcServer\ECRecoverRequest;
use RpcServer\GreeterClient;

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
        if (env('ENABLE_ETHEREUM_PROXY') != 'true') {
            //hack so that this doesn't break other people who aren't running in docker yet
            return true;
        }
        $client = self::getRpcClient();

        $msg = new ECRecoverRequest();
        $msg->setAddress($signerAddress);
        $msg->setSigned($signedMessage);
        $msg->setPlaintext($plainTextMessage);

        list($reply, $status) = $client->PerformECRecover($msg)->wait();

        //to dat ECrecover
        return $reply;
    }
}
