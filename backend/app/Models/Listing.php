<?php

namespace App\Models;

use App\Helpers\EthereumConverter;
use Illuminate\Database\Eloquent\Model;

class Listing extends Model
{
    public $timestamps = false;
    
    /**
     * Defines the relationship between Listings and Cards.
     */
    public function cards()
    {
        return $this->hasOne(Card::class, 'id', 'card_id');
    }

    public function getPriceAttribute($value)
    {
        //price is stores as a bigint in the db, need to make it a properly formatted float.
        return EthereumConverter::convertETHPriceToFloat($value);
    }
}
