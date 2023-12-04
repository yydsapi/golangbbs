/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "index/photoelement"}}
<!DOCTYPE html>
<html lang="en">
<head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="description" content="Default Project in Go using Goji, mgo, Gorilla sessions">
    <meta name="keywords" content="go, golang, mgo, mongo, goji, microframework, gorilla, sessions, mvc">
    <link href="/static/favicon.ico" rel="icon" type="image/x-icon">
    <title>{{.WebTitle}}</title>
   <script>
  try{!Promise}catch(e){location.href="/static/error.htm"}
       </script>
 <!-- import CSS -->
  <link rel="stylesheet" href="/static/elementUI/css/element.css">
    <link rel="stylesheet" href="/static/common.css">
 <!-- import Vue before Element -->
  <script src="/static/elementUI/js/vue.js"></script>
 <!-- import JavaScript -->
  <script src="/static/elementUI/js/element.js"></script>
  <script src="/static/elementUI/axios.min.js"></script>
          <script src="/static/httpVueLoader.js"></script>
    <script src="/static/common.js"></script>
</head>
<body bgcolor="#000000"><input type="hidden" name="_csrf" id="_csrf" value="{{.Csrf}}">
<div id="appcarousel">  
    <el-carousel :height="bannerH +'px'">
     <el-carousel-item v-for="(item,index) in BannerImg" :key="index" :align="center">
      <img src="/static/tmpimg/1.jpg" v-if="index == 0" class="bannerImg" />
       <img src="/static/tmpimg/2.jpg" v-if="index == 1" class="bannerImg" />
       <img src="/static/tmpimg/3.jpg" v-if="index == 2" class="bannerImg" />
       <img src="/static/tmpimg/4.jpg" v-if="index == 3" class="bannerImg" />
     </el-carousel-item>
   </el-carousel>
</div>
</body>
<script type="text/javascript">
             new Vue({
        el: '#appcarousel',
        data () {
            return {
              BannerImg: [{
                  item: '/static/tmpimg/1.jpg',
                  index: 0
                },{
                  item: '/static/tmpimg/2.jpg',
                  index: 1
                },{
                  item: '/static/tmpimg/3.jpg',
                  index: 2
                },{
                  item: '/static/tmpimg/4.jpg',
                  index: 3
                }],
                bannerH:200,
            }
          },
methods:{
    setBannerH(){
      this.bannerH = window.screen.availHeight-110
    }
  },
  mounted(){
    this.setBannerH()
    window.addEventListener('resize', () => {
      this.setBannerH()
    }, false)
  },
  created(){}

    })
</script>
{{end}}