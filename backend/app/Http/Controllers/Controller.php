<?php

namespace App\Http\Controllers;

use Illuminate\Foundation\Auth\Access\AuthorizesRequests;
use Illuminate\Foundation\Bus\DispatchesJobs;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Routing\Controller as BaseController;

class Controller extends BaseController
{
    //format of response message objects: [id, http_status_code, message]
    //success
    const RESPONSE_MESSAGE_SUCCESS = [1, 200, 'ok'];

    //error
    const RESPONSE_MESSAGE_ERROR_JWT_ERROR = [10, 403, 'No JWT provided with request!'];
    const RESPONSE_MESSAGE_ERROR_JWT_EXPIRED = [11, 403, 'JWT has expired!'];
    const RESPONSE_MESSAGE_ERROR_JWT_INVALID = [12, 403, 'JWT is invalid!'];
    const RESPONSE_MESSAGE_ERROR_UNAUTHORIZED = [13, 401, 'Not authorized!'];
    const RESPONSE_MESSAGE_ERROR_DIGEST_SIGNING_MISMATCH = [14, 403, 'Signed message signature does not match! Login failed.'];
    const RESPONSE_MESSAGE_ALREADY_FOLLOWING = [15, 400, 'Already Following'];
    const RESPONSE_MESSAGE_ERROR_NOT_FOUND = [16, 404, 'Not found!'];
    use AuthorizesRequests, DispatchesJobs, ValidatesRequests;
}
