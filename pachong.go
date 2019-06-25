package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

//var htmltext string = `

<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="trafficjunky-site-verification" content="445w343y8" />
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta http-equiv="x-dns-prefetch-control" content="on">
<meta name="renderer" content="webkit">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>若菜奈央 - 演员 - 影片 - AVMOO</title>
<meta name="author" content="AVMOO">
<meta name="keywords" content="若菜奈央,演员,热门,AVMOO">
<meta name="description" content="若菜奈央 - 演员 - 影片 - AVMOO - 你的线上日本成人影片情报站。管理你的影片并分享你的想法。">
<link rel="apple-touch-icon" href="https://avmoo.asia/app/jav/View/img/apple-touch-icon.png">
<link rel="shortcut Icon" href="https://avmoo.asia/app/jav/View/img/favicon.ico">
<link rel="bookmark" href="https://avmoo.asia/app/jav/View/img/favicon.ico">
<link rel="dns-prefetch" href="https://jp.netcdn.space" />
<link rel="dns-prefetch" href="https://us.netcdn.space" />
<link rel="dns-prefetch" href="https://ads.exoclick.com" />
<link rel="dns-prefetch" href="https://syndication.exoclick.com" />
<link rel="dns-prefetch" href="https://adserver.juicyads.com" />
<link rel="dns-prefetch" href="https://j.traffichunt.com" />
<script language="javascript">function $ROOT_URL(){return "https://avmoo.asia"}function $APP(){return "jav"}function $APP_URL(){return "/index.php?app=jav&"}function $APP_INFO_URL(){return "/index.php/jav"}function $APP_REWRITE_URL(){return "/jav"}function $APP_VIEW_URL(){return "https://avmoo.asia/app/jav/View"}function $APP_UPLOAD_URL(){return "/app/jav/Upload"}</script><link rel='stylesheet' type='text/css' href='https://avmoo.asia/app/jav/View/css/app.min.css?v=1476953808'>
<!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries --><!--[if lt IE 9]><script src='https://avmoo.asia/script/html5shiv/3.7.2/html5shiv.min.js?v=1476953808'></script>
<script src='https://avmoo.asia/script/respond/1.4.2/respond.min.js?v=1476953808'></script>
<![endif]--><!--[if lt IE 8]><link rel='stylesheet' type='text/css' href='https://avmoo.asia/script/bootstrap.ie7/1.0/bootstrap.ie7.min.css?v=1476953808' >
<![endif]--><script>
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
ga('create', 'UA-74041965-4', {'sampleRate': 50});
ga('send', 'pageview');
</script>
</head>

  <body>
    <nav class="navbar navbar-default navbar-fixed-top top-bar">
      <div class="container-fluid">
        <div class="navbar-header">
        	<a href="https://avmoo.asia/cn" class="logo"></a>
            <div class="btn-group pull-right visible-xs-inline" role="group" style="margin-top:8px;margin-right:8px;">
                <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
                  <span class="glyphicon glyphicon-globe"></span> 简体中文                  <span class="caret"></span>
                </button>
                <ul class="dropdown-menu" role="menu">
                    <li><a href="https://avmoo.asia/en/star/ozw">English</a></li>
                    <li><a href="https://avmoo.asia/ja/star/ozw">日本语</a></li>
                    <li><a href="https://avmoo.asia/tw/star/ozw">正體中文</a></li>
                    <li><a href="https://avmoo.asia/cn/star/ozw">简体中文</a></li>
                </ul>
            </div>
        </div>

        <div id="navbar" class="collapse navbar-collapse">
          <form class="navbar-form navbar-left fullsearch-form" action="https://avmoo.asia/cn/search" onsubmit="return false">
            <div class="input-group">
              <input name="keyword" type="text" class="form-control" placeholder="搜寻 识别码, 影片, 演员">
              <span class="input-group-btn">
                <button class="btn btn-default" type="submit">搜寻</button>
              </span>
            </div>
          </form>
          <ul class="nav navbar-nav">
            <li ><a href="https://avmoo.asia/cn">全部</a></li>
			<li ><a href="https://avmoo.asia/cn/released">已发布</a></li>
			<li ><a href="https://avmoo.asia/cn/popular">热门</a></li>
            <li ><a href="https://avmoo.asia/cn/actresses">女优</a></li>
            <li ><a href="https://avmoo.asia/cn/genre">类别</a></li>
          </ul>
          <ul class="nav navbar-nav navbar-right">
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><span class="glyphicon glyphicon-globe" style="font-size:12px;"></span> <span class="hidden-sm">简体中文</span> <span class="caret"></span></a>
              <ul class="dropdown-menu" role="menu">
                <li><a href="https://avmoo.asia/en/star/ozw">English</a></li>
                <li><a href="https://avmoo.asia/ja/star/ozw">日本语</a></li>
                <li><a href="https://avmoo.asia/tw/star/ozw">正體中文</a></li>
                <li><a href="https://avmoo.asia/cn/star/ozw">简体中文</a></li>
              </ul>
            </li>
          </ul>
        </div><!--/.nav-collapse -->
      </div>
    </nav>
<div class="row visible-xs-inline footer-bar">
    <div class="col-xs-3 text-center">
        <a id="menu" class="btn btn-default trigger-overlay"><i class="glyphicon glyphicon-align-justify"></i></a>
    </div>
        <div class="col-xs-3 text-center">
            </div>
    <div class="col-xs-3 text-center">
                <a id="prev" class="btn btn-default" href="/cn/star/ozw/page/2" style="display:none"><i class="glyphicon glyphicon-chevron-right"></i></a>
            </div>
        <div class="col-xs-3 text-center">
        <a id="back" class="btn btn-default" href="javascript:window.history.back()"><i class="glyphicon glyphicon-share-alt flipx"></i></a>
    </div>
