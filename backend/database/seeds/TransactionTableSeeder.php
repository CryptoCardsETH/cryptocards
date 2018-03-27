<?php

use Illuminate\Database\Seeder;
use App\Helpers\EthereumConverter;

class TransactionTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        factory(App\Models\Transaction::class, 30)->create();
    }
}
