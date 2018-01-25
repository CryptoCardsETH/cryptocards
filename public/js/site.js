/*
	@ Harris Christiansen (code@HarrisChristiansen.com)
	Project: Cryptocards - github.com/CryptoCardsETH/backend
*/


$(document).ready(function() { 
	$(".sortable").tablesorter();

	$('.validate').bValidator();

	$(".datepicker").datepicker({
		dateFormat: 'yy-mm-dd',
	});

	$("#alert").click(function(evt) {
		displayMessage("Example alert!");
	});
	
	// Turn Table Rows Into Links
	$('.table-links tbody tr').click(function() {
		var url = $(this).find("a").attr("href");
		if (url) {
			window.location = url;
		}
	});
	
	$(".categoriesTypeahead").autocomplete({
		source: "/categories/typeahead",
		minLength: 2,
	});
});

function displayMessage(msg, isFailure=false) {
	if (isFailure) {
		$('<div class="alert alert-danger" role="alert">'+msg+"</div>").appendTo('#alerts').delay(5000).slideUp(300);
	} else {
		$('<div class="alert alert-success" role="alert">'+msg+"</div>").appendTo('#alerts').delay(5000).slideUp(300);
	}
}