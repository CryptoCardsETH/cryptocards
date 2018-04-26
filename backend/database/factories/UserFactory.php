<?php

use App\Helpers\EthereumConverter;
use App\Models\Card;
use Faker\Generator as Faker;

/*
|--------------------------------------------------------------------------
| Model Factories
|--------------------------------------------------------------------------
|
| This directory should contain each of the model factory definitions for
| your application. Factories provide a convenient way to generate new
| model instances for testing / seeding your application's database.
|
*/

$factory->define(App\Models\User::class, function (Faker $faker) {
    return [

        'nickname' => $faker->unique()->word.'-'.$faker->unique()->md5,
        'address'  => '0xfakex'.$faker->unique()->sha1,
    ];
});

$factory->define(Card::class, function (Faker $faker) {
    return [
        'token_id'   => Card::max('token_id') + 1,
        'name'       => $faker->word,
        'created_at' => $faker->dateTimeThisMonth($max = 'now', $timezone = null),
    ];
});

$factory->define(App\Models\Attribute::class, function (Faker $faker) {
    return [
        'name' => $faker->word,
    ];
});

$factory->define(App\Models\Transaction::class, function (Faker $faker) {
    return [
        'card_id'    => App\Models\Card::all()->random()->id,
        'user_id'    => App\Models\User::all()->random()->id,
        'price'      => EthereumConverter::convertETHPriceToInt($faker->randomFloat(13, 0, 2)),
        'created_at' => $faker->dateTimeThisMonth($max = 'now', $timezone = null),
    ];
});

$factory->define(App\Models\Listing::class, function (Faker $faker) {
    return [
        'card_id' => App\Models\Card::all()->random()->id,
        'user_id' => App\Models\User::all()->random()->id,
        'price'   => EthereumConverter::convertETHPriceToInt($faker->randomFloat(13, 0, 2)),
    ];
});

$factory->define(App\Models\Battle::class, function (Faker $faker) {
    $winner_id = factory(App\Models\BattleGroup::class)->create()->id;
    return [
        'group_1'        => $winner_id,
        'group_2'        => factory(App\Models\BattleGroup::class)->create()->id,
        'group_winner'   => $winner_id,
        'entrance_fee'   => EthereumConverter::convertETHPriceToInt($faker->randomFloat(13, 0, 2)),
        'created_at' => $faker->dateTimeThisMonth($max = 'now', $timezone = null),
    ];
});

$factory->define(App\Models\BattleGroup::class, function (Faker $faker) {
    return [
        'user_id'        => App\Models\User::all()->random()->id,
        'token_id'        => $faker->randomNumber(5),
    ];
});
