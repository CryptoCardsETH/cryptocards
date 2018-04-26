<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;

class AddTokenIdToBattleGroupsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('battle_groups', function (Blueprint $table) {
            $table->integer('token_id')->unique();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::table('battle_groups', function (Blueprint $table) {
            $table->dropColumn('token_id');
        });
    }
}
