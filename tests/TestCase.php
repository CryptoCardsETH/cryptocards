<?php

namespace Tests;

use Illuminate\Foundation\Testing\TestCase as BaseTestCase;

abstract class TestCase extends BaseTestCase
{
    use CreatesApplication;
    public function authenticatedJSON($method, $route, $params, $token)
    {
        return $this->json($method, $route, $params, ['HTTP_Authorization' => 'Bearer '.$token]);
    }
}
