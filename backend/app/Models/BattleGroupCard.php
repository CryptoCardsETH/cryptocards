<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class BattleGroupCard extends Model
{
    protected $guarded = ['id'];

    public function card()
    {
        return $this->belongsTo(Card::class);
    }
}
