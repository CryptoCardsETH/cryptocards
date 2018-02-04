<?php

namespace App\Http\Controllers;

use Illuminate\Foundation\Bus\DispatchesJobs;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Foundation\Auth\Access\AuthorizesRequests;

class Controller extends BaseController
{
    const RESPONSE_MESSAGE_ERROR_JWT_ERROR = 'No JWT provided with request!';
    const RESPONSE_MESSAGE_ERROR_JWT_EXPIRED = 'JWT has expired!';
    const RESPONSE_MESSAGE_ERROR_JWT_INVALID = 'JWT is invalid!';
    const RESPONSE_MESSAGE_ERROR_UNAUTHORIZED = 'Not authorized!';
    const RESPONSE_MESSAGE_ERROR_DIGEST_SIGNING_MISMATCH = 'Signed message signature does not match! Login failed.';
    const RESPONSE_MESSAGE_SUCCESS = 'Success!';

    const RESPONSE_MESSAGE_HTTP_CODES = [
        self::RESPONSE_MESSAGE_ERROR_JWT_ERROR => 403,
        self::RESPONSE_MESSAGE_ERROR_JWT_EXPIRED => 403,
        self::RESPONSE_MESSAGE_ERROR_JWT_INVALID => 403,
        self::RESPONSE_MESSAGE_ERROR_UNAUTHORIZED => 401,
        self::RESPONSE_MESSAGE_ERROR_DIGEST_SIGNING_MISMATCH => 403,
        self::RESPONSE_MESSAGE_SUCCESS => 200,
    ];

    use AuthorizesRequests, DispatchesJobs, ValidatesRequests;
}
