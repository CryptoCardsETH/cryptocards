<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;

class CreateBattlesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('battles', function (Blueprint $table) {
            $table->increments('id');
            $table->integer('group_1')->unsigned();
            $table->foreign('group_1')->references('id')->on('battle_groups');
            $table->integer('group_2')->unsigned();
            $table->foreign('group_2')->references('id')->on('battle_groups');
            $table->integer('group_winner')->unsigned()->nullable();
            $table->foreign('group_winner')->references('id')->on('battle_groups');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('battles');
    }
}
