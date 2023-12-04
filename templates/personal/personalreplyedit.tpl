/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "bbs/personal/editreply"}}
{{template "header" .}}
<div id="app" class="mycenter">
  <div  class="webtitle"><h5>{{.PageTitle}}</h5></div>
<el-form  :model="commonForm" @submit.native.prevent ref="commonForm" label-width="100px" id="myForm" class="demo-commonForm" size="medium">
                <el-form-item :label="$t('common.content')" prop="ck-content" >
                  <textarea name="content" id="ck-content" style="height:400px">{{.Replycontent}}</textarea>
                </el-form-item>
                
                <el-form-item>
                  <el-button type="primary" native-type="submit" @click="submitForm('commonForm')">[[[$t('common.submit')]]]</el-button>
                  <el-button @click="resetForm('commonForm')">[[[$t('common.reset')]]]</el-button>
                </el-form-item>
</el-form>
              </div>
          </div>
          </div>
 <script src="/static/ckeditor/build/ckeditor.js"></script>
  <script src="/static/ckeditor/build/translations/zh-cn.js"></script>
  <link rel="stylesheet" href="/static/ckeditor/build/custom.css">
  <script>

  var Main = {
      delimiters: ['[[[', ']]]'],
    data() {
      return {        
        commonForm: {
          
        }
      }
    },
    methods: {
      submitForm(formName) {

 if (editor.getData().toString().length<9) {
  this.$message({
                          message: this.$t('publish.contentsnbl'),
                          type: 'warning'
                        });
  return false;
}
            var form = document.getElementById("myForm");
            var formData = new FormData();
            for (const key of Object.keys(this.commonForm)) {
                formData.append(key, this.commonForm[key])
              }
            var ts=this
            formData.append("content",editor.getData());
            formData.append("Csrf", document.getElementById("_csrf").value)
            formData.append("id",{{.Id}});
                  axios.post('/bbs/updatereplypost',formData)
                  .then(function (response) {
                    if(response.data.info=="ok"){
                      mess=ts.$t('common.opsuccess')
                      dg1.dialogVisible1 = true;
                      setTimeout(function(){ location.href=response.data.returnURL; },1000);
                    }
                    else{
                      mess=ts.$t(response.data.info);
                      dg1.dialogVisible1 = true;
                    }
                  })
                  .catch(function (error) {
                    mess=ts.$t('common.serverError')
                    dg1.dialogVisible1 = true;
                    console.log(error);
                  });
      },
      resetForm(formName) {
        location.href=location.href;
      }

    }
  }

var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
    var data;
    ClassicEditor.create(document.getElementById('ck-content'), {
      language: 'en',
          ckfinder: {
            uploadUrl: '/bbs/upload',

          }
        }
    ).then(editor => {
    window.editor = editor;
    data = editor.getData();

    console.log(data);
    } )
    .catch(error => {
        console.log(error);
    });
</script>
{{template "footer" .}}
{{end}}
