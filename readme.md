# CryptoCards

[![Build Status](https://travis-ci.org/CryptoCardsETH/cryptocards.svg?branch=master)](https://travis-ci.org/CryptoCardsETH/cryptocards)

[![codecov](https://codecov.io/gh/CryptoCardsETH/cryptocards/branch/master/graph/badge.svg)](https://codecov.io/gh/CryptoCardsETH/cryptocards)


## Synopsis

[WIP] CryptoCards is a game centered around unique playing cards which you can battle against other players. Each card is one-of-a-kind and 100% owned by you; it cannot be replicated, taken away, or destroyed.  

Available at [GitHub.com/CryptoCardsETH/](https://github.com/CryptoCardsETH/)   

## Running 

### Locally

coming soon...

### With docker

Run `docker-compose up` from the root directory. Frontend will be on :5000 if you've run the setup proceedures below:

The frontend automatically installs yarn dependencies during the image build process.

To run `php artisan migrate` and `composer install` and other backend tooling, get a shell with `docker-compose exec fpm /bin/bash`

`docker-compose exec fpm php artisan test:proto` to test the laravel <-> golang communication over gRPC (give `ethereum_proxy` ample time to start up)

`./docker-tooling.sh generate` to generate run abigen and proto generators.

`ganache-cli` is running in deterministic mode, with network id of 999999. This means the first account should be `0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1` with private key of `4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d`

Mysql is exposed on port _3307_ (one higher than the default, so it doesn't conflict with any mysql on the docker host.)


## License

MIT License  

Copyright (c) 2018 [CryptoCards](https://github.com/CryptoCardsETH)  

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:  

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.  

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.  
