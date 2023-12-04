 /*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "index/video"}}
  <!DOCTYPE html>
<!--
/*
 * blueimp Gallery Demo
 * https://github.com/blueimp/Gallery
 *
 * Copyright 2013, Sebastian Tschan
 * https://blueimp.net
 *
 * Licensed under the MIT license:
 * https://opensource.org/licenses/MIT
 */
-->
<html lang="en">
  <head>
    <!--[if IE]>
      <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
    <![endif]-->
    <meta charset="utf-8" />
    <title>blueimp Gallery</title>
    <meta
      name="description"
      content="blueimp Gallery is a touch-enabled, responsive and customizable image and video gallery, carousel and lightbox, optimized for both mobile and desktop web browsers. It features swipe, mouse and keyboard navigation, transition effects, slideshow functionality, fullscreen support and on-demand content loading and can be extended to display additional content types."
    />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <script src="/static/js/jquery.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/html5video/main.css" />
    <link rel="stylesheet" href="/static/common.css">
<script src="/static/html5video/html5media.min.js"></script>
  </head>
  <body>
<div class="mycenter">
  <h2 class="top_title"><a href="/" title="Upper level"><img src="/static/img/return.png" class="imgreturn"></a> Random play</h2>
    <!-- The container for the list of example images --><br>
<video id="myvideo" width="400" height="250" controls preload="" >
  Your browser does not support HTML5 playback
  <source src="{{.UrlOne}}"></source>
</video>

</div>
    <script>
String.prototype.replaceAll = function(s1,s2){ 
return this.replace(new RegExp(s1,"gm"),s2); 
}
          var video = document.getElementById("myvideo");
    $(document).ready(function(){
        //setTimeout(function() {video.play();},1000);
    });

        var fs={{.UrlOne}}.split("/")
        document.getElementById("mediatitle").innerHTML=fs[fs.length-1]
    var vLists = "{{.Url}}";
    vLists = vLists.replaceAll(":","\/")
    var vList = vLists.split(",")
    var vLen = vList.length;
    var curr = 0;

    video.addEventListener('ended', function(){
       // alert("continues")
        play();
    });
    function play() {
        video.src = vList[curr];
        var s=vList[curr].split("/")
        //document.getElementById("mediatitle").innerHTML=s[s.length-1]
        video.load();
        video.play();
        curr++;
        if(curr >= vLen){
            curr = 0; //repeate
        }

    }
    </script><br>
    <!--<script src="js/demo/demo.js"></script>-->
    <br><br>
          <div align="center">
        <p class="text-footer">&copy; 2019 golangbbs.com qq:1269866868 mail:golangbbs@gmail.com </p>
    </div>
  </body>
</html>

{{end}}