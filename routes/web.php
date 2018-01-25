<?php

/*
|--------------------------------------------------------------------------
| Web Routes - CryptoCards - github.com/CryptoCardsETH/backend
|--------------------------------------------------------------------------
*/

Route::get('/', 'HomeController@getIndex')->name('home');
Route::group(['prefix' => 'profile'], function () {
	Route::get('', 'ProfileController@getIndex')->name('profile');
	Route::get('collection', 'ProfileController@getCollection')->name('collection');
	Route::get('settings', 'ProfileController@getSettings')->name('settings');
	Route::get('logout', 'ProfileController@getLogout')->name('logout');
});
Route::group(['prefix' => 'marketplace'], function () {
	Route::get('', 'MarketplaceController@getIndex')->name('marketplace');
});

Route::group(['prefix' => 'battleground'], function () {
	Route::get('', 'BattlegroundController@getIndex')->name('battleground');
});