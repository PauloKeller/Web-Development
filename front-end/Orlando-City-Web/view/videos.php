<?php include_once("header.php"); ?>

<link rel="stylesheet" href="lib/plyr/dist/plyr.css">


<section>

    <div id="call-to-action">

        <div class="container">

            <div class="row text-center">
                <h2>VIDEOS</h2>
                <hr>
            </div>

            <div class="row">

                <div class="player">
                    <video src="mp4/v.mp4" controls poster="img/highlights.jpg"></video>
                </div>
                <input type="range" id="volume" min="0" max="1" step="0.01" value="1">

                <button type="button" id="btn-play-pause" class="btn btn-success">PLAY</button>

            </div>

        </div>
    </div>

    <div id="news" class="container" style="top:0">

        <div class="row" class="row text-center">
            <h2>Latest News</h2>
            <hr>
        </div>

        <div class="row thumbnails">

            <div class= "item" data-video="vs">
                <div class="item-inner">
                    <img src="img/highlights.jpg" alt="noticia">
                    <h3>Orlando City Acquires </h3>
                </div>
            </div>
            <div class= "item" data-video="orlando_fundation">
                <div class="item-inner">
                    <img src="img/img3.jpg" alt="noticia">
                    <h3>Orlando City</h3>
                </div>
            </div>
            <div class= "item" data-video="orlando_fundation">
                <div class="item-inner">
                    <img src="img/img2.jpg" alt="noticia">
                    <h3>Orlando</h3>
                </div>
            </div>
            <div class= "item" data-video="orlando_fundation">
                <div class="item-inner">
                    <img src="img/highlights.jpg" alt="noticia">
                    <h3>Lorem ipsum dolor sit amet</h3>
                </div>
            </div>
            <div class= "item"  data-video="orlando_fundation">
                <div class="item-inner">
                    <img src="img/highlights.jpg" alt="noticia">
                    <h3>Lorem ipsum dolor sit amet, consec</h3>
                </div>
            </div>
            <div class= "item" data-video="v">
                <div class="item-inner">
                    <img src="img/highlights.jpg" alt="noticia">
                    <h3>Lorem ipsum dolor sit amet, consec</h3>
                </div>
            </div>
            <div class= "item" data-video="v">
                <div class="item-inner">
                    <img src="img/highlights.jpg" alt="noticia">
                    <h3>Lorem ipsum dolor sit amet, consec</h3>
                </div>
            </div>
            <div class= "item" data-video="v">
                <div class="item-inner">
                    <img src="img/highlights.jpg" alt="noticia">
                    <h3>Lorem ipsum dolor sit amet, consec</h3>
                </div>
            </div>
        </div>

    </div>

</section>

<?php include_once("footer.php"); ?>

<script src="lib/plyr/dist/plyr.js"></script>

<script>
$.get("lib/plyr/dist/plyr.svg", function(data) {
      var div = document.createElement("div");
      div.innerHTML = new XMLSerializer().serializeToString(data.documentElement);
      document.body.insertBefore(div, document.body.childNodes[0]);
    });
</script>
<script>
$(function(){
    $(".thumbnails .item").on("click", function(){

        $("video").attr({
            "scr":"mp4/"+$(this).data('video')+".mp4",
            "poster":"img/"+$(this).data('video')+".jpg"
        });

    });

    $("#volume").on("change", function(){

        $("video")[0].volume = parseFloat($(this).val());

    });

    $("#btn-play-pause").on("click", function(){

        var video = $("video")[0];

        if($(this).hasClass("btn-success")) {
            $(this).text("STOP");
            video.play();
        } else {
            $(this).text("PLAY");
            video.pause();
        }

        $(this).toggleClass("btn-success btn-danger");
    });

    plyr.setup();

});
</script>

