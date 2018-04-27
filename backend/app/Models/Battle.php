<?php

namespace App\Models;

use App\Helpers\EthereumConverter;
use Illuminate\Database\Eloquent\Model;

class Battle extends Model
{
    const FIELD_TOKEN_ID = 'token_id';
    protected $guarded = ['id'];

    public function group1()
    {
        return $this->hasOne(BattleGroup::class, 'id', 'group_1');
    }

    public function group2()
    {
        return $this->hasOne(BattleGroup::class, 'id', 'group_2');
    }

    public function groupwinner()
    {
        return $this->hasOne(BattleGroup::class, 'id', 'group_winner');
    }

    public function getEntranceFeeAttribute($value)
    {
        return EthereumConverter::convertETHPriceToFloat($value);
    }

    public function getFeeSumAttribute($value)
    {
        return EthereumConverter::convertETHPriceToFloat($value);
    }
}
