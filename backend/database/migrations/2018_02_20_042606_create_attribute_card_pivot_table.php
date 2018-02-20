<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;

class CreateAttributeCardPivotTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('attribute_card', function (Blueprint $table) {
            $table->integer('attribute_id')->unsigned()->index();
            $table->foreign('attribute_id')->references('id')->on('attributes')->onDelete('cascade');
            $table->integer('card_id')->unsigned()->index();
            $table->foreign('card_id')->references('id')->on('cards')->onDelete('cascade');
            $table->primary(['attribute_id', 'card_id']);
            $table->integer('value');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('attribute_card');
    }
}
