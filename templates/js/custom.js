var changeName;

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
var widthProgress = $('.progress').width();
var progressStop = $('.progress .bg-success');
var widthProgressStop = (progressStop.position().left+(progressStop.width()/2))-8;
var arrowStop = widthProgressStop/(widthProgress/100);
$('.spinner-arrow').slideDown();
// console.log(widthProgress);
// console.log(progressStop.position().left);
// console.log(progressStop.width());
// console.log(widthProgressStop);
// console.log(arrowStop);
spinnerChangeName();
    $('.spinner-arrow').animate({
        "left":"+=100%"
    }, 500, function () {
        $(this).animate({
            "left": "-8px"
        }, 1000, function () {
            $(this).animate({
                "left":"+=100%"
            }, 1500, function () {
                $(this).animate({ 
                    "left": "0%"
                }, 2000, function () {
                    $(this).animate({
                        "left":"+="+arrowStop+"%"
                    }, 2500, function () {
                        $(this).css("background", "red");
                        clearInterval(changeName);
                        $('.round-tittle').text($('.round-tittle').text() + ' wins $16.49');
                    });
                });
            });
        });
    });
    
};

//Change player name
var spinnerChangeName = function() {
var progressArrow = $('.progress-container .spinner-arrow');
var progressPlayers = progressArrow.next().children('.progress-bar');
      var posEl;
      var widthEl;
      var posArrow;
//console.log(progressPlayers);
//console.log(progressPlayers.length);
//console.log(progressArrow.position().left);

  changeName = setInterval(function() {

// var i = 0;
// for (i = 0; i < progressPlayers.length; i++) {
//  //console.log(progressPlayers);
//  //console.log(progressPlayers[i]);

// var element = progressPlayers[i];

//       minPosEl = $(element).position().left;
//       widthEl = $(element).width();
//       posArrow = progressArrow.position().left;
//       maxPosEl = minPosEl+widthEl;

//       if (minPosEl < posArrow){

//         if (posArrow < maxPosEl) {
//         console.log($(element).text());
//         $('.round-tittle').text($(element).text());
//       }
//     }
// }

    progressPlayers.map(function(indx, element){
      
      minPosEl = $(element).position().left;
      widthEl = $(element).width();
      posArrow = progressArrow.position().left;
      maxPosEl = minPosEl+widthEl;

      // console.log('posEl',minPosEl);
      // console.log('posArrow',posArrow);
      // console.log('widthEl',maxPosEl);

      if (minPosEl <= posArrow){
       if (posArrow <= maxPosEl) {
        // console.log(indx);
        //console.log($(element).text());
        $('.round-tittle').text($(element).text());
      }
    }
    });
  }, 100);

};

