<?php
/*
    Battleground Controller - github.com/CryptoCardsETH/backend
    File created by Harris Christiansen (HarrisChristiansen.com)
*/

namespace App\Http\Controllers;

class BattlegroundController extends Controller
{
    public function getIndex()
    {
        return view('pages.battleground.index');
    }
}
