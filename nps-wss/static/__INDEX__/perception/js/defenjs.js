$(document).ready(function(){
	//攻击事件切换
	$(".attack_span").eq(0).click(function(){
		$(this).addClass("active_tile").siblings().removeClass("active_tile");
		$(".now_attack").fadeIn(200);
		$(".history_attack").fadeOut(100);
	});
	$(".attack_span").eq(1).click(function(){
		$(this).addClass("active_tile").siblings().removeClass("active_tile");
		$(".history_attack").fadeIn(200);
		$(".now_attack").fadeOut(100);
	});
	
	
});