</div>
    <div class="container-fluid">
                <div class="row">
			<div id="waterfall">
						                <div class="item">
                    <div class="avatar-box">
                        <div class="photo-frame">
                            <img src="https://jp.netcdn.space/mono/actjpgs/wakana_nao.jpg" title="">
                        </div>
						<div class="photo-info">
                            <span class="pb-10">若菜奈央</span>
																					<p>身高: 171cm</p>														<p>胸围: 88cm</p>							<p>腰围: 59cm</p>							<p>臀围: 88cm</p>							<p>出生地: 東京都</p>							<p>爱好: ヨガ、水泳、ピアノ</p>							<p><a href="https://avsox.asia/cn/search/%E8%8B%A5%E8%8F%9C%E5%A5%88%E5%A4%AE" target="_blank" style="color:#CC0000;">更多无码影片</a></p>
						</div>
					</div>
                </div>
													<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/7779">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/hndb00143/hndb00143ps.jpg" title="嬉しそうにニヤニヤ僕を見下ろしてからかい腰振り騎乗位で全力中出ししてくる痴女お姉さんBEST4時間">
                        </div>
                        <div class="photo-info">
	                       <span>嬉しそうにニヤニヤ僕を見下ろしてからかい腰振り騎乗位で全力中出ししてくる痴女お姉さんBEST4時間 <br><date>HNDB-143</date> / <date>2019-07-20</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/76wh">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/5642hodv21396/5642hodv21396ps.jpg" title="魅惑のセクシーランジェリー 恥ずかしいけどこんな姿の私を見て。あなたとセックスしたいから。似合う？興奮してくれる？…ねぇ、抱いて。">
                        </div>
                        <div class="photo-info">
	                       <span>魅惑のセクシーランジェリー 恥ずかしいけどこんな姿の私を見て。あなたとセックスしたいから。似合う？興奮してくれる？…ねぇ、抱いて。 <br><date>HODV-21396</date> / <date>2019-07-05</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/75fs">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/dazd00095/dazd00095ps.jpg" title="圧倒的存在感 揺れる波打つ至高のデカ尻7時間BEST">
                        </div>
                        <div class="photo-info">
	                       <span>圧倒的存在感 揺れる波打つ至高のデカ尻7時間BEST <br><date>DAZD-095</date> / <date>2019-06-22</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/76h6">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/13rvg00097/13rvg00097ps.jpg" title="禁断介護BEST vol.14">
                        </div>
                        <div class="photo-info">
	                       <span>禁断介護BEST vol.14 <br><date>RVG-097</date> / <date>2019-06-19</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/74s3">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/dvaj00398/dvaj00398ps.jpg" title="中出ししたザーメンをチ○ポで膣内シェイクする怒涛の追撃ピストン性交">
                        </div>
                        <div class="photo-info">
	                       <span>中出ししたザーメンをチ○ポで膣内シェイクする怒涛の追撃ピストン性交  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>DVAJ-398</date> / <date>2019-06-09</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/75sv">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/td020dvaj00268/td020dvaj00268ps.jpg" title="【お得】出会った瞬間合体！会ったら即ハメ、とにかく合体＆無許可中出し。打合せ中に男優忍び寄り生ハメ。驚く彼女をよそに立ちバックで突きまくる。椅子に座って背面騎乗、床に座って対面座位。立ちバックからテーブル使って正常位。気持ち良すぎて暴発中出し。 若菜奈央">
                        </div>
                        <div class="photo-info">
	                       <span>【お得】出会った瞬間合体！会ったら即ハメ、とにかく合体＆無許可中出し。打合せ中に男優忍び寄り生ハメ。驚く彼女をよそに立ちバックで突きまくる。椅子に座って背面騎乗、床に座って対面座位。立ちバックからテーブル使って正常位。気持ち良すぎて暴発中出し。 若菜奈央  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>DVAJ-268</date> / <date>2019-05-31</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/74x4">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/2wsp00165/2wsp00165ps.jpg" title="エステで火照る爆乳妻">
                        </div>
                        <div class="photo-info">
	                       <span>エステで火照る爆乳妻 <br><date>WSP-165</date> / <date>2019-05-24</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/74ve">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/1ienf00008/1ienf00008ps.jpg" title="淫語かたりかけ自画撮りぐちゅぐちゅ連続絶頂指オナニー4">
                        </div>
                        <div class="photo-info">
	                       <span>淫語かたりかけ自画撮りぐちゅぐちゅ連続絶頂指オナニー4  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>IENF-008</date> / <date>2019-05-23</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/743e">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/td019dvaj00279/td019dvaj00279ps.jpg" title="【特価】巨乳でスーパーボディ奈央ちゃんに中出し！いきなりパンツ下してバックで生ハメ。椅子を利用したバックでチ○ポとマ○コの高さばっちり奥突き。椅子に座って背面座位で動いて貰い片足上げバックでイク。場所を移して再度バックで突いて、奈央ちゃんマ○コに中出し！">
                        </div>
                        <div class="photo-info">
	                       <span>【特価】巨乳でスーパーボディ奈央ちゃんに中出し！いきなりパンツ下してバックで生ハメ。椅子を利用したバックでチ○ポとマ○コの高さばっちり奥突き。椅子に座って背面座位で動いて貰い片足上げバックでイク。場所を移して再度バックで突いて、奈央ちゃんマ○コに中出し！  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>DVAJ-279</date> / <date>2019-04-26</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/71qg">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/kibd00240/kibd00240ps.jpg" title="生意気ギャルの恥ずかしいケツ穴丸見え！ブンブン杭打ち騎乗位BEST">
                        </div>
                        <div class="photo-info">
	                       <span>生意気ギャルの恥ずかしいケツ穴丸見え！ブンブン杭打ち騎乗位BEST  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>KIBD-240</date> / <date>2019-04-13</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/71hz">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/mkck00232/mkck00232ps.jpg" title="絶品ボディ膣内に大量射精！真正中出し134発8時間">
                        </div>
                        <div class="photo-info">
	                       <span>絶品ボディ膣内に大量射精！真正中出し134発8時間 <br><date>MKCK-232</date> / <date>2019-04-07</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/72ei">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/h_237nacx00028/h_237nacx00028ps.jpg" title="くねらす絶品ボディ中出し14人">
                        </div>
                        <div class="photo-info">
	                       <span>くねらす絶品ボディ中出し14人  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>NACX-028</date> / <date>2019-04-01</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/708d">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/kibd00239/kibd00239ps.jpg" title="生意気ギャルをイカせてもやめない容赦ないハードピストン！「さっきからイッてんだよ」状態で連続イキ追撃ピストンBEST">
                        </div>
                        <div class="photo-info">
	                       <span>生意気ギャルをイカせてもやめない容赦ないハードピストン！「さっきからイッてんだよ」状態で連続イキ追撃ピストンBEST  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>KIBD-239</date> / <date>2019-03-16</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/71ms">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/td018dvaj00310/td018dvaj00310ps.jpg" title="【特価】Fカップ巨乳奈央ちゃんの調教ハードセックス。調教仕上げはバックから突かれ感謝の言葉と共にイクと背面騎乗での突き上げでもイカされる。バックに戻っての激ピスでお尻に発射されると、2人目はおっぱいを弄りながら正常位で発射。3人目も正常位のままで口に発射！">
                        </div>
                        <div class="photo-info">
	                       <span>【特価】Fカップ巨乳奈央ちゃんの調教ハードセックス。調教仕上げはバックから突かれ感謝の言葉と共にイクと背面騎乗での突き上げでもイカされる。バックに戻っての激ピスでお尻に発射されると、2人目はおっぱいを弄りながら正常位で発射。3人目も正常位のままで口に発射！  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>DVAJ-310</date> / <date>2019-03-15</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/7016">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/dvaj00382/dvaj00382ps.jpg" title="マゾメス奴隷調教BEST15人5時間">
                        </div>
                        <div class="photo-info">
	                       <span>マゾメス奴隷調教BEST15人5時間  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>DVAJ-382</date> / <date>2019-03-11</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/7004">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/mbyd00286/mbyd00286ps.jpg" title="私、実は夫の上司に犯され続けてます… 総集編vol.5">
                        </div>
                        <div class="photo-info">
	                       <span>私、実は夫の上司に犯され続けてます… 総集編vol.5 <br><date>MBYD-286</date> / <date>2019-03-10</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/70ep">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/h_479gah00124/h_479gah00124ps.jpg" title="生中出し解禁 8時間 2">
                        </div>
                        <div class="photo-info">
	                       <span>生中出し解禁 8時間 2 <br><date>GAH-124</date> / <date>2019-02-22</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6zr0">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/h_067nash00006/h_067nash00006ps.jpg" title="ド迫力BODY爆乳熟女22人！全裸エクササイズ3">
                        </div>
                        <div class="photo-info">
	                       <span>ド迫力BODY爆乳熟女22人！全裸エクササイズ3  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>NASH-006</date> / <date>2019-02-08</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6zfg">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/td017dvaj00310/td017dvaj00310ps.jpg" title="【特価】若菜奈央ちゃんの作品がお手頃価格で登場！バックで突かれ、言葉責めを浴びせられる。尻をビンタされ突かれまくられイク。騎乗位で首を絞められ下から突かれ涙を浮かべながらイク。正常位では激ピスで連続イキ。最後もイケとの命令で、一緒にイカせてもらう！">
                        </div>
                        <div class="photo-info">
	                       <span>【特価】若菜奈央ちゃんの作品がお手頃価格で登場！バックで突かれ、言葉責めを浴びせられる。尻をビンタされ突かれまくられイク。騎乗位で首を絞められ下から突かれ涙を浮かべながらイク。正常位では激ピスで連続イキ。最後もイケとの命令で、一緒にイカせてもらう！  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>DVAJ-310</date> / <date>2019-02-02</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6xtm">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/41hodv021357/41hodv021357ps.jpg" title="綺麗なお姉さんは好きですか？清楚な顔して淫らに喘ぎ乱れる濃密セックス！！">
                        </div>
                        <div class="photo-info">
	                       <span>綺麗なお姉さんは好きですか？清楚な顔して淫らに喘ぎ乱れる濃密セックス！！ <br><date>HODV-21357</date> / <date>2019-02-01</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6x9h">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/cjob00040/cjob00040ps.jpg" title="肉厚デカ尻にプレスされて射精させられたボク。4時間ベスト">
                        </div>
                        <div class="photo-info">
	                       <span>肉厚デカ尻にプレスされて射精させられたボク。4時間ベスト  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>CJOB-040</date> / <date>2019-01-20</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6wog">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/dvaj00372/dvaj00372ps.jpg" title="中出し寸前の一番気持ちいい生ピストン82連発">
                        </div>
                        <div class="photo-info">
	                       <span>中出し寸前の一番気持ちいい生ピストン82連発 <br><date>DVAJ-372</date> / <date>2019-01-13</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6wm0">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/mizd00121/mizd00121ps.jpg" title="暗転なし！ブツ切りなし！ MOODYZ最新1年分中出し全部入り超絶技巧編集ベスト">
                        </div>
                        <div class="photo-info">
	                       <span>暗転なし！ブツ切りなし！ MOODYZ最新1年分中出し全部入り超絶技巧編集ベスト  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>MIZD-121</date> / <date>2019-01-12</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6woh">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/mbyd00284/mbyd00284ps.jpg" title="欲求不満な団地妻と孕ませオヤジの汗だく濃厚中出し不倫8タイトル大ボリューム8時間BEST vol.2">
                        </div>
                        <div class="photo-info">
	                       <span>欲求不満な団地妻と孕ませオヤジの汗だく濃厚中出し不倫8タイトル大ボリューム8時間BEST vol.2 <br><date>MBYD-284</date> / <date>2019-01-12</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6w3k">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/41hodv021347/41hodv021347ps.jpg" title="メガネ女子 清楚な瞳の奥にある淫らな本性">
                        </div>
                        <div class="photo-info">
	                       <span>メガネ女子 清楚な瞳の奥にある淫らな本性  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>HODV-21347</date> / <date>2019-01-11</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6x5j">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/84kmvr00519/84kmvr00519ps.jpg" title="【VR】出すならどっぷりナカに！中出しSUPER BEST">
                        </div>
                        <div class="photo-info">
	                       <span>【VR】出すならどっぷりナカに！中出しSUPER BEST <br><date>KMVR-519</date> / <date>2018-12-28</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6xcz">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/2wsp00154/2wsp00154ps.jpg" title="イイオンナ、お貸しします。3 4時間">
                        </div>
                        <div class="photo-info">
	                       <span>イイオンナ、お貸しします。3 4時間  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>WSP-154</date> / <date>2018-12-28</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6xg2">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/td016dvaj00231/td016dvaj00231ps.jpg" title="【特価】女子大生「若菜奈央」ちゃんが、ごみ屋敷のような部屋で無言の男に犯される！抵抗するもバックで入れられるが、なんと奈央のマ〇コはびしょ濡れ。男の荒い息遣いと奈央の喘ぎ声、そしていやらしい音だけが部屋に響く！…">
                        </div>
                        <div class="photo-info">
	                       <span>【特価】女子大生「若菜奈央」ちゃんが、ごみ屋敷のような部屋で無言の男に犯される！抵抗するもバックで入れられるが、なんと奈央のマ〇コはびしょ濡れ。男の荒い息遣いと奈央の喘ぎ声、そしていやらしい音だけが部屋に響く！…  <i class="glyphicon glyphicon-fire" style="color:#cc0000"></i><br><date>DVAJ-231</date> / <date>2018-12-28</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6w1w">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/h_1245avvb00014/h_1245avvb00014ps.jpg" title="【VR】高精細映像4K完全収録！ドッピュン！ドクッドク孕ませ中出しSEX 長尺VR120分">
                        </div>
                        <div class="photo-info">
	                       <span>【VR】高精細映像4K完全収録！ドッピュン！ドクッドク孕ませ中出しSEX 長尺VR120分 <br><date>AVVB-014</date> / <date>2018-12-21</date></span>
	                    </div>
	                </a>
				</div>
							<div class="item">
	                <a class="movie-box" href="https://avmoo.asia/cn/movie/6vlp">
                        <div class="photo-frame">
    	                    <img src="https://jp.netcdn.space/digital/video/hndb00128/hndb00128ps.jpg" title="禁欲女×絶倫男 ナマで覚醒！本能剥き出し真正中出し性交！！シリーズ全6タイトルBEST">
                        </div>
                        <div class="photo-info">
	                       <span>禁欲女×絶倫男 ナマで覚醒！本能剥き出し真正中出し性交！！シリーズ全6タイトルBEST <br><date>HNDB-128</date> / <date>2018-12-16</date></span>
	                    </div>
	                </a>
				</div>
									</div>
		</div>
    </div>
             <div class="hidden-xs mtb-20 text-center">
             <div type="text/data-position" style="display:none">W3siaWQiOiJleG9jX2pfTF83Mjh4OTAiLCJhZHNwb3QiOiJqX0xfNzI4eDkwIiwid2VpZ2h0IjoiMyIsImZjYXAiOmZhbHNlLCJzY2hlZHVsZSI6ZmFsc2UsIm1heFdpZHRoIjpmYWxzZSwibWluV2lkdGgiOiI3NjgiLCJ0aW1lem9uZSI6ZmFsc2UsImV4Y2x1ZGUiOmZhbHNlLCJkb21haW4iOmZhbHNlLCJjb2RlIjoiPHNjcmlwdCB0eXBlPVwidGV4dFwvamF2YXNjcmlwdFwiPlxyXG52YXIgYWRfaWR6b25lID0gXCI4MTMzMDhcIixcclxuXHQgYWRfd2lkdGggPSBcIjcyOFwiLFxyXG5cdCBhZF9oZWlnaHQgPSBcIjkwXCI7XHJcbjxcL3NjcmlwdD5cclxuPHNjcmlwdCB0eXBlPVwidGV4dFwvamF2YXNjcmlwdFwiIHNyYz1cImh0dHBzOlwvXC9hZHMuZXhvY2xpY2suY29tXC9hZHMuanNcIj48XC9zY3JpcHQ+XHJcbjxub3NjcmlwdD48YSBocmVmPVwiaHR0cDpcL1wvbWFpbi5leG9jbGljay5jb21cL2ltZy1jbGljay5waHA/aWR6b25lPTgxMzMwOFwiIHRhcmdldD1cIl9ibGFua1wiPjxpbWcgc3JjPVwiaHR0cHM6XC9cL3N5bmRpY2F0aW9uLmV4b2NsaWNrLmNvbVwvYWRzLWlmcmFtZS1kaXNwbGF5LnBocD9pZHpvbmU9ODEzMzA4Jm91dHB1dD1pbWcmdHlwZT03Mjh4OTBcIiB3aWR0aD1cIjcyOFwiIGhlaWdodD1cIjkwXCI+PFwvYT48XC9ub3NjcmlwdD4ifSx7ImlkIjoianVpY19qX0xfNzI4eDkwIiwiYWRzcG90Ijoial9MXzcyOHg5MCIsIndlaWdodCI6IjMiLCJmY2FwIjpmYWxzZSwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6ZmFsc2UsIm1pbldpZHRoIjoiNzY4IiwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxpZnJhbWUgYm9yZGVyPTAgZnJhbWVib3JkZXI9MCBtYXJnaW5oZWlnaHQ9MCBtYXJnaW53aWR0aD0wIHdpZHRoPTczNiBoZWlnaHQ9OTggc2Nyb2xsaW5nPW5vIGFsbG93dHJhbnNwYXJlbmN5PXRydWUgc3JjPVwvXC9hZHNlcnZlci5qdWljeWFkcy5jb21cL2Fkc2hvdy5waHA/YWR6b25lPTM3MTcwOD48XC9pZnJhbWU+In1d</div>             </div>
             
             <div class="visible-xs-block text-center">
             <div type="text/data-position" style="display:none">W3siaWQiOiJleG9jX2pfTV8zMDB4MjUwIiwiYWRzcG90Ijoial9NXzMwMHgyNTAiLCJ3ZWlnaHQiOiIyIiwiZmNhcCI6ZmFsc2UsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOiI3NjgiLCJtaW5XaWR0aCI6ZmFsc2UsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6ZmFsc2UsImRvbWFpbiI6ZmFsc2UsImNvZGUiOiI8c2NyaXB0PlxyXG5hZF9pZHpvbmUgPSBcIjEwMzEwNDJcIjtcclxuYWRfd2lkdGggPSBcIjMwMFwiO1xyXG5hZF9oZWlnaHQgPSBcIjI1MFwiO1xyXG5pZih0b3A9PT1zZWxmKSB2YXIgcD1kb2N1bWVudC5VUkw7IGVsc2UgdmFyIHA9ZG9jdW1lbnQucmVmZXJyZXI7dmFyIGR0PW5ldyBEYXRlKCkuZ2V0VGltZSgpO1xyXG52YXIgZXhvRG9jdW1lbnRQcm90b2NvbCA9IChkb2N1bWVudC5sb2NhdGlvbi5wcm90b2NvbCAhPSBcImh0dHBzOlwiICYmIGRvY3VtZW50LmxvY2F0aW9uLnByb3RvY29sICE9IFwiaHR0cDpcIikgPyBcImh0dHBzOlwiIDogZG9jdW1lbnQubG9jYXRpb24ucHJvdG9jb2w7XHJcbmlmKHR5cGVvZihhZF9zdWIpID09ICd1bmRlZmluZWQnKSB2YXIgYWRfc3ViID0gXCJcIjtcclxuaWYodHlwZW9mKGFkX3RhZ3MpID09ICd1bmRlZmluZWQnKSB2YXIgYWRfdGFncyA9IFwiXCI7XHJcbnZhciBhZF90eXBlID0gYWRfd2lkdGggKyAneCcgKyBhZF9oZWlnaHQ7XHJcbmlmKGFkX3dpZHRoID09ICcxMDAlJyAmJiBhZF9oZWlnaHQgPT0gJzEwMCUnKSBhZF90eXBlID0gJ2F1dG8nO1xyXG52YXIgYWRfc2NyZWVuX3Jlc29sdXRpb24gPSBzY3JlZW4ud2lkdGggKyAneCcgKyBzY3JlZW4uaGVpZ2h0O1xyXG5kb2N1bWVudC53cml0ZSgnPGlmcmFtZSBmcmFtZWJvcmRlcj1cIjBcIiBzY3JvbGxpbmc9XCJub1wiIHdpZHRoPVwiJyArIGFkX3dpZHRoICsgJ1wiIGhlaWdodD1cIicgKyBhZF9oZWlnaHQgKyAnXCIgc3JjPVwiJyArIGV4b0RvY3VtZW50UHJvdG9jb2wgKyAnXC9cL3N5bmRpY2F0aW9uLmV4b2NsaWNrLmNvbVwvYWRzLWlmcmFtZS1kaXNwbGF5LnBocD9pZHpvbmU9JyArIGFkX2lkem9uZSArICcmdHlwZT0nICsgYWRfdHlwZSArICcmcD0nICsgZXNjYXBlKHApICsgJyZkdD0nICsgZHQgKyAnJnN1Yj0nICsgYWRfc3ViICsgJyZ0YWdzPScgKyBhZF90YWdzICsgJyZzY3JlZW5fcmVzb2x1dGlvbj0nICsgYWRfc2NyZWVuX3Jlc29sdXRpb24gKyAnXCI+PFwvaWZyYW1lPicpO1xyXG48XC9zY3JpcHQ+In0seyJpZCI6Imp1aWNfal9NXzMwMHgyNTAiLCJhZHNwb3QiOiJqX01fMzAweDI1MCIsIndlaWdodCI6IjMiLCJmY2FwIjpmYWxzZSwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6Ijc2OCIsIm1pbldpZHRoIjpmYWxzZSwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxpZnJhbWUgYm9yZGVyPTAgZnJhbWVib3JkZXI9MCBtYXJnaW5oZWlnaHQ9MCBtYXJnaW53aWR0aD0wIHdpZHRoPTMwOCBoZWlnaHQ9MjU4IHNjcm9sbGluZz1ubyBhbGxvd3RyYW5zcGFyZW5jeT10cnVlIHNyYz1cL1wvYWRzZXJ2ZXIuanVpY3lhZHMuY29tXC9hZHNob3cucGhwP2Fkem9uZT0zNzE3MjY+PFwvaWZyYW1lPiJ9XQ==</div>             </div>
             
            <div class="text-center hidden-xs mtb-20">
                            <ul class="pagination pagination-lg mtb-0">
                    <li class="active"><a name="numbar"  href="/cn/star/ozw/page/1">1</a></li><li><a name="numbar"  href="/cn/star/ozw/page/2">2</a></li><li><a name="numbar"  href="/cn/star/ozw/page/3">3</a></li><li><a name="numbar"  href="/cn/star/ozw/page/4">4</a></li><li><a name="numbar"  href="/cn/star/ozw/page/5">5</a></li><li><a name="numbar"  href="/cn/star/ozw/page/6">6</a></li><li><a name="numbar"  href="/cn/star/ozw/page/7">7</a></li><li><a name="nextpage"  href="/cn/star/ozw/page/2">下一页 <span class="glyphicon glyphicon-chevron-right"></span></a></li>                </ul>
                        </div>
