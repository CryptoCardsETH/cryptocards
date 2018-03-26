<?php

namespace App\Http\Middleware;

use App\Http\Controllers\Controller;
use Closure;

class CheckAdmin
{
    /**
     * Handle an incoming request.
     *
     * @param \Illuminate\Http\Request $request
     * @param \Closure                 $next
     *
     * @return mixed
     */
    public function handle($request, Closure $next)
    {
        if (!auth()->user()->admin) {
            return response()->build(Controller::RESPONSE_MESSAGE_ERROR_UNAUTHORIZED);
        }

        return $next($request);
    }
}
