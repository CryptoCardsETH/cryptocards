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

Route::put('/follow/{id}', 'ProfileController@follow');
Route::get('/users', 'ProfileController@getAllUsers');
Route::get('/users/{id}', 'ProfileController@getUserDetail');
Route::get('/me/transactions', 'ProfileController@getMyTransactions');
Route::get('/me/notifications', 'ProfileController@getMyNotifications');

Route::get('/battles', 'BattlegroundController@getAllBattles');

Route::get('/cards', 'MarketplaceController@getAllCards');
Route::get('/cards/{id}', 'MarketplaceController@getCardDetail');
Route::put('/cards/{id}', 'MarketplaceController@updateCard');
Route::get('/listings', 'MarketplaceController@getAllListings');
Route::put('/cards/{id}/transaction', 'MarketplaceController@putTransaction');
Route::get('/cards/{id}/transactions', 'MarketplaceController@getCardTransactions');

Route::get('/stats/counts', 'StatsController@getCounts');
Route::get('/stats/transactionReport', 'StatsController@getTransactionReport');

Route::get('/contracts', 'ContractController@getContractAddresses');
Route::put('/contracts/ingest', 'ContractController@contractAddressIngest');

Route::get('/', function () {
    return Response::build(\App\Http\Controllers\Controller::RESPONSE_MESSAGE_SUCCESS, 'hi');
});
