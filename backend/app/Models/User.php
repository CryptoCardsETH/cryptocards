<?php

namespace App\Models;

use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;
use Tymon\JWTAuth\Contracts\JWTSubject;

class User extends Authenticatable implements JWTSubject
{
    use Notifiable;
    const FIELD_EMAIL = 'email';
    const FIELD_NICKNAME = 'nickname';
    protected $guarded = ['id'];

    public function cards()
    {
        return $this->hasMany(Card::class);
    }

    /**
     * Get the identifier that will be stored in the subject claim of the JWT.
     *
     * @return mixed
     */
    public function getJWTIdentifier()
    {
        return $this->getKey();
    }

    /**
     * Return a key value array, containing any custom claims to be added to the JWT.
     *
     * @return array
     */
    public function getJWTCustomClaims()
    {
        return ['address'=>$this->address];
    }

    public function getToken()
    {
        return auth()->login($this);
    }

    public function followers()
    {
        return $this->belongsToMany('App\Models\User', 'follows', 'user_id', 'follower_id')->withTimestamps();
    }

    public function following()
    {
        return $this->belongsToMany('App\Models\User', 'follows', 'follower_id', 'user_id')->withTimestamps();
    }

    /**
     * Follow user_id.
     *
     * @throws ModelNotFoundException if $user_id could not be found.
     *                                http 404 page is returned if exception is not caught.
     *
     * @return bool True - if user is successfully following user_id, False - if user is already following user_id.
     */
    public function follow($user_id)
    {
        if (!$this->following->contains($user_id)) {
            $user = self::findOrFail($user_id);
            $user->followers()->attach($this->id);

            return true;
        } else {
            return false;
        }
    }

    /**
     * Gets all battles a user has participated in, as per their BattleGroups.
     */
    public function getAllBattles()
    {
        $user_id = $this->id;

        return Battle::whereHas('group1', function ($query) use ($user_id) {
            $query->where('user_id', $user_id);
        })->orWhereHas('group2', function ($query) use ($user_id) {
            $query->where('user_id', $user_id);
        })->with(
            'group1.user',
            'group1.group_cards.card.user',
            'group2.user',
            'group2.group_cards.card.user',
            'groupwinner.user'
        )->get();
    }

    /**
     * Gets all BattleGroups a user has.
     */
    public function getAllBattleGroups()
    {
        return BattleGroup::with('group_cards.card.user', 'user')->where('user_id', $this->id)->get();
    }

    public static function getByAddress($address)
    {
        return self::firstOrCreate(['address'=>$address]);
    }
}
