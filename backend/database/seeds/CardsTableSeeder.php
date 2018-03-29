<?php

use Illuminate\Database\Seeder;

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
        factory(App\Models\Card::class, 10)->create();

        DB::table('attribute_card')->insert(
            [
                'attribute_id' => factory(App\Models\Attribute::class)->create()->id,
                'card_id'      => factory(App\Models\Card::class)->create()->id,
                'value'        => $faker->randomFloat(2, 0, 10),
            ]
        );
    }
}