<div type="text/data-position" style="display:none">W3siaWQiOiJhZHN0X2pfUE9QVU5ERVIiLCJhZHNwb3QiOiJqX1BPUFVOREVSIiwid2VpZ2h0IjoiNSIsImZjYXAiOiIyIiwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6ZmFsc2UsIm1pbldpZHRoIjoiNzY4IiwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxzY3JpcHQgdHlwZT0ndGV4dFwvamF2YXNjcmlwdCcgc3JjPSdodHRwczpcL1wvcGwxMTA1MjUucHVodG1sLmNvbVwvNzBcLzgyXC85Y1wvNzA4MjljMzgyMTZlMWMwNDYxNmFkYjQ2NzJlZGEzNDIuanMnPjxcL3NjcmlwdD4ifSx7ImlkIjoiY2xpY19qX1BPUFVOREVSIiwiYWRzcG90Ijoial9QT1BVTkRFUiIsIndlaWdodCI6IjciLCJmY2FwIjoiMiIsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOmZhbHNlLCJtaW5XaWR0aCI6ZmFsc2UsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6ZmFsc2UsImRvbWFpbiI6ZmFsc2UsImNvZGUiOiI8c2NyaXB0IHR5cGU9XCJ0ZXh0XC9qYXZhc2NyaXB0XCIgc3JjPVwiXC9cL3h4bGFyZ2Vwb3AuY29tXC9hcHUucGhwP3pvbmVpZD0zNjkwODlcIj48XC9zY3JpcHQ+In0seyJpZCI6ImV4b2Nfal9QT1BVTkRFUiIsImFkc3BvdCI6ImpfUE9QVU5ERVIiLCJ3ZWlnaHQiOiI2IiwiZmNhcCI6IjIiLCJzY2hlZHVsZSI6ZmFsc2UsIm1heFdpZHRoIjpmYWxzZSwibWluV2lkdGgiOiI3NjgiLCJ0aW1lem9uZSI6ZmFsc2UsImV4Y2x1ZGUiOmZhbHNlLCJkb21haW4iOmZhbHNlLCJjb2RlIjoiPHNjcmlwdCBzcmM9XCJcL1wvc3luZGljYXRpb24uZXhvY2xpY2suY29tXC9zcGxhc2gucGhwP2lkem9uZT0xMDA4MDk0JnR5cGU9M1wiPjxcL3NjcmlwdD4ifSx7ImlkIjoiZXhvbV9qX1BPUFVOREVSIiwiYWRzcG90Ijoial9QT1BVTkRFUiIsIndlaWdodCI6IjYiLCJmY2FwIjoiMiIsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOiI3NjgiLCJtaW5XaWR0aCI6ZmFsc2UsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6ZmFsc2UsImRvbWFpbiI6ZmFsc2UsImNvZGUiOiI8c2NyaXB0IHNyYz1cIlwvXC9zeW5kaWNhdGlvbi5leG9jbGljay5jb21cL3NwbGFzaC5waHA/aWR6b25lPTEwMjYzMTImdHlwZT0xMVwiPjxcL3NjcmlwdD5cclxuPHNjcmlwdD5cclxuJChkb2N1bWVudCkucmVhZHkoZnVuY3Rpb24oKSB7XHJcbiAgICBpZiAodHlwZW9mIGV4b1VybCAhPSBcInVuZGVmaW5lZFwiKSB7XHJcbiAgICAgICAgZXhvVXJsID0gZXhvVXJsLnJlcGxhY2UoJ2h0dHA6XC9cLycsICdodHRwczpcL1wvJyk7XHJcbiAgICB9XHJcbiAgICAkKFwiYVwiKS5iaW5kKFwiY2xpY2tcIiwgZnVuY3Rpb24oZXZlbnQpIHtcclxuICAgICAgICBleG9Nb2JpbGVQb3AoKTtcclxuICAgICAgICAkKHRoaXMpLnVuYmluZChcImNsaWNrXCIpO1xyXG4gICAgfSk7XHJcbn0pO1xyXG48XC9zY3JpcHQ+In0seyJpZCI6Imp1aWNfal9QT1BVTkRFUiIsImFkc3BvdCI6ImpfUE9QVU5ERVIiLCJ3ZWlnaHQiOiIyIiwiZmNhcCI6IjEiLCJzY2hlZHVsZSI6ZmFsc2UsIm1heFdpZHRoIjpmYWxzZSwibWluV2lkdGgiOiI3NjgiLCJ0aW1lem9uZSI6ZmFsc2UsImV4Y2x1ZGUiOlsiemgtY24iXSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjwhLS0gSnVpY3lBZHMgUG9wVW5kZXJzIHYzIFN0YXJ0IC0tPlxyXG48c2NyaXB0IHR5cGU9XCJ0ZXh0XC9qYXZhc2NyaXB0XCIgc3JjPVwiaHR0cHM6XC9cL2pzLmp1aWN5YWRzLmNvbVwvanAucGhwP2M9NzQ1NDMzejJ0MjU2cjJxMncyODQ2Mzg0JnU9aHR0cHMlM0ElMkYlMkZidHNvLnB3XCI+PFwvc2NyaXB0PlxyXG48IS0tIEp1aWN5QWRzIFBvcFVuZGVycyB2MyBFbmQgLS0+In1d</div>
<!-- Modal -->
<div class="modal fade" id="advertisingModal" tabindex="-1" role="dialog" aria-labelledby="advertisingModalLabel" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="advertisingModalLabel">Advertising</h4>
      </div>
      <div class="modal-body">
	    <p>Please contact following agents for advertising on AVMOO</p>
	    <p><a href="https://www.exoclick.com/?login=james666" target="_blank" style="color: #D80456;">ExoClick</a> / <a href="https://manage.juicyads.com/juicysites.php?id=128293" target="_blank" style="color: #D80456;">JuicyAds</a> / <a href="http://www.clickadu.com/?rfd=0l1" target="_blank" style="color: #D80456;">ClickADu</a></p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
      </div>
    </div>
  </div>
