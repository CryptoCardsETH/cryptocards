<?php
// GENERATED CODE -- DO NOT EDIT!

namespace RpcServer;

/**
 */
class GreeterClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \RpcServer\BlankRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetBlank(\RpcServer\BlankRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/rpcServer.Greeter/GetBlank',
        $argument,
        ['\RpcServer\BlankReply', 'decode'],
        $metadata, $options);
    }

}
