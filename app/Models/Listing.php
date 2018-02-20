<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Listing extends Model
{
    /**
     * Defines the relationship between Listings and Cards.
     */
    public function cards()
    {
        return $this->hasOne(Card::class, 'id', 'card_id');
    }
}
