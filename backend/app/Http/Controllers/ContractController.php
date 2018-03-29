<?php

namespace App\Http\Controllers;

use App\Models\Contract;
use Illuminate\Support\Facades\Request;
use Log;

class ContractController extends Controller
{
    public function contractAddressIngest()
    {
        if (env('APP_DEBUG') != 'true') {
            //this feature is only really for the development workflow, so no auth needed.
            return response()->build(self::RESPONSE_MESSAGE_ERROR_UNAUTHORIZED);
        }

        $data = json_decode(Request::getContent(), true);
        Log::info($data);

        foreach ($data as $contractName => $data) {
            $address = $data['address'];
            Log::info($contractName.' @ '.$address);
            Contract::updateAddress($contractName, $address, $data['transactionHash']);
        }

        return response()->build(self::RESPONSE_MESSAGE_SUCCESS);
    }

    public function getContractAddresses()
    {
        $c = Contract::all()->keyBy(Contract::FIELD_NAME);

        return response()->build(self::RESPONSE_MESSAGE_SUCCESS, $c);
    }
}
