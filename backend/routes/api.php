<?php

use Illuminate\Support\Facades\Response;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::post('/auth', 'AuthController@auth');

Route::get('/me', 'ProfileController@me');
Route::put('/me', 'ProfileController@updateMe');

Route::get('/me/cards', 'ProfileController@getMyCards');

Route::get('/cards', 'MarketplaceController@getAllCards');
Route::get('/cards/{id}', 'MarketplaceController@getCardDetail');
Route::put('/cards/{id}', 'MarketplaceController@updateCard');
Route::get('/listings', 'MarketplaceController@getAllListings');
Route::put('/cards/{id}/buy', 'MarketplaceController@putTransaction');

Route::get('/', function () {
    return Response::build(\App\Http\Controllers\Controller::RESPONSE_MESSAGE_SUCCESS, 'hi');
});
