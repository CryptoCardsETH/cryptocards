<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class BattleGroupCard extends Model
{
    public function card()
    {
        return $this->hasOne(Card::class, 'id');
    }
}
