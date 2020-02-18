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
  		spinnerAnimation();
  	}
  }
}, 1000);

//Spinner anim
var spinnerAnimation = function() {
var widthProgress = $('.progress').width()-6;
var progressStop = $('.progress .bg-success');
var widthProgressStop = (progressStop.position().left+progressStop.width()/2)-6;
$('.spinner-arrow').slideDown();
//console.log(widthProgress);
//console.log(widthProgressStop);
    $('.spinner-arrow').animate({
        "margin-left":"+="+widthProgress+"px"
    }, 500, function () {
        $(this).animate({
            "margin-left": "0px"
        }, 1000, function () {
            $(this).animate({
                "margin-left":"+="+widthProgress+"px"
            }, 1500, function () {
                $(this).animate({ 
                    "margin-left": "0px"
                }, 2000, function () {
                    $(this).animate({
                        "margin-left":"+="+widthProgressStop+"px"
                    }, 2500, function () {
                        $(this).css("background", "red")
                    });
                });
            });
        });
    });
};



