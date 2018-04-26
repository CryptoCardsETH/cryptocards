<?php
/*
    Marketplace Controller - github.com/CryptoCardsETH/backend
    File created by Harris Christiansen (HarrisChristiansen.com)
*/

namespace App\Http\Controllers;

use App\Helpers\EthereumConverter;
use App\Models\Card;
use App\Models\Listing;
use App\Models\Transaction;
use Illuminate\Support\Facades\Request;

class MarketplaceController extends Controller
{
    public function __construct()
    {
        $this->middleware('jwt.auth', ['except' => ['getAllCards', 'getCardDetail', 'getAllListings', 'estimateValue']]);
    }

    /**
     * Gets all the cards.
     *
     * @return mixed cards
     */
    public function getAllCards()
    {
        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, Card::with('attributes', 'user')->get());
    }

    public function getCardDetail($card_id)
    {
        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, Card::with('attributes', 'user')->find($card_id));
    }

    public function updateCard($card_id)
    {
        $user = auth()->user();
        $card = Card::find($card_id);
        if (!$card->isUserOwner($user)) {
            return response()->build(self::RESPONSE_MESSAGE_ERROR_UNAUTHORIZED);
        }

        $data = json_decode(Request::getContent(), true);
        foreach ($data as $key => $value) {
            if (in_array($key, [
                Card::FIELD_HIDDEN_TOGGLE,
            ])) {
                $card->$key = $value;
            }
        }

        $card->save();

        return $this->getCardDetail($card_id);
    }

    public function putTransaction($card_id)
    {
        //update card user's id with current id
        $user = auth()->user();
        $card = Card::find($card_id);
        $listing = Listing::where('card_id', $card_id)->first();
        if ($card->isUserOwner($user)) {
            return response()->build(self::RESPONSE_MESSAGE_ERROR_UNAUTHORIZED, 'User already owns the card');
        }

        $data = json_decode(Request::getContent(), true);
        $card->user_id = $user->id;

        $card->save();

        //add transaction
        $transaction = new Transaction();
        $transaction->card_id = $card_id;
        $transaction->user_id = $user->id;
        $price = EthereumConverter::convertETHPriceToInt($listing->price);
        $transaction->price = $price;

        $transaction->save();

        //remove card from listings table
        Listing::where('card_id', $card_id)->delete();

        return $this->getCardDetail($card_id);
    }

    public function getCardTransactions($card_id)
    {
        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, Card::find($card_id)->transactions);
    }

    public function getAllListings()
    {
        $listings = Listing::with(['cards'])->get();

        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, $listings);
    }

    public function estimateValue($card_id)
    {
        $card = Card::findOrFail($card_id);
        if ($card->transactions()->count() >= 1) {
            $avg_price = $card->transactions()->orderBy('created_at', 'DESC')->limit(5)->avg('price');
            $value = EthereumConverter::convertETHPriceToFloat(round($avg_price + ($avg_price * .5), 13));
        } else {
            $value = EthereumConverter::convertETHPriceToFloat(rand(250000000, 1250000000) * pow(10, 4));
        }

        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, $value);
    }
}
