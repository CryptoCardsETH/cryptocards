<?php
/*
	Profile Controller - github.com/CryptoCardsETH/backend
	File created by Harris Christiansen (HarrisChristiansen.com)
*/

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class ProfileController extends Controller {
	
	public function getIndex() {
		return view('pages.profile.index');
	}
	
	public function getCollection() {
		return view('pages.profile.collection');
	}
	
	public function getSettings() {
		return view('pages.profile.settings');
	}
	
	public function getLogout() {
		return redirect()->route('home')->with('alert', "You have been logged out");
	}
	
}
