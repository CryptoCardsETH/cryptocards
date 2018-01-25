<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateItemsTable extends Migration
{
    public function up()
    {
        Schema::create('items', function (Blueprint $table) {
            $table->increments('id');
            $table->string('name')->nullable();
            $table->integer('category_id')->foreign()->references('id')->on('categories')->nullable();
            $table->text('description')->nullable();
            $table->text('url')->nullable();
            $table->integer('quantity')->default(1);
            $table->timestamps();
            $table->softDeletes();
        });
    }
    
    public function down()
    {
        Schema::dropIfExists('items');
    }
}
