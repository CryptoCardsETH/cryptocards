<?php

namespace App\Models;

use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;
use Illuminate\Database\Eloquent\Model;
use Tymon\JWTAuth\Contracts\JWTSubject;

class User extends Authenticatable implements JWTSubject
{
    use Notifiable;
    const FIELD_EMAIL = 'email';
    const FIELD_NICKNAME = 'nickname';
    protected $guarded = ['id'];

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

    /**
     * Return boolean, whether or not user is following user_id
     * 
     * @return boolean
     */
    public function isFollowing($user_id) {
        return Follow::where([
            ['user_id', '=', $user_id],
            ['follower_id', '=', $this->id],
        ])->exists();
    }
}
