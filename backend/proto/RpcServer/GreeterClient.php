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

    /**
     * @param \RpcServer\CreateCardRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function CreateCard(\RpcServer\CreateCardRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/rpcServer.Greeter/CreateCard',
        $argument,
        ['\RpcServer\BlankReply', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \RpcServer\CardsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetCardsByOwner(\RpcServer\CardsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/rpcServer.Greeter/GetCardsByOwner',
        $argument,
        ['\RpcServer\CardsReply', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \RpcServer\BattleGroupInfoRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function RequestBattleGroupInfo(\RpcServer\BattleGroupInfoRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/rpcServer.Greeter/RequestBattleGroupInfo',
        $argument,
        ['\RpcServer\BattleGroupInfoReply', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \RpcServer\CardInfoRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function RequestCardInfo(\RpcServer\CardInfoRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/rpcServer.Greeter/RequestCardInfo',
        $argument,
        ['\RpcServer\CardInfoReply', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \RpcServer\BattleInfoRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function RequestBattleInfo(\RpcServer\BattleInfoRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/rpcServer.Greeter/RequestBattleInfo',
        $argument,
        ['\RpcServer\BattleInfoReply', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \RpcServer\CoreContractAddress $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function TestThings(\RpcServer\CoreContractAddress $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/rpcServer.Greeter/TestThings',
        $argument,
        ['\RpcServer\BlankReply', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \RpcServer\ECRecoverRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function PerformECRecover(\RpcServer\ECRecoverRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/rpcServer.Greeter/PerformECRecover',
        $argument,
        ['\RpcServer\ECRecoverReply', 'decode'],
        $metadata, $options);
    }

}