</div>

<footer class="footer hidden-xs">
	<div class="container-fluid">
<p><a href="https://avmoo.asia/cn/terms">Terms</a> / <a href="https://avmoo.asia/cn/privacy">Privacy</a> / <a href="https://avmoo.asia/cn/usc">2257</a> / <a href="http://www.rtalabel.org/" target="_blank" rel="external nofollow">RTA</a> / <a href="#advertisingModal" role="button" data-toggle="modal">Advertising</a> / <a class="contactus" href="javascript:;" role="button" data-toggle="modal">Contact</a> / <a href="https://tellme.pw/avmoo" target="_blank">Guide</a> | Links: <a href="https://avmoo.asia" target="_blank">AVMOO</a> / <a href="https://avsox.asia" target="_blank">AVSOX</a> / <a href="https://avmemo.asia" target="_blank">AVMEMO</a><br>Copyright © 2013 AVMOO. All Rights Reserved. All other trademarks and copyrights are the property of their respective holders. The reviews and comments expressed at or through this website are the opinions of the individual author and do not reflect the opinions or views of AVMOO. AVMOO is not responsible for the accuracy of any of the information supplied here.</p>
	</div>
</footer>

<div class="visible-xs-block footer-bar-placeholder"></div>

<!-- ////////////////////////////////////////////////// -->
<div class="overlay overlay-contentscale">
    <nav>
        <ul>
            <li>
            <form class="fullsearch-form" action="https://avmoo.asia/cn/search" onsubmit="return false">
               <div class="input-group col-xs-offset-2 col-xs-8">
                  <input name="keyword" type="text" class="form-control" placeholder="搜寻 识别码, 影片, 演员">
                  <span class="input-group-btn">
                    <button class="btn btn-default" type="submit">搜寻</button>
                  </span>
               </div>
            </form>
            </li>
            <li><a href="https://avmoo.asia/cn/released">已发布</a></li>
            <li><a href="https://avmoo.asia/cn/popular">热门</a></li>
            <li><a href="https://avmoo.asia/cn/actresses">女优</a></li>
            <li><a href="https://avmoo.asia/cn/genre">类别</a></li>
        </ul>
    </nav>
    <div class="row overlay-close"><i class="glyphicon glyphicon-remove" style="color:#fff;font-size: 24px;margin:30px;"></i></div>
