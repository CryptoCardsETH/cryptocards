<?php

use Illuminate\Database\Migrations\Migration;

class AddDefaultToBattleGroupsSortOrder extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('battle_group_cards', function ($table) {
            $table->integer('sort_order')->default(0)->change();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        //
    }
}
