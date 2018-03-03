<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Friend extends Model
{
    //
    public static function areFriends($user_id1, $user_id2)
    {
        return Friend::where([
            ['user_id', '=', $user_id1],
            ['friend_id', '=', $user_id2]
        ])->orWhere([
            ['user_id', '=', $user_id2],
            ['friend_id', '=', $user_id1]
        ])->exists();
    }
}
