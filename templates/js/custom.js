//Close alert
$(function(){
    window.setTimeout(function(){
        $('.alert').alert('close');
    },5000);
});

//Timer
var secondOne = $('.sec-one').text();
var secondTwo = $('.sec-two').text();
var int = setInterval(function() {
  if (secondTwo > 0) {
    secondTwo--;
    $('.sec-two').text(secondTwo);
  } else {
  	if (secondOne > 0) {
  		secondOne--;
  		secondTwo = 9;
  		$('.sec-one').text(secondOne);
  		$('.sec-two').text(secondTwo);
  	} else {
  		$('.sec-one').text(secondOne);
  		$('.sec-two').text(secondTwo);
  		clearInterval(int); 
  	}
  }
}, 1000);