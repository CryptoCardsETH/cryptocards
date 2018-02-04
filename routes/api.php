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
Route::get('/me', 'AuthController@me');

Route::get('/', function () {
    return Response::build(\App\Http\Controllers\Controller::RESPONSE_MESSAGE_SUCCESS,'hi');
});
