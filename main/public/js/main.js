(function( $ ) {
  
	$('#btnnext').click(function(e){

		$('#wrapper').addClass('hide');
		$('form').removeClass('hide');

	})

	$('form').submit(function(e) {
  		var val = $('#inputPassword').val();

  		var data = { Password: val };
  		
  		
  		$.ajax({
  			type: "POST",
  			url: '/post',
  			data: JSON.stringify(data),
  			success: function(resp){
  				var temp = $('#sha');
  				temp.html("Success: " +  resp);
  				temp.removeClass('hide');
  			},
 			error:function(resp){
 				console.log(resp);

 			}
		});

  		e.preventDefault();
	});




})( jQuery );