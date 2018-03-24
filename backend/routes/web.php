<?php

/*
|--------------------------------------------------------------------------
| Web Routes - CryptoCards - github.com/CryptoCardsETH/backend
|--------------------------------------------------------------------------
*/

Route::get('/', function () {
    return Response::build(\App\Http\Controllers\Controller::RESPONSE_MESSAGE_SUCCESS, 'CryptoCards API');
});
