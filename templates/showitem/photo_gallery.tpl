 /*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "index/photogallery"}}
  <script src="/static/elementUI/js/vue.js"></script>
  <script type="text/javascript" src="/static/vue-gallery/blueimp-helper.js"></script> 
  <script type="text/javascript" src="/static/vue-gallery/blueimp-gallery.js"></script> 
  <script type="text/javascript" src="/static/vue-gallery/blueimp-gallery-fullscreen.js"></script> 
  <script type="text/javascript" src="/static/vue-gallery/vue-gallery.js"></script> 
  <link rel="stylesheet" type="text/css" href="/static/vue-gallery/blueimp-gallery.min.css">
  
 
<div id="app">
  <gallery :images="images" :index="index" @close="index = null"></gallery>
  <div
    class="image"
    v-for="image, imageIndex in images"
    @click="index = imageIndex"
    :style="{ backgroundImage: 'url(' + image + ')', width: '300px', height: '200px' }"
  ></div>
</div>
 
<script type="text/javascript">
  new Vue({
    el: '#app',
    data: function () {
      return {
        images: [
          '/static/tmpimg/1.jpg',
          '/static/tmpimg/2.jpg',
          '/static/tmpimg/3.jpg',
          '/static/tmpimg/4.jpg'
        ],
        index: null
      };
    },
 
    components: {
      'gallery': VueGallery
    }
  });
</script>
{{end}}