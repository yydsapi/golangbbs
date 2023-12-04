/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "index/about"}}
{{template "header" .}}
        <h4 class="card-title">{{.Title}}</h4>
               <div id="app"  class="mycenter">
                <span v-html="commonForm.content"></span> 
              </div>
          </div>
          </div>
  <script>
    var ts=this;
  var Main = {
    delimiters: ['[[[', ']]]'],
    data() {
      return {
        commonForm: {
          content: this.$t('about.content')
        }
      }
    },
    methods: {

    }
  }

var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
</script>
{{template "footer" .}}
{{end}}
