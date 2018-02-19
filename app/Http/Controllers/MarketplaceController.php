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
        $this->middleware('jwt.auth', ['except' => ['getAllCards','getCardDetail']]);
    }

    /**
     * Gets all the cards.
     * @return mixed cards
     */
    public function getAllCards() {
        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, Card::with('attributes','user')->get());
    }

    public function getCardDetail($card_id) {
        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, Card::with('attributes','user')->find($card_id));
    }
	
}
