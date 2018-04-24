<?php

namespace App\Console\Commands;

use App\Models\User;
use App\Notifications\BattleCompletion;
use Illuminate\Console\Command;
use Notification;

class BattleNotificationTest extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'battlenotification';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'battle notification test';

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
        Notification::send(User::all(), new BattleCompletion());
    }
}
