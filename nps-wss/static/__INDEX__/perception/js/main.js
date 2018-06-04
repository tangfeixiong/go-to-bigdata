/**
 * 主要是两个函数：
 * Myhole.start
 * Myhole.attack
 * Myhole.isStart
 * 
 * 请注意一下：必须在start完成之后才可以attack。
 * Myhole.start([team:'队名',avatar:'图标.jpg'...],function onsuccess);
 * Myhole.attack(from , to , time);
 * Myhole.isStart();
 * 
 * 
 * 本地开发，http://39.108.48.141/gj/index.html  由于跨域无法获得json数据。可以通过左边网页查看
 * 
 * 
 * */
var gdata ='';
var Main={
	init:function(){
		Main.initSocket();
		Main.initFinaltoplist();
		Main.initEchart1();
		Main.initEchart2();
		Main.initEchart3();
		Main.initEchart4();
	},
	initSocket:function(){
		//var url='ws://192.168.1.209:9001';
		//var url='ws://localhost:8080/sxdwg/socket';
		var idx = window.location.pathname.lastIndexOf('/');
		var prefix = window.location.pathname.substring(0, idx);
		var url= 'ws://' + window.location.host + prefix + '/ws';
		var ws = new ReconnectingWebSocket(url, null, {debug: true, reconnectInterval: 2000});
		//var ws = new WebSocket(url);
		
		ws.onopen=function(event){
		   if (window.console && window.console.log) {
              console.log(event);
           }
		
		   var myMap = {"1": {"target": {"ip_v4": "127.0.0.1", "port": 80}, "sources": [{"ip_v4": "127.0.0.1"}]}};
		   var myObj = {"stages": myMap};
		   ws.send(JSON.stringify(myObj));
		};
		
		ws.onmessage=function(event){
			var o = eval('(' + event.data + ')');
			if(o.finaltoplist){
				Main.renderfinaltoplist(o.finaltoplist);
				gdata = o.finaltoplist;
			}
			/*已经被取消
			 * if(o.currentoplist){
				Main.rendercurrentoplist(o.currentoplist);
			}*/
			if(o.attackdata){
				Main.renderattack(o.attackdata);
			}
			if(o.radardata){
				Main.renderradardata(o.radardata);
			}
			if(o.piedata){
				Main.renderpiedata(o.piedata);
			}
			if(o.bardata){
				Main.renderbardata(o.bardata);
				
				var obj = o.bardata;
				var objtop5 = '';
				var change = '';
				var seriesinfo='';
				var num;
				var bardata;
				if( gdata != '' ){
					gdata.sort(function(a,b){return b.score-a.score;});
					num ='';
					bardata='';
					for(var i=0;i<gdata.length;i++){
							if(i>4) break;
							objtop5 += '"' + gdata[i].teamname + '",';
							$.each(Config.teams,function(j,k){
								if( gdata[i].teamname == k.team ){
									num = k.number;
									return false;
								}
							});
							
							$.each(obj.teams,function(m,n){
								if( gdata[i].teamname == n ){
									bardata = obj.number[m];
									return false;
								}
							});
							
							seriesinfo += '{name:"' +gdata[i].teamname + '",value:[' + gdata[i].score + ',' + gdata[i].add_score + ',' + gdata[i].del_score + ',' 
								 bardata + ',' + num + ']},'

							
					}
					change = eval("["+objtop5+"]");
					seriesinfo = eval("["+seriesinfo+"]");
					var	option = {
						legend: {
							data: change,
						},
						series: [{
							data:seriesinfo
						}]
					};
					Main.renderradardata(option);
				}

			}
			if(o.caution){
				Main.rendercaution(o.caution);
			}
			if(o.barTTeamdata){
				Main.renderbarTTeamdata(o.barTTeamdata);
			}

			if(o.finaltoplist[0]['dispart']==0){
				$(".defen_time").html("比赛已经结束");
			}else if($("#t_s").html()=="中"){
				var s = o.finaltoplist[0]['dispart'];
	    		time();  
			    function time(){  
			        console.log(formatSeconds(s));  
			        if (s == 0) {  
			        	$(".defen_time").html("比赛已经结束");
			        }else{  
			            s--;  
			            setTimeout(function() {  
			                time();  
			            },  
			            1000)  
			        }  
			    }

				function formatSeconds(value) {  
			        var theTime = parseInt(value);// 秒  
			        var theTime1 = 0;// 分  
			        var theTime2 = 0;// 小时  
			        if(theTime > 60) {  
			            theTime1 = parseInt(theTime/60);  
			            theTime = parseInt(theTime%60);  
			            if(theTime1 > 60) {  
			                theTime2 = parseInt(theTime1/60);  
			                theTime1 = parseInt(theTime1%60);  
			            }  
			        }  
			        document.getElementById("t_h").innerHTML = parseInt(theTime2) + "：";
			        document.getElementById("t_i").innerHTML = parseInt(theTime1) + "：";
			        document.getElementById("t_s").innerHTML = parseInt(theTime);
			    }    
			}

		}
	},
	rendercaution:function(data){
		var str = '<img src="img/513.png">';
		for(var i=0;i<data.length;i++){
			str += '警告:' + data[i].attack_type + ':' + data[i].from + '违规使用' + data[i].attack_name + '攻击' + data[i].to +',请立即停止非法攻击操作!';
		}
		$("#scroll_begin").html(str);
	},
	renderfinaltoplist:function(data){
		$('#finaltoplist').html('');
		for(var i=0;i<data.length;i++){
			$.each(Config.teams,function(j,k){
				if(data[i].teamname == k.team){
					data[i].avatar = k.avatar;
					return false;
				}
			});
			var html=Main.getTemplateHtml(i,data[i]);
			$('#finaltoplist').append(html);
		}
	},
	rendercurrentoplist:function(data){
		$('#currentoplist').html('');
		for(var i=0;i<data.length;i++){
			var html=Main.getTemplateHtml(i,data[i]);
			$('#currentoplist').append(html);
		}
	},
	renderattack:function(data){
		var strli='';
		var str='';
		var oppie='';
		if(!Myhole.isStart()){
			//初始化
			Myhole.start(Config.teams,function(){
				$.each(data,function(i,v){
					Myhole.attack(v.from,v.to,v.time);
					Main.appendlog(v.from,v.to,v.time);

					strli += Main.renderrattackpacket(v);
				});
			});
		}else{
			$.each(data,function(i,v){
				Myhole.attack(v.from,v.to,v.time);
				Main.appendlog(v.from,v.to,v.time);
				
				strli += Main.renderrattackpacket(v);				
			});
		}
		$("#attack_con_main").append(strli);
		Config.teams.sort(function(a,b){return b.number-a.number;});
		$.each(Config.teams,function(j,k){	
			if(j < 10 && k.number >0 ){
				str+='<ul class="attack_cout_data clear">'+
					'<li>'+k.number+'</li>'+
					'<li><img src="'+k.avatar+'">'+k.team+'</li>'+
			    		'</ul>';	
				oppie += '{value:' + k.number + ', name:"' + k.team + '"},';
			}
		});
		$(".attack_cout_main").html(str);
		
		//更新饼图
		var obj = eval("["+oppie+"]");
		var option = {
			   series: [
			        {
			            data:obj
			        }
			    ]
		};
		Main.renderpiedata(option);
	},
	formatDuring:function(mss) {
		    var hours = parseInt((mss % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))+8;
		    var minutes = parseInt((mss % (1000 * 60 * 60)) / (1000 * 60));   
		    var seconds = (mss % (1000 * 60)) / 1000;
		    seconds=seconds.toFixed(3);
		    if (hours<10) {hours="0"+hours;};
		    if (minutes<10) {minutes="0"+minutes;};
		    if (seconds<10) {seconds="0"+seconds;};
		    return  hours + ":" + minutes + ":" + seconds;
	}, 
	renderrattackpacket:function(v){	
		var strli='';
		var unixTimestamp = Main.formatDuring(v.time) ;
		$.each(Config.teams,function(j,k){
			if(v.from == k.team){
				k.number++;
			}
		});
		strli='<ul class="attack_con clear">'+
			'<li>' + unixTimestamp + '</li>'+
			'<li>' + v.from + '</li>'+
			'<li>' + v.fromip + '</li>'+
			'<li>' + v.to + '</li>'+
			'<li>' + v.toip + '</li>'+
			'<li>' + v.protocol + '</li>'+
			'<li>' + v.port+'</li>' +
			'<li>' + v.DataSize + 'KB</li>'+
			'</ul>';
		return strli;
	},
	renderradardata:function(data){
		/*{
			legend:{data: ['fox队', '大虾队', '大圣队', '高手高手高手队', '名字很长很长的队']},
			radar:{indicator:[{name:'总分'},{name:'得分'},{name:'失分'},{name: '数据包大小'},{name:'频率'}]},
			series:[{data:[{name:'大虾队',value : [1,3,5,7,9]},{name:'fox队',value:[2, 12, 14, 18,42]},{name:'大圣队',value:[3, 12, 14,3,41]},{name:'高手高手高手队',value:[3, 12, 14,4,41]},{name:'名字很长很长的队',value:[3, 12, 14,5,41]}]}]
		}*/
		Main.chart1.setOption(data);
	},
	renderpiedata:function(data){
		/*{
			series:[{data:[{name:'大虾队',value : [1,3,5,7,9]},{name:'fox队',value:[2, 12, 14, 18,42]},{name:'大圣队',value:[3, 12, 14,3,41]},{name:'高手高手高手队',value:[3, 12, 14,4,41]},{name:'名字很长很长的队',value:[3, 12, 14,5,41]}]}]
		}*/
		Main.chart2.setOption(data);
	},
	renderbardata:function(data){
		/*{
			yAxis:{data:['队伍一','队伍二','队伍三','队伍四','队伍5','队伍6','队伍7','队伍8','队伍9','队伍0']},
			series:[{data:[1,3,5,7,9 ,2,4,6,8,10]}]
		}*/
		var objteams=Array();
		var objnumber=Array();
		$.each(data.teams,function(j,k){
				if(j<10){
					objteams[j]=k;
				}
		});
		$.each(data.number,function(j,k){
				if(j<10){
					objnumber[j]=k;
				}
		});

		var option = {
			 yAxis: {
			        data: objteams
			    },
				 series: [
			        {
			            type: 'bar',
			            itemStyle:{
			            	normal:{
			            		opacity:0.6
			            	}
			            },
			            data: objnumber
			        },
			    ]
			
		};
		Main.chart3.setOption(option);
	},
	getTemplateHtml:function(i,obj){
		switch(i){
			case 0:{
				return '<ul class="Score_ul score_one">'+
				'<li class="Score_li"><img src="img/defenimg/jiangbei.png" /></li>'+
				'<li class="Score_li"><img src="'+obj.avatar+'" /></li>'+
				'<li class="Score_li score_di">'+obj.teamname+'</li>'+
				'<li class="Score_li score_font">'+obj.score+'</li>'+
				'<li class="Score_li score_font">'+obj.add_score+'</li>'+
				'<li class="Score_li score_font">'+obj.del_score+'</li>'+
				'<div class="defen_click"></div>'+
			'</ul>';
			};
			case 1:{
				return '<ul class="Score_ul score_tow">'+
				'<li class="Score_li"><img src="img/defenimg/yinpai.png" /></li>'+
				'<li class="Score_li"><img src="'+obj.avatar+'" /></li>'+
				'<li class="Score_li score_di">'+obj.teamname+'</li>'+
				'<li class="Score_li score_font">'+obj.score+'</li>'+
				'<li class="Score_li score_font">'+obj.add_score+'</li>'+
				'<li class="Score_li score_font">'+obj.del_score+'</li>'+
				'<div class="defen_click"></div>'+
			'</ul>';
			};
			case 2:{
				return '<ul class="Score_ul score_three">'+
				'<li class="Score_li"><img src="img/defenimg/tongpai.png" /></li>'+
				'<li class="Score_li"><img src="'+obj.avatar+'" /></li>'+
				'<li class="Score_li score_di">'+obj.teamname+'</li>'+
				'<li class="Score_li score_font">'+obj.score+'</li>'+
				'<li class="Score_li score_font">'+obj.add_score+'</li>'+
				'<li class="Score_li score_font">'+obj.del_score+'</li>'+
				'<div class="defen_click"></div>'+
			'</ul>';
			};
			default:{
				return '<ul class="Score_ul">'+
				'<li class="Score_li">'+Main.pad(i+1,3)+'</li>'+
				'<li class="Score_li"><img src="'+obj.avatar+'" /></li>'+
				'<li class="Score_li score_di">'+obj.teamname+'</li>'+
				'<li class="Score_li score_font">'+obj.score+'</li>'+
				'<li class="Score_li score_font">'+obj.add_score+'</li>'+
				'<li class="Score_li score_font">'+obj.del_score+'</li>'+
				'<div class="defen_click"></div>'+
			'</ul>';
			}
		}
	},
	/* 质朴长存法*/  
	pad:function(num, n) {  
	    var len = num.toString().length;  
	    while(len < n) {  
	        num = "0" + num;  
	        len++;  
	    }  
	    return num;  
	},
	appendlog:function(from,to,time){
		$(".attacklog:gt(50)").remove();
		$('#attacklog').prepend('<ul class="attack_ul attacklog">'+
				'<li class="attack_li">'+from+'</li>'+
				'<li class="attack_li">'+to+'</li>'+
				'<li class="attack_li">'+moment(time,'x').format('HH:mm:ss')+'</li>'+
				'<div class="defen_click"></div>'+
			'</ul>');
	},
	initFinaltoplist:function(){
		$('#finaltoplist').html('');
		$.each(Config.teams,function(j,k){
			var teamlist = 	'{"avatar":"' + k.avatar + '","teamname":"' + k.team + '","score":0,"add_score":0,"del_score":0},';
			var obj = eval("["+teamlist+"]"); 
			var html=Main.getTemplateHtml(j,obj[0]);
			$('#finaltoplist').append(html);
		});
	},
	chart1:null,
	initEchart1:function(){
		var chart1 = echarts.init(document.getElementById('chart1'));
		Main.chart1=chart1;
		var teams = '';
		$.each(Config.teams,function(j,k){
			if(j >4){ return false;}
			teams += "'"+ k.team + "'" + ',';
		});
		var obj = eval("["+teams+"]");
		
		option = {
				color: ['#2d398f','#5bc0de','#fbb450','#62c462','#6495ed',
		            '#ff69b4','#ba55d3','#cd5c5c','#ffa500','#40e0d0',
		            '#1e90ff','#ff6347','#7b68ee','#00fa9a','#ffd700',
		            '#6699FF','#ff6666','#3cb371','#b8860b','#30e0e0'],
				tooltip: {},
				legend: {
					data: obj,
					right:'right',
			        textStyle: {
			            color: '#AAA'
			        },
			        formatter: function (name) {//大于3个字的队伍自动拆检
			        	if(name.length>3){
			        		return name.substring(0,3);
			        	}else{
			        		return name;	
			        	}
			        }
			    },
				radar: {
					center:['40%','50%'],
					/*radius: 50,*/
			        shape: 'circle',
			        splitLine: {
			            lineStyle: {
			                color: [
			                    'rgba(238, 197, 102, 0.1)', 'rgba(238, 197, 102, 0.2)',
			                    'rgba(238, 197, 102, 0.4)', 'rgba(238, 197, 102, 0.6)',
			                    'rgba(238, 197, 102, 0.8)', 'rgba(238, 197, 102, 1)'
			                ].reverse()
			            }
			        },
			        splitArea: {
			            show: false
			        },
			        indicator: [
			           { name: 'Score'},
			           { name: 'Attack'},
			           { name: 'Lose'},
			           { name: 'Defense'},
			           { name: 'Rate'}
			        ]
			    },
			    series: [{
			        type: 'radar',
			        lineStyle:{
		            	normal:{
		            		opacity:0.6
		            	}
		            },
			        data : [{name:'大虾队',value : [10, 12, 14, 18]},{name:'fox队',value:[15, 12, 14, 18]}]
			        // data:[]
			    }]
			};
		chart1.setOption(option);
		//test data;  可以删
		//var iii=1;
		//setInterval(function(){
		//	chart1.setOption({
		//		series:[{data:[{name:'大虾队',value : [10, iii++, 14, 18,12]},{name:'fox队',value:[iii, 12, 14, 18,42]},{name:'大圣队',value:[3, 12, 14,iii-5,41]},{name:'高手高手高手队',value:[3, 12, 14,iii-5,41]},{name:'名字很长很长的队',value:[3, 12, 14,iii-5,41]}]}]
		//	});
		//},1000);
	},
	chart2:null,
	initEchart2:function(){
		var chart2 = echarts.init(document.getElementById('chart2'));
		Main.chart2=chart2;
		option = {
				color: ['#2d398f','#5bc0de','#fbb450','#62c462','#6495ed',
		            '#ff69b4','#ba55d3','#cd5c5c','#ffa500','#40e0d0',
		            '#1e90ff','#ff6347','#7b68ee','#00fa9a','#ffd700',
		            '#6699FF','#ff6666','#3cb371','#b8860b','#30e0e0'],
			    tooltip: {
			        trigger: 'item',
			        formatter: "{a} <br/>{b}: {c} ({d}%)"
			    },
			    series: [
			        {
			            name:'攻击频率',
			            type:'pie',
			            radius: ['50%', '80%'],
			            avoidLabelOverlap: false,
			            z:1000,
			            itemStyle:{
			            	normal:{
			            		opacity:0.6
			            	}
			            },
			            data:[{value:335, name:'大虾队'},{value:310, name:'fox队'},{value:234, name:'广告队'}]
			            // data:[]
			        }
			    ]
			};
		chart2.setOption(option);
		//test data;  可以删
		//var iii=10;
		//setInterval(function(){
		//	iii++;
		//	chart2.setOption({
		//		series:[{data:[{value:iii, name:'大虾队'},{value:10, name:'fox队'},{value:30, name:'大圣队'}]}]
		//	});
		//},1000);
	},
	chart3:null,
	initEchart3:function(){
		//让chart3 自适应当前浏览器宽度
		// var aw=$(window).width();
		// var lw=$('#finaltoplist').width();
		// var rw=$('#chart1panel').width();
		// $('#chart3panel').css('width',1525);
		// $('#chart3panel').css('right',20);
		
		var chart3 = echarts.init(document.getElementById('chart3'));
		Main.chart3=chart3;
		var teams = '';
		$.each(Config.teams,function(j,k){
			if(j >9){ return false;}
			teams += "'"+ k.team + "'" + ',';
		});
		var obj = eval("["+teams+"]");
		option = {
			    tooltip: {
			        trigger: 'axis',
			        axisPointer: {
			            type: 'shadow'
			        }
			    },
			    grid: {
			    	top:5,
			        left:0,
			        right:0,
			        bottom:0,
			        containLabel: true
			    },
			    xAxis: {
			        type: 'value',
			        splitLine: {show:false},
			        axisLabel:{
			        	show:false,
			            textStyle: {
				            color: '#AAA'
				        },
			        },
			    },
			    yAxis: {
			        type: 'category',
			        axisLabel:{
			            formatter:function(name){//大于3个字的队伍自动拆检
			            	if(name.length>3){
				        		return name.substring(0,3);
				        	}else{
				        		return name;	
				        	}
			            },
			            textStyle: {
				            color: '#AAA'
				        },
			        },
			        data: obj
			    },
			    series: [
			        {
			            type: 'bar',
			            itemStyle:{
			            	normal:{
			            		opacity:0.6
			            	}
			            },
			            // data: []
			            data:[{value:335, name:'大虾队'},{value:310, name:'fox队'},{value:234, name:'广告队'},{value:335, name:'大虾队'},{value:310, name:'fox队'},{value:234, name:'广告队'},{value:335, name:'大虾队'},{value:310, name:'fox队'},{value:234, name:'广告队'}]
			        },
			      
			    ]
			};
		chart3.setOption(option);
		//test data;  可以删
		//var iii=1;
		//setInterval(function(){
		//	iii++;
		//	chart3.setOption({
		//		series:[{data:[1,iii,3,4,5,6,5,4,iii+2,2]}]
		//	});
		//},1000);
		
	},


	renderbarTTeamdata:function(data){
		/*{
			series:[{data:[{name:'大虾队',value : [1,3,5,7,9]},{name:'fox队',value:[2, 12, 14, 18,42]},{name:'大圣队',value:[3, 12, 14,3,41]},{name:'高手高手高手队',value:[3, 12, 14,4,41]},{name:'名字很长很长的队',value:[3, 12, 14,5,41]}]}]
		}*/
		var option = {
			 yAxis: {
			        data: data.teams
			    },
				 series: [
			        {
			            type: 'bar',
			            itemStyle:{
			            	normal:{
			            		opacity:0.6
			            	}
			            },
			            data: data.number
			        },
			    ]
			
		};
		Main.chart4.setOption(data);
	},


	//柱状图
	chart4:null,
	initEchart4:function(){		
		var chart4 = echarts.init(document.getElementById('chart4'));
		Main.chart4=chart4;
		var teams = '';
		$.each(Config.teams,function(j,k){
			if(j >9){ return false;}
			teams += "'"+ k.team + "'" + ',';
		});
		var obj = eval("["+teams+"]");
		option = {
			    tooltip: {
			        trigger: 'axis',
			        // axisPointer: {
			        //     type: 'shadow'
			        // }
			    },
    			toolbox: {
			        show : true,
			        feature : {
			            // mark : {show: true},
			            // dataView : {show: true, readOnly: false},
			            // magicType : {show: true, type: ['line', 'bar']},
			        }
			    },
			    //曲线在容器中的位置
			    grid: {
			    	top:40,
			        left:20,
			        right:20,
			        bottom:20,
			        containLabel: true
			    },
			    yAxis: {
			        type: 'value',
			        splitLine: {show:false},
			        axisLabel:{
			        	// show:false,
			            textStyle: {
				            color: '#fff'
				        },

			        },
			        axisLine:{lineStyle:{ color:'#006bb0',width:'2'}},//纵坐标的颜色
			    },
			    xAxis: {
			        type: 'category',
			        // boundaryGap : false,
			        axisLabel:{
			            formatter:function(name){//大于3个字的队伍自动拆检
			            	if(name.length>3){
				        		return name.substring(0,3);
				        	}else{
				        		return name;	
				        	}
			            },
			            //设置横坐标文字倾斜
			            // interval:0,  
               //          rotate:-60,//倾斜度 -90 至 90 默认为0  
               //          margin:2, 
			            textStyle: {
				            color: '#fff',
				            fontSize: '12',
				            fontWeight:'bolder',
				        },

			        },
			        axisLine:{lineStyle:{ color:'#006bb0',width:'2'}},//横坐标的颜色
			        data: ['总分','得分','失分']
			    },
			    series: [
			        {
			            type: 'bar',
			            // itemStyle:{
			            // 	normal:{
			            // 		opacity:0.6
			            // 	}
			            // },
			            // 顶部展示数字即折线点和折线条的样式
			            itemStyle: {
			            	normal:{ 
			            		color: function (value){ return "#"+("00000"+((Math.random()*16777215+0.5)>>0).toString(16)).slice(-6); },
			            		areaStyle: {
			            			type: 'default',
			            			// color:'#6dc0d8'
			            		},
			            		label: {  
                                	show: true,//是否展示  
                                	textStyle: {  
	                                    fontWeight:'bolder',  
	                                    fontSize : '12',  
	                                    fontFamily : '微软雅黑',
	                                    color:'#fff'  
                                	}
                                },
                                lineStyle:{
                                	// color:'#6dc0d8'
                                }  
                         	}   

			        	},
			        	data: [600,410,903,487,505,620,25,444,112,232]
			        },
			      
			    ]
			};
		chart4.setOption(option);
		//test data;  可以删
		// var iii=1;
		// setInterval(function(){
		// 	iii++;
		// 	chart3.setOption({
		// 		series:[{data:[1,iii,3,4,5,6,5,4,iii+2,2]}]
		// 	});
		// },1000);	
	},
}