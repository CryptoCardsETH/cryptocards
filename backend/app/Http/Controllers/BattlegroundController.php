<?php
/*
    Battleground Controller - github.com/CryptoCardsETH/backend
    File created by Harris Christiansen (HarrisChristiansen.com)
*/

namespace App\Http\Controllers;

use App\Models\Battle;

class BattlegroundController extends Controller
{
    public function getAllBattles()
    {
        $battles = Battle::with('group1.group_cards.card', 'group2.group_cards.card')->get();

        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, $battles);
    }
}
