<?php
namespace App\Providers;

use App\Http\Controllers\Controller;
use Illuminate\Support\ServiceProvider;
use Illuminate\Support\Facades\Response;
class ResponseServiceProvider extends ServiceProvider
{
    /**
     * Register the application's response macros.
     * success: for API success
     * error: for API error.
     * @return void
     */
    public function boot()
    {
        Response::macro('success', function ($message, $data) {
            return Response::json([
                'success'       => true,
                'message'       => $message,
                'data'          => $data,
            ], Controller::RESPONSE_MESSAGE_HTTP_CODES[$message]);
        });
        Response::macro('error', function ($message, $data = null) {
            return Response::json([
                'success'       => false,
                'message'       => $message,
                'data'          => $data,
            ], Controller::RESPONSE_MESSAGE_HTTP_CODES[$message]);
        });
    }

    /**
     * Register the application services.
     *
     * @return void
     */
    public function register()
    {
        //
    }
}