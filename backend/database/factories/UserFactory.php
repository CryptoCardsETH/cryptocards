<?php

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

        'nickname' => $faker->unique()->firstName,
//        'email' => $faker->unique()->safeEmail,
        'address' => '0xfakex'.$faker->unique()->sha1,
    ];
});

$factory->define(App\Models\Card::class, function (Faker $faker) {
    return [
        'name' => $faker->word,
    ];
});

$factory->define(App\Models\Attribute::class, function (Faker $faker) {
    return [
        'name' => $faker->word,
    ];
});
