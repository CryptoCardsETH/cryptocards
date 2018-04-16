<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Card extends Model
{
    const FIELD_HIDDEN_TOGGLE = 'hidden';
    protected $appends = ['imageURL'];
    protected $guarded = ['id'];

    public function attributes()
    {
        return $this->belongsToMany('App\Models\Attribute')->withPivot('value');
    }

    public function user()
    {
        return $this->BelongsTo('App\Models\User');
    }

    public function getImageURLAttribute()
    {
        //TODO: real url
        return 'http://via.placeholder.com/350?text=card+id'.$this->id;
    }

    public function isUserOwner(User $user)
    {
        return $user->id === $this->user_id;
    }

    public function transactions()
    {
        return $this->hasMany('App\Models\Transaction');
    }

    public static function getByTokenId($tokenId)
    {
        $card = self::firstOrNew(['token_id'=>$tokenId]);
        if (!$card->exists) {
            $card->name = 'temp name';
            $card->save();
        }

        return $card;
    }
}
