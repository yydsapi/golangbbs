/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "header"}}
<!DOCTYPE html>
<html lang="en">
<head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="description" content="Default Project in Go using Goji, mgo, Gorilla sessions">
	<meta name="keywords" content="go, golang, bbs, blog, mysql, golangbbs, mongo, microframework, win, sessions, mvc">
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
  <script src="/static/langs/locale/en.js"></script>
  <script src="/static/langs/locale/zh-CN.js"></script>
  <script src="/static/langs/vue-i18n.js"></script>
  <script src="/static/common.js"></script>  
</head>
<body><input type="hidden" name="_csrf" id="_csrf" value="{{.Csrf}}"><img src="/static/img/{{.RightImgSrc}}" class="alignright" onclick="window.open('https://github.com/kdhly/golangbbs')" >
<div id="appdialog1">        
            <el-dialog :visible.sync="dialogVisible1" :title="$t('common.tip')" ref="dailog" @open="open()">
              <div id="dialogcontent1"></div></el-dialog></div>
<div id="appdialog2"><el-dialog :title="$t('common.tip')" :visible.sync="dialogVisible2" width="30%" center @open="open()">
  <div id="dialogcontent2"></div><span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="appdialog2Confirm">[[[$t('common.confirm')]]]</el-button>
    <el-button @click="dialogVisible2 = false">[[[$t('common.cancel')]]]</el-button>
</el-dialog></div> 
<div id="appmenu" class="mycenter" >
<el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
{{.NavMenuStr}}<img src=/static/img/photos.png style="width:30px;height:30px;position:relative;margin-left:15px;top:4px;border:0px;cursor:pointer;" onclick="location.href='/bbs/photos'" /><img src=/static/img/musics.png style="width:30px;height:30px;position:relative;margin-left:25px;top:4px;border:0px;cursor:pointer;" onclick="location.href='/bbs/musics'" /><img src=/static/img/videos.png style="width:30px;height:30px;position:relative;margin-left:25px;top:4px;border:0px;cursor:pointer;" onclick="location.href='/bbs/videos'" />
</el-menu>
 </div>
 <script type="text/javascript">
  Vue.config.delimiters = ['[[[', ']]]'];
  if({{.LanguageStr}}.indexOf("cn")==-1){
      Vue.locale('en', ELEMENT.lang.en)
  }else{
      Vue.locale('zh-CN', ELEMENT.lang.zhCN)
      Vue.config.lang = "zh-CN"
  }
</script>
{{end}}
