<?php

namespace App\Providers;

use Illuminate\Support\Facades\Response;
use Illuminate\Support\ServiceProvider;

class ResponseServiceProvider extends ServiceProvider
{
    /**
     * Register the application's response macros.
     * success: for API success
     * error: for API error.
     *
     * @return void
     */
    public function boot()
    {
        Response::macro('build', function ($responseMessage, $data = null) {
            return Response::json([
                'success'     => $responseMessage[0] === 1, //response message id 1 is success
                'message'     => $responseMessage[2],
                'data'        => $data,
                'response_id' => $responseMessage[0],
            ], $responseMessage[1]);
        });
    }
}
