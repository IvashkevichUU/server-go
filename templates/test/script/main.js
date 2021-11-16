let Application			= PIXI.Application,
	Container			= PIXI.Container,
	loader				= PIXI.loader,
	resources			= PIXI.loader.resources,
	Graphics			= PIXI.Graphics,
	TextureCache		= PIXI.utils.TextureCache,
	Sprite				= PIXI.Sprite;

let app = new Application({
	width: 1390,
	height: 640,
	antialias: true,
	autoResize: true,
});

document.body.appendChild( app.view );

let ladder_active = 0, ladder_previous;

loader
	.add("images/atlas_web_0.json")
	.add("images/atlas_web_1.json")
	.add("images/atlas_web_2.json")
	.load( main );

function main() {
	id_0 = resources["images/atlas_web_0.json"].textures;
	id_1 = resources["images/atlas_web_1.json"].textures;
	id_2 = resources["images/atlas_web_2.json"].textures;

	gameScene = new Container();
	app.stage.addChild( gameScene );

	bg = new Sprite( id_2["background.jpg"] );
	bg.anchor.set(0, 0);
	gameScene.addChild( bg );

	//Other

	let plant_3 = new Sprite( id_0["plant_1.png"] );
	plant_3.position.set( 1135, 164 );
	gameScene.addChild( plant_3 );

		let book_table = new Sprite( id_0["book_table.png"] );
	book_table.position.set( 834, 0 );
	gameScene.addChild( book_table );

	//Ladder

	let ladder = new Container();
	ladder.position.set( app.stage.width, app.stage.height / 1.68 );
	gameScene.addChild( ladder );

	let ladder_0 = new Sprite( id_1["ladder_0.png"] );
	ladder_0.anchor.set(1, 0.5);
	ladder_0.position.set( 0, 0 );
	ladder.addChild( ladder_0 );

	let ladder_1 = new Sprite( id_1["ladder_1.png"] );
	ladder_1.anchor.set(1, 0.6);
	ladder_1.position.set( 0, 0 );
	ladder_1.alpha = 0;
	ladder.addChild( ladder_1 );

	let ladder_2 = new Sprite( id_1["ladder_2.png"] );
	ladder_2.anchor.set(1, 0.6);
	ladder_2.position.set( 0, 0 );
	ladder_2.alpha = 0;
	ladder.addChild( ladder_2 );

	let ladder_3 = new Sprite( id_1["ladder_3.png"] );
	ladder_3.anchor.set(1, 0.6);
	ladder_3.position.set( 0, 0 );
	ladder_3.alpha = 0;
	ladder.addChild( ladder_3 );

	//Ladder menu

	let ladder_menu = new Container();
	ladder_menu.position.set( app.stage.width / 2 + 200, 70 );
	gameScene.addChild( ladder_menu );

	let ladder_cont_1 = new Container();
	ladder_cont_1.position.set( 0, 0 );
	ladder_cont_1.alpha = 0;
	ladder_menu.addChild( ladder_cont_1 );

	let ladder_cont_2 = new Container();
	ladder_cont_2.position.set( 0, 0 );
	ladder_cont_2.alpha = 0;
	ladder_menu.addChild( ladder_cont_2 );

	let ladder_cont_3 = new Container();
	ladder_cont_3.position.set( 0, 0 );
	ladder_cont_3.alpha = 0;
	ladder_menu.addChild( ladder_cont_3 );

	let menu_clean_1 = new Sprite( id_0["menu_clean.png"] );
	menu_clean_1.anchor.set(0.5, 0.47);
	menu_clean_1.position.set( 0, 0);
	menu_clean_1.interactive = false;
	menu_clean_1.buttonMode = true;
	ladder_cont_1.addChild( menu_clean_1 );

	let menu_green_1 = new Sprite( id_0["menu_green.png"] );
	menu_green_1.anchor.set(0.5, 0.5);
	menu_green_1.position.set( menu_clean_1.position.x, menu_clean_1.position.y);
	menu_green_1.alpha = 0;
	ladder_cont_1.addChild( menu_green_1 );

	let menu_ladder_1 = new Sprite( id_0["ladder_menu_1.png"] );
	menu_ladder_1.anchor.set(0.5, 0.5);
	menu_ladder_1.position.set( menu_clean_1.position.x, menu_clean_1.position.y);
	ladder_cont_1.addChild( menu_ladder_1 );


	let menu_clean_2 = new Sprite( id_0["menu_clean.png"] );
	menu_clean_2.anchor.set(0.5, 0.47);
	menu_clean_2.position.set( menu_clean_1.position.x + menu_clean_1.width, menu_clean_1.position.y);
	menu_clean_2.interactive = false;
	menu_clean_2.buttonMode = true;
	ladder_cont_2.addChild( menu_clean_2 );

	let menu_green_2 = new Sprite( id_0["menu_green.png"] );
	menu_green_2.anchor.set(0.5, 0.5);
	menu_green_2.position.set( menu_clean_2.position.x, menu_clean_2.position.y);
	menu_green_2.alpha = 0;
	ladder_cont_2.addChild( menu_green_2 );

	let menu_ladder_2 = new Sprite( id_0["ladder_menu_2.png"] );
	menu_ladder_2.anchor.set(0.5, 0.5);
	menu_ladder_2.position.set( menu_clean_2.position.x, menu_clean_2.position.y);
	ladder_cont_2.addChild( menu_ladder_2 );


	let menu_clean_3 = new Sprite( id_0["menu_clean.png"] );
	menu_clean_3.anchor.set(0.5, 0.47);
	menu_clean_3.position.set( menu_clean_2.position.x + menu_clean_1.width, menu_clean_1.position.y);
	menu_clean_3.interactive = false;
	menu_clean_3.buttonMode = true;
	ladder_cont_3.addChild( menu_clean_3 );

	let menu_green_3 = new Sprite( id_0["menu_green.png"] );
	menu_green_3.anchor.set(0.5, 0.5);
	menu_green_3.position.set( menu_clean_3.position.x, menu_clean_3.position.y);
	menu_green_3.alpha = 0;
	ladder_cont_3.addChild( menu_green_3 );

	let menu_ladder_3 = new Sprite( id_0["ladder_menu_3.png"] );
	menu_ladder_3.anchor.set(0.5, 0.5);
	menu_ladder_3.position.set( menu_clean_3.position.x, menu_clean_3.position.y);
	ladder_cont_3.addChild( menu_ladder_3 );

	//OK button

	let ok_btn = new Sprite( id_0["ok_btn.png"] );
	ok_btn.anchor.set(0.5, 0);
	ok_btn.position.set(0,-200);
	ok_btn.alpha = 0;
	ok_btn.interactive = true;
	ok_btn.buttonMode = true;
	ladder_menu.addChild( ok_btn );

	//Hammer button

	hammer = new Sprite( id_0["hammer.png"] );
	hammer.x = ladder.position.x - ladder.width / 2;
	hammer.y = ladder.position.y;
	hammer.anchor.set(0.5, 1);
	hammer.interactive = true;
	hammer.buttonMode = true;
	hammer.alpha = 0;
	gameScene.addChild( hammer );

	TweenLite.from(hammer.scale, 0.6, {x:0, y:0, ease:"easeOutBack", delay:1.5});
	TweenLite.to(hammer, 0.6, {alpha:1, ease:"easeOutBack", delay:1.5});
	TweenMax.to(hammer, 1.5, {y:hammer.position.y+10, repeat:-1, yoyo:true, ease:"easeInOutBack", repeatDelay:5});

	let final = new Container();
	final.position.set( app.stage.width / 2, app.stage.height / 2 );
	app.stage.addChild( final );

	let final_screen = new Sprite( id_0["final_screen.png"] );
	final_screen.anchor.set(0.5, 0.5);
	final_screen.position.set(0, -final_screen.height /3);
	final_screen.alpha = 0;
	final.addChild( final_screen );


	//Continue button

	let continue_btn = new Sprite( id_0["continue_btn.png"] );
	continue_btn.anchor.set(0.5, 0.5);
	continue_btn.position.set( 0, final.position.y / 2.5 );
	continue_btn.interactive = true;
	continue_btn.buttonMode = true;
	final.addChild( continue_btn );
	TweenMax.to(continue_btn.scale, 1, {x:1.1, y:1.1, repeat:-1, yoyo:true}); 


	//Other

	let austin = new Sprite( id_0["austin.png"] );
	austin.anchor.set(0, 1);
	austin.position.set( app.stage.width / 2, app.stage.height / 2 );
	gameScene.addChild( austin );

	let globus = new Sprite( id_0["globus.png"] );
	globus.position.set( 87, 109 );
	gameScene.addChild( globus );

	let logo = new Sprite( id_0["logo.png"] );
	logo.position.set( -app.stage.width / 2 + 32, -app.stage.height / 2 + 5 );
	final.addChild( logo );

	let plant_1 = new Sprite( id_0["plant_1.png"] );
	plant_1.position.set( 456, -42 );
	gameScene.addChild( plant_1 );

	let plant_2 = new Sprite( id_0["plant_2.png"] );
	plant_2.position.set( 1122, 438 );
	gameScene.addChild( plant_2 );

	let soffa = new Sprite( id_0["soffa.png"] );
	soffa.position.set( 127, 324 );
	gameScene.addChild( soffa );

	let table = new Sprite( id_0["table.png"] );
	table.position.set( 202, 196 );
	gameScene.addChild( table );


	//Events

	hammer.on('pointertap', () => {

		TweenLite.killTweensOf(hammer.scale);
		TweenLite.to(hammer.scale, 0.6, {x:0, y:0, alpha:0, ease:"easeInBack"});
		TweenLite.from(ladder_cont_1.scale, 0.6, {x:0, y:0, ease:"easeOutBack", delay:0.3});
		TweenLite.from(ladder_cont_2.scale, 0.6, {x:0, y:0, ease:"easeOutBack", delay:0.6});
		TweenLite.from(ladder_cont_3.scale, 0.6, {x:0, y:0, ease:"easeOutBack", delay:0.9});
		TweenLite.to(ladder_cont_1, 0.6, { alpha:1, ease:"easeInOutCubic", delay:0.3});
		TweenLite.to(ladder_cont_2, 0.6, { alpha:1, ease:"easeInOutCubic", delay:0.6});
		TweenLite.to(ladder_cont_3, 0.6, { alpha:1, ease:"easeInOutCubic", delay:0.9});
		menu_clean_1.interactive = menu_clean_2.interactive = menu_clean_3.interactive = true;

	});

	menu_clean_1.on('pointertap', () => {

		ok_btn.alpha = menu_green_2.alpha = menu_green_3.alpha = 0;
		ok_btn.position.set( menu_clean_1.position.x, menu_clean_1.position.y);
		TweenLite.killTweensOf(menu_green_1);
		TweenLite.to(menu_green_1, 0.1, { alpha:1, ease:"easeInCubic"});
		TweenLite.killTweensOf(ok_btn);
		TweenLite.to(ok_btn, 0.3, {y:menu_clean_1.height/3, alpha:1, ease:"easeOutElastic"});
		ladder_previous = ladder_active;
		ladder_active = 1;
		ladder_change(ladder_active, ladder_previous);

	});

	menu_clean_2.on('pointertap', () => {

		ok_btn.alpha = menu_green_1.alpha = menu_green_3.alpha = 0;
		ok_btn.position.set( menu_clean_2.position.x, menu_clean_2.position.y);
		TweenLite.killTweensOf(menu_green_2);
		TweenLite.to(menu_green_2, 0.1, { alpha:1, ease:"easeInCubic"});
		TweenLite.killTweensOf(ok_btn);
		TweenLite.to(ok_btn, 0.3, {y:menu_clean_2.height/3, alpha:1, ease:"easeOutElastic"});
		ladder_previous = ladder_active;
		ladder_active = 2;
		ladder_change(ladder_active, ladder_previous);	

	});

	menu_clean_3.on('pointertap', () => {

		ok_btn.alpha = menu_green_1.alpha = menu_green_2.alpha = 0;
		ok_btn.position.set( menu_clean_3.position.x, menu_clean_3.position.y);
		TweenLite.killTweensOf(menu_green_3);
		TweenLite.to(menu_green_3, 0.1, { alpha:1, ease:"easeInCubic"});
		TweenLite.killTweensOf(ok_btn);
		TweenLite.to(ok_btn, 0.3, {y:menu_clean_3.height/3, alpha:1, ease:"easeOutElastic"});
		ladder_previous = ladder_active;
		ladder_active = 3;
		ladder_change(ladder_active, ladder_previous);

	});

	ok_btn.on('pointertap', () => {
		menu_clean_1.interactive = menu_clean_2.interactive = menu_clean_3.interactive = false;
		ok_btn.interactive = false;
		TweenLite.killTweensOf(ok_btn);
		TweenLite.to(ok_btn.scale, 0.3, {x:0, y:0, ease:"easeOutCubic"});
		TweenLite.to(ok_btn, 0.3, {alpha:0, ease:"easeOutCubic"})
		TweenLite.killTweensOf(ladder_menu);
		TweenLite.to(ladder_menu, 0.6, {alpha:0, ease:"easeInOutCubic", delay:0.3});
		setTimeout(function (){
			finish();
		}, 1500);
		
	});

	continue_btn.on('pointertap', () => {

		alert('go to store with ladder: '+ ladder_active);

	});

function ladder_change(ladder_active, ladder_previous) {
	switch(ladder_previous) {
		case 1:
		TweenLite.killTweensOf(ladder_1);
		TweenLite.to(ladder_1, 0.4, {y:ladder_1.position.y+50, alpha:0, ease:"easeOutCubic"});
		break;
		case 2: 
		TweenLite.killTweensOf(ladder_2);
		TweenLite.to(ladder_2, 0.4, {y:ladder_2.position.y+50, alpha:0, ease:"easeOutCubic"});
		break;
		case 3:
		TweenLite.killTweensOf(ladder_3);
		TweenLite.to(ladder_3, 0.4, {y:ladder_3.position.y+50, alpha:0, ease:"easeOutCubic"}); 
		break;
		default:
		TweenLite.killTweensOf(ladder_0);
		TweenLite.to(ladder_0, 0.4, {y:ladder_0.position.y+50, alpha:0, ease:"easeOutCubic"});
	}
	switch(ladder_active) {
		case 1:
		ladder_1.position.set(0,-100);
		TweenLite.killTweensOf(ladder_1);
		TweenLite.to(ladder_1, 0.6, {y:ladder_1.position.y+100, alpha:1, ease:"easeOutCubic"});
		break;
		case 2: 
		ladder_2.position.set(0,-100);
		TweenLite.killTweensOf(ladder_2);
		TweenLite.to(ladder_2, 0.6, {y:ladder_2.position.y+100, alpha:1, ease:"easeOutCubic"});
		break;
		case 3:
		ladder_3.position.set(0,-100);
		TweenLite.killTweensOf(ladder_3);
		TweenLite.to(ladder_3, 0.6, {y:ladder_3.position.y+100, alpha:1, ease:"easeOutCubic"}); 
		break;
		default:
		ladder_0.position.set(0,-100);
		TweenLite.killTweensOf(ladder_0);
		TweenLite.to(ladder_0, 0.6, {y:ladder_0.position.y+100, alpha:1, ease:"easeOutCubic"});
	}
}

function finish () {
	let colorMatrix = new PIXI.filters.ColorMatrixFilter();
	gameScene.filters = [colorMatrix];
	colorMatrix.brightness(1);

	TweenLite.killTweensOf(final_screen);
	TweenLite.to(final_screen, 1, { alpha:1, ease:"easeInOutCubic", delay:0.3});

	let f_brightness = 1;
	let timerId = setInterval(function() {
		if (f_brightness <= 0.4) {clearInterval(timerId)}
	  f_brightness -= 0.05;
	  colorMatrix.brightness(f_brightness);
	}, 50);

}

}

