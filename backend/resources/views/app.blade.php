<!doctype html>
<html><head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<meta http-equiv="Content-type" content="text/html;charset=UTF-8">

	<title>@yield('page-title') {{ env('APP_NAME') }}</title>
	<link REL="icon" HREF="/images/fav.png">

	<meta name="author" content="CryptoCards Team - Harris Christiansen, Nicky Semenza, Jason Shipp, Ben Bemis, David Worley, Lena Arafa">
	<meta name="description" content="CryptoCards is a game centered around unique playing cards which you can battle against other players. Each card is one-of-a-kind and 100% owned by you; it cannot be replicated, taken away, or destroyed.">
	<meta name="keywords" content="cryptocards, html5, php, laravel, game, ethereum, playing, cards, battle, cryptocurrency, collect">
	
	<!-- CSRF Token -->
	<meta name="csrf-token" content="{{ csrf_token() }}">

	<!-- BValidator -->
	<link href="/css/plugins/bvalidator.css" rel="stylesheet" type="text/css">
	
	<!-- jQuery UI -->
	<link rel="stylesheet" href="//code.jquery.com/ui/1.11.4/themes/smoothness/jquery-ui.css">
	
	<!-- Bootstrap CSS -->
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/css/bootstrap.min.css" integrity="sha384-PsH8R72JQ3SOdhVi3uxftmaW6Vc51MKb0q5P2rRUpPvrszuE4W1povHYgTpBfshb" crossorigin="anonymous">
	
	<!-- Bootstrap IE8 Support -->
	<!--[if lt IE 9]>
	    <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
		<script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
	<![endif]-->

	<!-- Fonts -->
	<link rel="stylesheet" href="//maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css">
	
	<!-- Site CSS -->
	<link rel="stylesheet" type="text/css" href="/css/site.css">

</head><body>

	<!-- ========== Navbar ========== -->
	<nav class="navbar navbar-expand-lg navbar-light bg-light">
		<a class="navbar-brand" href="{{ route('home') }}">{{ env('APP_NAME') }}</a>
		<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>

		<div class="collapse navbar-collapse" id="navbarSupportedContent">
			<ul class="navbar-nav mr-auto">
				<li class="nav-item">
					<a class="nav-link {{ Request::route()->named('marketplace') ? 'active':'' }}" href="{{ route('marketplace') }}" role="button" aria-expanded="false" aria-label="Marketplace">Marketplace</a>
				</li>
				<li class="nav-item">
					<a class="nav-link {{ Request::route()->named('collection') ? 'active':'' }}" href="{{ route('collection') }}" role="button" aria-expanded="false" aria-label="Card Collection">Card Collection</a>
				</li>
				<li class="nav-item">
					<a class="nav-link {{ Request::route()->named('battleground') ? 'active':'' }}" href="{{ route('battleground') }}" role="button" aria-expanded="false" aria-label="Battlegroud">Battleground</a>
				</li>
				<li class="nav-item dropdown">
					<a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Account</a>
					<div class="dropdown-menu" aria-labelledby="navbarDropdown">
						<a class="dropdown-item" href="{{ route('settings') }}">Settings</a>
						<div class="dropdown-divider"></div>
						<a class="dropdown-item" href="{{ route('logout') }}">Logout</a>
					</div>
				</li>
			</ul>
			<form class="form-inline my-2 my-lg-0">
				<input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search">
				<button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
			</form>
		</div>
	</nav>
	
	@if (session()->has('alert') || count($errors))
	<!-- ========== Alerts ========== -->
	<div id="alerts" class="container mt-3">
		@if (session()->has('alert'))
		<div class="alert {{ session()->get('alert-class', 'alert-success') }}" role="alert">{{ session()->get('alert') }}</div>
		<?php session()->forget('alert'); ?>
		@endif
		@foreach ($errors->all() as $error)
		<div class="alert alert-danger" role="alert">{{ $error }}</div>
		@endforeach
	</div>
	@endif
	
	
	<!-- ========== Content ========== -->
	@yield('content')

	<!-- ========== Javascript ========== -->
	
	<!-- jQuery / jQuery UI -->
	<script src="https://code.jquery.com/jquery-3.2.1.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
	<script src="https://code.jquery.com/ui/1.12.1/jquery-ui.min.js" integrity="sha256-VazP97ZCwtekAsvgPBSUwPFKdrwD3unUfSGVYrahUqU=" crossorigin="anonymous"></script>
	<script type="text/javascript" src="/js/plugins/jquery.cookie.js"></script>
	
	<!-- BValidator -->
	<script type="text/javascript" src="/js/plugins/jquery.bvalidator-yc.js"></script>

	<!-- Bootstrap JS -->
	<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.3/umd/popper.min.js" integrity="sha384-vFJXuSJphROIrBnz7yo7oB41mKfc8JzQZiCq4NCceLEaO4IHwicKwpJf9c9IpFgh" crossorigin="anonymous"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/js/bootstrap.min.js" integrity="sha384-alpBpkh1PFOepccYVYDB4do5UnbKysX5WZXm3XxPqe5iKTfUKjNkCk9SaVuEZflJ" crossorigin="anonymous"></script>
	
	<!-- Tablesorter -->
	<script type="text/javascript" src="/js/plugins/jquery.tablesorter.min.js"></script>
	
	<!-- Site JS -->
	<script type="text/javascript" src="/js/site.js"></script>
</body></html>