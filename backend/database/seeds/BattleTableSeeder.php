<?php

use Illuminate\Database\Seeder;

class BattleTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        factory(App\Models\Battle::class, 1)->create();
        factory(App\Models\Battle::class, 1)->create();
        factory(App\Models\Battle::class, 1)->create();
        factory(App\Models\Battle::class, 1)->create();
        factory(App\Models\Battle::class, 1)->create();
        factory(App\Models\Battle::class, 1)->create();
        factory(App\Models\Battle::class, 1)->create();
        factory(App\Models\Battle::class, 1)->create();
    }
}
