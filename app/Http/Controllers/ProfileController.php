<?php

namespace App\Http\Controllers;

use App\Models\Card;

class ProfileController extends Controller {

    public function __construct()
    {
        $this->middleware('jwt.auth');
    }

    /**
     * Gets the info about the current user.
     * @return mixed user
     */
    public function me()
    {
        return response()->build(self::RESPONSE_MESSAGE_SUCCESS,auth()->user());
    }

    /**
     * Gets all the cards that the user is an owner of.
     * @return mixed cards
     */
    public function getMyCards() {
        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, Card::where("user_id",auth()->user()->id)->get());
    }

	
}
