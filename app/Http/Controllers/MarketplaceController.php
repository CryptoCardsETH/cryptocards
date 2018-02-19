<?php
/*
	Marketplace Controller - github.com/CryptoCardsETH/backend
	File created by Harris Christiansen (HarrisChristiansen.com)
*/

namespace App\Http\Controllers;

use App\Models\Card;
use Illuminate\Http\Request;

class MarketplaceController extends Controller {

    public function __construct()
    {
        $this->middleware('jwt.auth', ['except' => ['getAllCards']]);
    }

    /**
     * Gets all the cards.
     * @return mixed cards
     */
    public function getAllCards() {
        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, Card::with('attributes')->get());
    }
	
}
