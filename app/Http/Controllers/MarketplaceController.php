<?php
/*
	Marketplace Controller - github.com/CryptoCardsETH/backend
	File created by Harris Christiansen (HarrisChristiansen.com)
*/

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MarketplaceController extends Controller {
	
	public function getIndex() {
		return view('pages.marketplace.index');
	}
	
}
