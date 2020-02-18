//Close alert
$(function(){
    window.setTimeout(function(){
        $('.alert').alert('close');
    },5000);
});

//Timer
var secondOne = $('.sec-one').text();
var secondTwo = $('.sec-two').text();
var int = setInterval(function() { // запускаем интервал
  if (secondTwo > 0) {
    secondTwo--; // вычитаем 1
    $('.sec-two').text(secondTwo); // выводим получившееся значение в блок
  } else {
  	if (secondOne > 0) {
  		secondOne--;
  		secondTwo = 9;
  		$('.sec-one').text(secondOne);
  		$('.sec-two').text(secondTwo);
  	} else {
  		$('.sec-one').text(secondOne);
  		$('.sec-two').text(secondTwo);

  		learInterval(int); // очищаем интервал, чтобы он не продолжал работу при _Seconds = 0c
  	}

  }
}, 1000);