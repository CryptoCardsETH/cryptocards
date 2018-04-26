<?php

use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $this->call(UserTableSeeder::class);
        $this->call(CardsTableSeeder::class);
        $this->call(TransactionTableSeeder::class);
        $this->call(ListingTableSeeder::class);
        $this->call(BattleTableSeeder::class);
    }
}
