<?php

use Illuminate\Database\Seeder;
use Faker\Generator as Faker;
class CardsTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $faker = \Faker\Factory::create();

        DB::table('attribute_card')->insert(
            [
                'attribute_id' => factory(App\Models\Attribute::class)->create()->id,
                'card_id' => factory(App\Models\Card::class)->create()->id,
                'value' => $faker->randomFloat(2,0,10)
            ]
        );
    }
}
