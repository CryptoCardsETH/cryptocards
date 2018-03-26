<?php

namespace App\Http\Controllers;

use App\Models\Card;
use App\Models\Listing;
use App\Models\User;

class StatsController extends Controller
{
    public function __construct()
    {
        $this->middleware(['jwt.auth', 'admin']);
    }

    public function getCounts()
    {
        $counts = [
            'user'    => User::count(),
            'card'    => Card::count(),
            'listing' => Listing::count(),
        ];

        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, $counts);
    }
}
