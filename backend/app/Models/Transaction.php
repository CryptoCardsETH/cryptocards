<?php

namespace App\Models;

use App\Helpers\EthereumConverter;
use Illuminate\Database\Eloquent\Model;

class Transaction extends Model
{
    const FIELD_TRANSACTION_TIME = 'transaction_time';

    public function transactionCard()
    {
        return $this->hasOne(Card::class, 'id', 'card_id');
    }

    public function transactionUser()
    {
        return $this->hasOne('App\Models\User');
    }

    public function getPriceAttribute($value)
    {
        //price is stores as a bigint in the db, need to make it a properly formatted float.
        return EthereumConverter::convertETHPriceToInt($value);
    }
}
