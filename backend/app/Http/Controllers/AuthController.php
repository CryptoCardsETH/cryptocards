<?php

namespace App\Http\Controllers;

use App\Helpers\EthereumHelper;
use App\Models\User;

class AuthController extends Controller
{
    public function __construct()
    {
        $this->middleware('jwt.auth', ['except' => ['auth']]);
    }

    public function auth()
    {
        $payload = json_decode(request()->getContent(), true);

        $signerAddress = $payload['address'];
        $signedMessage = $payload['signed'];
        $plainTextMessage = $payload['plaintext'];

        //Check signing
        if (!EthereumHelper::didAddressReallySignMessage($signerAddress, $signedMessage, $plainTextMessage)) {
            return response()->build(self::RESPONSE_MESSAGE_ERROR_DIGEST_SIGNING_MISMATCH);
        }

        //Get user based on address

        $user = User::getByAddress($signerAddress);
        $token = auth()->login($user);

        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, ['token'=>$token]);
    }
}
