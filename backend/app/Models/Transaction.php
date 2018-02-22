<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Transaction extends Model
{
    const FIELD_TRANSACTION_TIME = 'transaction_time';
    
    public function attributes()
    {
        return $this->belongsToMany('App\Models\Attribute')->withPivot('value');
    }
    
    public function cards()
    {
        return $this->hasOne(Card::class, 'id', 'card_id');
    }
}
