<?php

namespace App\Http\Controllers;

use App\Models\User;

class AuthController extends Controller {

    public function __construct()
    {
        $this->middleware('jwt.auth', ['except' => ['auth']]);
    }

    /**
     * TODO:
     * This needs to verify that the message was properly signed.
     * Just like web3js does: http://web3js.readthedocs.io/en/1.0/web3-eth-accounts.html#recover
     * @param $signerAddress
     * @param $signedMessage
     * @param $plainTextMessage
     * @return bool
     */
    static function didAddressReallySignMessage($signerAddress, $signedMessage, $plainTextMessage)
    {
        return true;
    }

    public function auth() {
        $payload = json_decode(request()->getContent(), true);

        $signerAddress = $payload['address'];
        $signedMessage = $payload['signed'];
        $plainTextMessage = $payload['plaintext'];

        //Check signing
        if(! self::didAddressReallySignMessage($signerAddress,$signedMessage,$plainTextMessage)) {
            return ['error'=>'Signing mismatch.'];
        }

        //Get user based on address
        $user = User::firstOrCreate(['address'=>$signerAddress]);

        $token = auth()->login($user);

        return ['token'=>$token];
    }
    public function me()
    {
        return response()->json(auth()->user());
    }

}
