<?php
/*
	Marketplace Controller - github.com/CryptoCardsETH/backend
	File created by Harris Christiansen (HarrisChristiansen.com)
*/

namespace App\Http\Controllers;

use App\Models\Card;
use App\Models\Listing;
use Illuminate\Http\Request;

class MarketplaceController extends Controller {

    public function __construct()
    {
        $this->middleware('jwt.auth', ['except' => ['getAllCards', 'getAllListings']]);
    }

    /**
     * Gets all the cards.
     * @return mixed cards
     */
    public function getAllCards() {
        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, Card::get());
    }

    public function getAllListings() {
        $listings = Listing::join('cards','cards.id','=','listings.card_id')->select('cards.id','cards.name','listings.*')->get();
        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, $listings);
    }
	
}