</div>
<script src='https://avmoo.asia/app/jav/View/js/app.min.js?v=1476953808'></script>  </body>
</html>
`


// 1. 通过演员获取演员信息，以及演员电影url
// 2. 循环电影URL， 通过获取电影信息
// 3. 获取电影图片信息
// 4. 获取磁力链接信息

type Movie struct{
	Title string `json:"title"`
	Logo string	`json:"logo"`
	Info string `json:"info"`
	ReleaseTime string `json:"release_time"`
	Length string `json:"length"`
	Director string `json:"director"`
	Producers string `json:"producers"`
	Issuer string `json:"issuer"`
	Series string `json:"series"`
	Genre string `json:"genre"`
	Performer string `json:"performer"`
}

type Performer struct{
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Age string `json:"age"`
	Height string `json:"height"`
	Cup string `json:"cup"`
	Bust string `json:"bust"`
	Waist string `json:"waist"`
	Hips string `json:"hips"`
	Hometown string `json:"hometown"`
	Hobby string `json:"hobby"`
	Logo string `json:"logo"`
}

type Link struct{
	Title string `json:"title"`
	Url string `json:"url"`
	Name string `json:"name"`
	Size string `json:"size"`
	Date string `json:"date"`
}

type Image struct{
	Title string `json:"title"`
	Url string `json:"url"`
	Name string `json:"name"`
}

type GetMovie struct{
	Title string `json:"title"`
}

func getHtml(url string)(string, error){
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println("获取网页出错, ", err)
		return "", err
	}
	html_text := ""
	buf := make([]byte, 4096)
	for{
		n, err := resp.Body.Read(buf)
		if err != nil && err.Error() != "EOF"{
			fmt.Println("读取异常, " ,err)
			return "", err
		}
		if n == 0{
			break
		}
		html_text += string(buf[:n])
	}
	return html_text, err
}
func (this *GetMovie) getPerformerUrls(htmltext string)([]string, error){
	var urls []string
	html_text := strings.ReplaceAll(strings.ReplaceAll(htmltext,"\n", ""),"  ","")
	urls_ret := regexp.MustCompile(`href="(https://avmoo.asia/cn/star/.*?)"`)
	urls_list := urls_ret.FindAllStringSubmatch(html_text, -1)
	for _, url_item := range urls_list{
		urls = append(urls, url_item[1])
	}
	return urls, nil
}

func (this *GetMovie) getPerformerInfo(htmltext string)(Performer,[]string, error){
	performer := Performer{}
	var movie_urls []string
	html_text := strings.ReplaceAll(strings.ReplaceAll(htmltext,"\n", ""),"  ","")
	info_ret := regexp.MustCompile(`<span class="pb-10">.*?更多无码影片</a>`)
	info_texts := info_ret.FindStringSubmatch(html_text)
	info_text := info_texts[0]
	fmt.Println(info_text)


	name_ret := regexp.MustCompile(`<span class="pb-10">(.*?)</span>`)
	name_text := name_ret.FindStringSubmatch(info_text)
	performer.Name = name_text[1]

	birthday_ret := regexp.MustCompile(`<p>生日: (.*?)</p>`)
	birthday_texts := birthday_ret.FindStringSubmatch(info_text)
	if len(birthday_texts) > 1{
		performer.Birthday = birthday_texts[1]
	}

	age_ret := regexp.MustCompile(`<p>年龄: (.*?)</p>`)
	age_texts := age_ret.FindStringSubmatch(info_text)
	if len(age_texts) > 1{
		performer.Age = age_texts[1]
	}

	height_ret := regexp.MustCompile(`<p>身高: (.*?)</p>`)
	height_texts := height_ret.FindStringSubmatch(info_text)
	if len(height_texts) > 1{
		performer.Height = height_texts[1]
	}

	cup_ret := regexp.MustCompile(`<p>罩杯: (.*?)</p>`)
	cup_texts := cup_ret.FindStringSubmatch(info_text)
	if len(cup_texts) > 1{
		performer.Cup = cup_texts[1]
	}

	bust_ret := regexp.MustCompile(`<p>胸围: (.*?)</p>`)
	bust_texts := bust_ret.FindStringSubmatch(info_text)
	if len(bust_texts) > 1{
		performer.Bust = bust_texts[1]
	}

	waist_ret := regexp.MustCompile(`<p>腰围: (.*?)</p>`)
	waist_texts := waist_ret.FindStringSubmatch(info_text)
	if len(waist_texts) > 1{
		performer.Waist = waist_texts[1]
	}

	hips_ret := regexp.MustCompile(`<p>臀围: (.*?)</p>`)
	hips_texts := hips_ret.FindStringSubmatch(info_text)
	if len(hips_texts) > 1{
		performer.Bust = hips_texts[1]
	}

	urls_ret := regexp.MustCompile(`<a class="movie-box" href="(https://avmoo.asia/cn/movie/.*?)">`)
	urls_list := urls_ret.FindAllStringSubmatch(html_text, -1)
	for _, url_item := range urls_list{
		movie_urls = append(movie_urls, url_item[1])
	}
	return performer,movie_urls, nil
}

func (this *GetMovie) getMovieInfo(htmltext string)(Movie, error){
	// 将回车替换掉
	movie := Movie{}
	html_text := strings.ReplaceAll(strings.ReplaceAll(htmltext,"\n", ""),"  ","")
	// 获取container中少量元素
	container_ret := regexp.MustCompile(`<h3>.*下载`)
	container_texts := container_ret.FindStringSubmatch(html_text)
	container_text := container_texts[0]

	info_text_ret := regexp.MustCompile(`<h3>.*?<h4>`)
	info_texts := info_text_ret.FindStringSubmatch(container_text)
	info_text := info_texts[0]

	// 获取info
	ret := regexp.MustCompile(`<h3>(.*?)</h3`)
	infos := ret.FindStringSubmatch(info_text)

	//fmt.Println(infos[1])

	// 获取title
	title := strings.Fields(infos[1])[0]
	movie.Title = title
	this.Title = title

	// 获取logo封面链接
	logo_ret := regexp.MustCompile(`<a class="bigImage" href="(.*?)"`)
	logos := logo_ret.FindStringSubmatch(info_text)
	movie.Logo = logos[1]

	// 获取发行时间
	release_time_ret := regexp.MustCompile(`发行时间:</span> (.*?)</p>`)
	release_times := release_time_ret.FindStringSubmatch(info_text)
	if len(release_times) > 1{
		movie.ReleaseTime = release_times[1]
	}

	// 获取长度
	length_ret := regexp.MustCompile(`长度:</span> (.*?)</p>`)
	lengths := length_ret.FindStringSubmatch(info_text)
	if len(lengths) > 1{
		movie.Length = lengths[1]
	}
	//fmt.Println(length)

	// 获取导演
	director_ret := regexp.MustCompile(`导演:</span> <a href="https://avmoo.asia/cn/director/.*?">(.*?)</a></p>`)
	directors := director_ret.FindStringSubmatch(info_text)
	if len(directors) > 1{
		movie.Director = directors[1]
	}

	// 获取制作商
	producers_ret := regexp.MustCompile(`制作商: </p><p><a href="https://avmoo.asia/cn/studio/.*?">(.*?)</a></p>`)
	producers := producers_ret.FindStringSubmatch(info_text)
	if len(producers) > 1{
		movie.Producers = producers[1]
	}

	// 获取发行商
	issuer_ret := regexp.MustCompile(`发行商: </p><p><a href="https://avmoo.asia/cn/label/.*?">(.*?)</a></p>`)
	issuer := issuer_ret.FindStringSubmatch(info_text)
	if len(issuer) > 1{
		movie.Issuer = issuer[1]
	}

	// 获取系列
	series_ret := regexp.MustCompile(`系列: </p><p><a href="https://avmoo.asia/cn/series/.*?">(.*?)</a></p>`)
	seriess := series_ret.FindStringSubmatch(info_text)
	if len(seriess) > 1{
		movie.Series = seriess[1]
	}

	// 获取类别
	genre_ret := regexp.MustCompile(`<a href="https://avmoo.asia/cn/genre/.*?">(.*?)</a>`)
	genres := genre_ret.FindAllStringSubmatch(info_text, -1)
	genre := ""
	for _, genre_item := range genres{
		genre += "," + genre_item[1]
	}
	movie.Genre = genre[1:]

	// 获取演员HTML
	performer_html_ret := regexp.MustCompile(`<h4>演员.*?推荐</h4>`)
	performer_htmls := performer_html_ret.FindStringSubmatch(container_text)
	performer_html := ""
	if len(performer_htmls) != 0{
		performer_html = performer_htmls[0]
	}

	// 获取演员
	performer_ret := regexp.MustCompile(`<span>(.*?)</span>`)
	performer_list := performer_ret.FindAllStringSubmatch(performer_html, -1)
	performer := ""
	for _, performer_item := range performer_list{
		performer += "," + performer_item[1]
	}
	movie.Performer = performer[1:]
	return movie, nil
}

func (this *GetMovie) getImage(htmltext string)([]Image, error){
	images := []Image{}
	//image := Image{}
	html_text := strings.ReplaceAll(strings.ReplaceAll(htmltext,"\n", ""),"  ","")
	images_ret := regexp.MustCompile(`<h4>样品图像</h4>.*?<h4>推荐</h4>`)
	images_text := images_ret.FindStringSubmatch(html_text)[0]
	//fmt.Println(images_text)

	image_ret := regexp.MustCompile(`href="(.*?)" title="(.*?)"`)
	image_list := image_ret.FindAllStringSubmatch(images_text, -1)
	//fmt.Println(image_list)
	for _, image_item := range image_list{
		//fmt.Println(image_item, len(image_item))
		image := Image{}
		image.Title = this.Title
		image.Url = image_item[1]
		image.Name = image_item[2]
		images = append(images, image)
	}
	return images, nil
}

type Mysql struct{
	Db sql.DB
}

func (this *Mysql)Init()error{
	db, err := sql.Open("mysql", "root:xiaoxin@tcp(127.0.0.1:3306)/video?charset=utf8")
	if err != nil{
		return err
	}
	err = db.Ping()
	if err != nil{
		return err
	}
	this.Db = *db
	return nil
}
func(this *Mysql) SaveMovie(){
	stmt, err := this.Db.Prepare("insert into")
}
func (this *Mysql) SavePerformer(){

}
func (this *Mysql) SaveLink(){

}
func (this *Mysql) SaveImage(){

}

func main(){
	get_movie := GetMovie{}
	// 获取演员页面
	performer_text, err := getHtml("https://avmoo.asia/cn/actresses")
	if err != nil{
		fmt.Println(err)
		return
	}
	performer_urls, err := get_movie.getPerformerUrls(performer_text)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(performer_urls)

	for index, per_url := range performer_urls{
		performer_info_text, err := getHtml(per_url)
		if err != nil{
			fmt.Println(err)
			continue
		}
		performer, movie_urls, err := get_movie.getPerformerInfo(performer_info_text)
		if err != nil{
			fmt.Println(err)
			continue
		}
		fmt.Println(performer, movie_urls)
		if index > 5{
			break
		}
	}
	//
	//html_text, err := getHtml("https://avmoo.asia/cn/movie/7704")
	////fmt.Println(html_text, err)
	//
	//movie, err := get_movie.getMovieInfo(html_text)
	//if err != nil{
	//	return
	//}
	//fmt.Println(movie)
	//images, err := get_movie.getImage(html_text)
	//if err != nil{
	//	return
	//}
	//fmt.Println(images)
}
