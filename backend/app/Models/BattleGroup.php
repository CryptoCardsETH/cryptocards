<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class BattleGroup extends Model
{
    const FIELD_TOKEN_ID = 'token_id';
    protected $guarded = ['id'];

    public function group_cards()
    {
        return $this->hasMany(BattleGroupCard::class, 'group_id');
    }

    public function user()
    {
        return $this->hasOne(User::class, 'id', 'user_id');
    }

    public static function getByTokenId($tokenId)
    {
        return self::firstOrCreate(['token_id' => $tokenId]);
    }
    public static function getNextTokenId()
    {
        return self::max('token_id') + 1;
    }
}
