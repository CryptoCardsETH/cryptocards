<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Card extends Model
{
    public function attributes()
    {
        return $this->hasMany('App\Models\CardAttribute');
    }
}
