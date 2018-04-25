<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use RpcServer\CoreContractAddress;

class Contract extends Model
{
    const FIELD_NAME = 'name';
    const FIELD_ADDRESS = 'address';
    const FIELD_TRANSACTION_HASH = 'transaction_hash';
    const CONTRACT_NAME_CORE = 'CryptoCardsCore';
    protected $fillable = ['name'];

    public static function updateAddress($name, $address, $transactionHash)
    {
        $c = self::firstOrNew([self::FIELD_NAME => $name]);
        $c[self::FIELD_ADDRESS] = $address;
        $c[self::FIELD_TRANSACTION_HASH] = $transactionHash;
        $c->save();
    }

    /*
     * Package the Model isntance into an Rpc Message
     */
    public static function getRpcCoreContractAddressMessage($name = self::CONTRACT_NAME_CORE)
    {
        $core = self::where('name', $name)->first();
        $ca = new CoreContractAddress();
        $ca->setAddress($core->address);
        return $ca;
    }
}
