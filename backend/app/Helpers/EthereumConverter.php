<?php

namespace App\Helpers;

class EthereumConverter
{
    public static function convertETHPriceToFloat($integerPrice)
    {
        return $integerPrice / pow(10, 13);
    }

    public static function convertETHPriceToInt($ethPriceFloat)
    {
        return  $ethPriceFloat * pow(10, 13);
    }
}
