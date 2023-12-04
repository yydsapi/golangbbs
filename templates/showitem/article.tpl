/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "bbs/show"}}
{{template "header" .}}
<link rel="stylesheet" type="text/css" href="/static/html5video/main.css" />
<script src="/static/html5video/html5media.min.js"></script>
<script src="/static/ckeditor/build/ckeditor.js"></script>
<script src="/static/ckeditor/build/translations/zh-cn.js"></script>
<link rel="stylesheet" href="/static/ckeditor/build/custom.css">
<div class="mycenter">
<div class="clearfix"><html> <h3><div class="bbsarticle"><a href="#" onclick="history.back(-1)" title="PrvPage"><img src="/static/img/return.png" class="imgreturn"></a> {{.Title}}</h3></html></div> 
<div id="apprate" style="width:20%;float:right;"><el-rate
  v-model="value5"
  disabled
  show-score
  text-color="#ff9900"
  score-template="{{.Raise_count}}">
</el-rate></div><div style="float: right;margin-right:25px;">{{.PrvPage}} {{.NextPage}}</div></div><div style="display:block;width:100%;line-height: 12px"> <img src="/static/img/space.png" border=0px height=12px /></div>
<div class="bbscontent">
<div id="appdivider"><el-divider content-position="right">{{.Author}} <font class="graw">{{.Add_time}}</font></el-divider></div>
	{{.Media}}{{.Attachment}}<div class="bbscontent0">{{.Content}}</div>
</div><br>
<div class="mycenter">
  <div ><img src='{{.Headicon}}' width=26px height=26px /> {{.Author}} <img src='/static/img/level/{{.UserLever}}.png' width=20px height=20px /><font class="graw">{{.Add_time}}</font></div>
    <div class="clearfix" >
    <div id="apprate2" style="width:50%;float:left;height:30px;">
    <el-rate
      v-model="value2"
      @change="currentSel"
      :colors="colors">
    </el-rate>
</div><div style="float: right;margin-right:85px">{{.PrvPage}} {{.NextPage}}</div><div style="width:90%;float:right;"><span align=right>{{.Edit_history}}</span></div>
</div>
{{.ReplayContent}}<br>{{.ShowTxt}}
<div id="app" style="display:{{.ShowReply}}">
  <el-form :model="commonForm" :rules="rules" ref="commonForm" label-width="2px" id="myForm" class="demo-commonForm" size="medium">
        <el-form-item label="" prop="ck-content" >
            <textarea name="content" id="ck-content" style="height:300px"></textarea>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="submitForm('commonForm')">[[[$t('common.reply')]]]</el-button>
            <el-button @click="resetForm('commonForm')">[[[$t('common.reset')]]]</el-button>
        </el-form-item>
  </el-form>
</div>
</div>
</div>
<script>
var Main = {
   delimiters: ['[[[', ']]]'],
    data() {
      return {
        commonForm: {
        },
        rules: {
        }
      }
    },
    methods: {
      submitForm(formName) {
        var ts=this
          if (editor.getData().toString().length>8) {
              var formData = new FormData();
              formData.append("myid", document.getElementById("bbsarticleid").innerText)
              formData.append("Csrf", document.getElementById("_csrf").value)
              formData.append("mycontent", editor.getData())
              axios.post('/bbs/bbsreply',formData)
              .then(function (response) {
               if(response.data.info=="ok"){
                     ts.$message({
                      message: ts.$t('publish.replysuccess'),
                      type: 'success'
                    });
                    editor.setData("")
                }else{
                     ts.$message({
                       message: 'error:'+ts.$t(response.data.info),
                      type: 'warning'
                    });
                }
              })
              .catch(function (error) {
                mess=ts.$t('common.serverError')
                dg1.dialogVisible1 = true;
                console.log(error);
              });
          } else {
              ts.$message({
                message: this.$t('publish.replysnbl'),
                type: 'warning'
              });
          }
      },
      resetForm(formName) {
        editor.setData("");
      }

    }
  }
var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
    var data;
    ClassicEditor.create(document.getElementById('ck-content'), {
      language: {{.CkeditLanguage}},
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
var appdivider = new Vue({
      el: '#appdivider',
      data: function() {
        return {}
      },
      methods: {
            open () {
                //this.$nextTick(() => document.getElementById("dialogcontent1").innerHTML=mess)
            }
        }
    })
var apprate = new Vue({
      el: '#apprate',
      data: function() {
        return {value5: {{.Raise_count}}}
      }
    })
var apprate2 = new Vue({
      el: '#apprate2',
      methods: {
            currentSel(selVal) {
              var ts=this
              var formData = new FormData();
              formData.append("myid", document.getElementById("bbsarticleid").innerText)
              formData.append("Csrf", document.getElementById("_csrf").value)
              formData.append("myraise", selVal.toString())
              axios.post('/bbs/raiseitem',formData)
              .then(function (response) {
               if(response.data.info=="ok"){
                   ts.$message({
                    message: ts.$t('publish.pfsecuss'),
                    type: 'success'
                  });
                }
                else{
                   ts.$message({
                    message: 'error:'+ts.$t(response.data.info),
                    type: 'warning'
                  });
                }
              })
              .catch(function (error) {
                mess=ts.$t('common.serverError')
                dg1.dialogVisible1 = true;
                console.log(error);
              });
            }
        },
     data() {
      return {
        value2: null,
        colors: ['#99A9BF', '#F7BA2A', '#FF9900']  // == { 2: '#99A9BF', 4: { value: '#F7BA2A', excluded: true }, 5: '#FF9900' }
      }
    }
})
</script>
<div id="bbsarticleid" style="display: none;">{{.Id}}</div>
{{template "footer" .}}
{{end}}