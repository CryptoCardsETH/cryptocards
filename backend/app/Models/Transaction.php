<?php

namespace App\Models;

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
}
