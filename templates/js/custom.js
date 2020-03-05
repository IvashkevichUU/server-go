var playerData = [
                  {'id':1,'name':'Player_1','room':1,'game':1,'bet':1,'rate':8},
                  {'id':2,'name':'Player_2','room':1,'game':1,'bet':4,'rate':31},
                  {'id':3,'name':'Player_3','room':1,'game':1,'bet':2,'rate':15},
                  {'id':4,'name':'Player_4','room':1,'game':1,'bet':6,'rate':46}
                  ];
var currentPlayerData = [];
var totalPrizeGame = 0;
var changeName;
var p = 0;
var inGame = false;
var currentID,currentName,currentColor,betPlayer;

//Close alert
$(function(){
    window.setTimeout(function(){
        $('.alert').alert('close');
    },100);
});

//Timer
var inGameProgress = function() {

  if (inGame == false) {

    inGame = true;

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

  } 

};

//Spinner anim
var spinnerAnimation = function() {
var widthProgress = $('.progress').width();
var progressStop = $('.progress .bg-success');
var widthProgressStop = Math.floor(Math.random() * 100) + 1;
var arrowStop = widthProgressStop/(widthProgress/100);

var stopWin = Math.floor(Math.random() * 100) + 1;

$('.spinner-arrow').slideDown();

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
                        "left":"+="+stopWin+"%"
                    }, 2500, function () {
                        $(this).css("background", "red");
                        clearInterval(changeName);
                        $('.round-tittle').text($('.round-tittle').text() + ' wins $' + totalPrizeGame);

                        window.setTimeout(function(){
                            //$('.progress').empty();
                            inGame = false;
                        },3000);

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

  changeName = setInterval(function() {


    progressPlayers.map(function(indx, element){
      
      minPosEl = $(element).position().left;
      widthEl = $(element).width();
      posArrow = progressArrow.position().left;
      maxPosEl = minPosEl+widthEl;

      if (minPosEl <= posArrow){
       if (posArrow <= maxPosEl) {
        $('.round-tittle').text($(element).text());
      }
    }
    });
  }, 100);

};

//Make bet
$( ".make-bet" ).on( "click", function(e) {
  e.preventDefault();

  if (inGame == true) {

    betPlayer = parseInt($('#exampleInputBet1').val());
    currentID = p;
    currentName = 'Gamer ' +p;
    currentColor = "#"+((1<<24)*Math.random()|0).toString(16);
    currentPlayerData.push({id:currentID,name:currentName,room:1,game:1,bet:betPlayer,rate:8,color:currentColor});

    totalPrizeRoundAndProcent();

  } else {

    inGameProgress();

  }

});  


//RESKIN
var updateUI = function() {
p = 0;
$('.progress').empty();
$('.table-current-game tbody').empty();
for(key in currentPlayerData) {
  p++;
  var strTable = "<tr><th scope='row'>"+p+"</th><td>"+currentPlayerData[key].name+"</td><td>"+currentPlayerData[key].bet+"</td><td>"+currentPlayerData[key].rate+"</td></tr>";
  var strProgress = "<div class='progress-bar' role='progressbar' style='width: "+currentPlayerData[key].rate+"%;background-color:"+currentPlayerData[key].color+"' aria-valuenow='"+currentPlayerData[key].rate+"' aria-valuemin='0' aria-valuemax='100'>"+currentPlayerData[key].name+"</div>";

  $('.progress').append(strProgress);
  $('.table-current-game tbody').append(strTable);

}

};


//Total prize round <--SERVER
var totalPrizeRoundAndProcent = function() {
totalPrizeGame = 0;
for(key in currentPlayerData) {
  totalPrizeGame += currentPlayerData[key].bet;
}

for(key in currentPlayerData) {
  currentPlayerData[key].rate = Math.round((currentPlayerData[key].bet / totalPrizeGame) * 100);
  console.log(key + " = " + currentPlayerData[key].bet + " rate " + currentPlayerData[key].rate);

}


console.log(" Total Prize " + totalPrizeGame);
console.log(currentPlayerData);

updateUI();


};

//totalPrizeRoundAndProcent();
// console.log(" Total Prize " + totalPrizeGame);
// console.log(currentPlayerData);