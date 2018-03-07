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
     * Return boolean, whether or not user is following user_id.
     *
     * @return bool
     */
    public function isFollowing($user_id)
    {
        return Follow::where([
            ['user_id', '=', $user_id],
            ['follower_id', '=', $this->id],
        ])->exists();
    }

    /**
     * Follow user_id.
     *
     * @return bool True - if user is successfully following user_id, False - if user is already following user_id.
     * @throws ModelNotFoundException if $user_id could not be found.
     * http 404 page is returned if exception is not caught.
     */
    public function follow($user_id)
    {
        if (!$this->following->contains($user_id)) {
            $user = User::findOrFail($user_id);
            $user->followers()->attach($this->id);
            return true;
        } else {
            return false;
        }
    }
}
