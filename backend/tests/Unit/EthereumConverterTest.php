<?php

namespace Tests\Unit;

use App\Helpers\EthereumConverter;
use Tests\TestCase;

class EthereumConverterTest extends TestCase
{
    /**
     * A basic test example.
     *
     * @return void
     */
    public function testConvertPriceIntToFloat()
    {
        $this->assertEquals(0.0000123456789, EthereumConverter::convertETHPriceToFloat(123456789));
        $this->assertEquals(0.01, EthereumConverter::convertETHPriceToFloat(100000000000));
        $this->assertEquals(1, EthereumConverter::convertETHPriceToFloat(10000000000000));
    }

    public function testConvertPriceFloatToInt()
    {
        $this->assertEquals(123456789, EthereumConverter::convertETHPriceToInt(0.0000123456789));
        $this->assertEquals(100000000000, EthereumConverter::convertETHPriceToInt(0.01));
        $this->assertEquals(10000000000000, EthereumConverter::convertETHPriceToInt(1));
    }

    public function testConvertPriceBackAndForthRandom()
    {
        $faker = \Faker\Factory::create();
        $float1 = $faker->randomFloat();
        $bigint1 = $faker->randomNumber(5);
        $this->assertEquals(EthereumConverter::convertETHPriceToFloat(EthereumConverter::convertETHPriceToInt($float1)), $float1);
        $this->assertEquals(EthereumConverter::convertETHPriceToInt(EthereumConverter::convertETHPriceToFloat($bigint1)), $bigint1);
    }
}
