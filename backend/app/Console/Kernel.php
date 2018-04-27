<?php

namespace App\Console;

use App\Console\Commands\BlockchainPlayground;
use App\Console\Commands\CreateCard;
use App\Console\Commands\IngestBattleGroupsFromBlockchain;
use App\Console\Commands\IngestBattlesFromBlockchain;
use App\Console\Commands\IngestCardsFromBlockchain;
use App\Console\Commands\ProtoTest;
use App\Console\Commands\SeedAttributesRandomly;
use Illuminate\Console\Scheduling\Schedule;
use Illuminate\Foundation\Console\Kernel as ConsoleKernel;

class Kernel extends ConsoleKernel
{
    /**
     * The Artisan commands provided by your application.
     *
     * @var array
     */
    protected $commands = [
        ProtoTest::class,
        IngestBattleGroupsFromBlockchain::class,
        IngestCardsFromBlockchain::class,
        IngestBattlesFromBlockchain::class,
        CreateCard::class,
        BlockchainPlayground::class,
        SeedAttributesRandomly::class,
    ];

    /**
     * Define the application's command schedule.
     *
     * @param \Illuminate\Console\Scheduling\Schedule $schedule
     *
     * @return void
     */
    protected function schedule(Schedule $schedule)
    {
        // $schedule->command('inspire')
        //          ->hourly();
        $schedule->command(IngestBattleGroupsFromBlockchain::class)->everyMinute();
        $schedule->command(IngestCardsFromBlockchain::class)->everyMinute();
    }

    /**
     * Register  the commands for the application.
     *
     * @return void
     */
    protected function commands()
    {
        $this->load(__DIR__.'/Commands');

        require base_path('routes/console.php');
    }
}
