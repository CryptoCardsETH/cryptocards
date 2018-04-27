<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use App\Models\Card;
use App\Models\Attribute;
use DB;
class SeedAttributesRandomly extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'attributes';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Command description';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        //
        $attributes = Attribute::all()->pluck('id')->toArray(); 
        DB::table('attribute_card')->truncate();
        foreach(Card::all() as $card) {
            foreach($attributes as $a) {
                $card->attributes()->attach($a,['value'=>rand(0,100)]);
            }
            $card->save();
        }
        
    }
}
