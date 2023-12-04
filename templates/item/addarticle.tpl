/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "bbs/addarticle"}}
{{template "header" .}}
<div id="app" class="mycenter">
  <div  class="webtitle"><h5>{{.PageTitle}}</h5></div>
<el-form :model="commonForm" @submit.native.prevent :rules="rules" ref="commonForm" label-width="130px" id="myForm" class="demo-commonForm" size="medium">
    <el-form-item :label="$t('common.title')" prop="mytitle" >
      <el-col :span="12"><el-input v-model="commonForm.mytitle"  :placeholder="$t('publish.entertitle')"  class="mytitle"></el-input></el-col>
    </el-form-item>
<el-row>
   <el-col :span="10">
    <el-form-item :label="$t('common.sort')" prop="category">
      <el-cascader
          v-model="commonForm.category"
          :options="options"
          ref="CateGoryStr" 
          expand-trigger="hover"
          :props="{ expandTrigger: 'hover' }"
          :placeholder="$t('publish.selectsort')"
          @change="handleChange">
          </el-cascader>
          </el-form-item>
    </el-col>
  <el-col :span="6">
      <el-form-item :label="$t('publish.isprivate')" prop="isprivate">                        
        <el-switch
          v-model="commonForm.isprivate"
          active-color="#13ce66"
          ref="isprivate" 
          inactive-color="#DEDEDE">
        </el-switch>
      </el-form-item>
  </el-col>
</el-row>
{{if ge .UserLevel 5}}
    <el-form-item label="Media" prop="ck-mediaurl" >
      <input type="file" name="image" id="file_upload" style="display:none">
      <el-button type="primary" @click="selectFiles('commonForm')">[[[$t('common.select')]]]</el-button>
      <el-button type="primary" @click="submitMedia('commonForm')">[[[$t('common.upload')]]]</el-button>
      <el-input v-model="commonForm.mediaurl" placeholder="" id="mediaurl" style="width:40%;"></el-input>
    </el-form-item>
    <div style="margin-left:95px;margin-top:-15px;color:#ff7575;font-size: 12px;vertical-align:middle;line-height: 30px;">[[[$t('publish.sttsu')]]]</div>
    <el-form-item :label="$t('common.share')" prop="strshare">                        
    <el-input v-model="commonForm.strshare" id="strshare" style="width:40%;margin-left:10px"></el-input>
    </el-form-item>
     {{end}} {{if ge .UserLevel 5}}
    <el-form-item :label="$t('common.attachment')" prop="" >
          <div style="width:50%">
          <el-upload
          class="upload-demo"
          action="/bbs/uploadatt"
          :on-preview="handlePreview"
          :on-remove="handleRemove"
          :before-remove="beforeRemove"
          :on-change="handleChangeUpload"
          :on-success="handleUploadSuccess"
          multiple
          :limit="3"
          :on-exceed="handleExceed"
          :data="CsrfStr" 
          :file-list="fileList">
          <el-button size="small" type="primary">[[[$t('publish.clicktoupload')]]]</el-button>
          <!--<div slot="tip" class="el-upload__tip">Only upload JPG / png files, and no more than 500 KB</div>-->
          </el-upload></div>
    </el-form-item>
    {{end}}
    <el-form-item :label="$t('common.content')" prop="ck-content" >
      <textarea name="content" id="ck-content" style="height:400px"></textarea>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" native-type="submit" @click="submitForm('commonForm')">[[[$t('publish.publishImmediately')]]]</el-button>
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
var flist
  var Main = {
    delimiters: ['[[[', ']]]'],
    data() {
      return {
        //value: ["",""],
        CsrfStr:{
            Csrf:'{{.Csrf}}'
        },
        options:JSON.parse({{.CateGoryStr}}),
        commonForm: {
          mytitle: '',
          category: {{.Categoryenall}}.split(","),
          mediaurl: '',
          isprivate: false,
          strshare:'',
        },
        rules: {
          mytitle: [
            { required: true, message: this.$t('publish.entertitle'), trigger: 'blur' },
          ],
          category: [{ required: true, message: this.$t('publish.selectsort'), trigger: 'change' }],
        },
        fileList: []
      }
    },
    methods: {
      handleUploadSuccess(response,file,fileList) {
        flist=fileList
      },
      handleRemove(file, fileList) {
        flist=fileList
                var atturl=""
        try {
            atturl=file.response.url
        }
        catch(err) {
            atturl=file.url
        }
        finally {   
        }
        var ts=this;
        if (atturl.indexOf("/media")!=-1){
          var formData = new FormData();
          formData.append('att', atturl);
          formData.append("Csrf", document.getElementById("_csrf").value)
          axios.post('/bbs/deleteatt',formData)
          .then(function (response) {
           
          })
          .catch(function (error) {
            mess=ts.$t('common.serverError')
            dg1.dialogVisible1 = true;
            console.log(error);
          });
        }
    },
      handlePreview(file) {
        //console.log(file);
      },
      handleExceed(files, fileList) {
        this.$message.warning(this.$t('publish.selectrestr1')+`${files.length}`+this.$t('publish.selectrestr2')+`${files.length + fileList.length}` + ` ` + this.$t('common.documents'));
      },
      beforeRemove(file, fileList) {
        return this.$confirm(this.$t('tips.comfirmdelete')+` ${ file.name } ?`);
      },
      handleChangeUpload(file, fileList) {
        this.fileList = fileList.slice(-3);
      },
      handleChange(value) {
        if(value[0]=="blog"){
         this.commonForm.isprivate=true
        }else{
          this.commonForm.isprivate=false
        }
      },
      submitForm(formName) {
        var ts=this
        if (editor.getData().toString().length<9) {
          this.$message({
             message: this.$t('publish.contentsnbl'),
             type: 'warning'
          });
          return false;
        }
        this.$refs[formName].validate((valid) => {
          if (valid) {
            var flisttmp=""
          try{
            for(var i=0;i<flist.length;i++){
                if (flisttmp==""){
                  flisttmp=flist[i].response.url
                }else{
                  flisttmp+="?"+flist[i].response.url
                }
            }
          }catch(e){flisttmp=""}
            var form = document.getElementById("myForm");
            var formData = new FormData();
            for (const key of Object.keys(this.commonForm)) {
                formData.append(key, this.commonForm[key])
              }
            var categoryens=this.$refs['CateGoryStr'].value
            if (categoryens=="[]") {
                mess=ts.$t('publish.selectsort')
                dg1.dialogVisible1 = true;
                return;
            }
            
            formData.append("attachment",flisttmp);
            try{formData.append("strshare", document.getElementById("strshare").value)}catch(e){formData.append("strshare", "")}
            formData.append("categoryen",categoryens[categoryens.length-1]);
            formData.append("categoryenall",this.$refs['CateGoryStr'].value);
            formData.append("categorycnall",this.$refs['CateGoryStr'].currentLabels);
            formData.append("content",editor.getData());
            formData.append("Csrf", document.getElementById("_csrf").value)
            try{
              var mediaurl=document.getElementById("mediaurl").value
              formData.append("mediaurltmp",mediaurl);
            }catch(e){}
            axios.post('/bbs/addarticlepost',formData)
            .then(function (response) {
              if(response.data.info=="ok"){
                mess=ts.$t('publish.publishsucc')
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
          } else {
            return false;
          }
        });
      },
      submitMedia(formName) {
        //alert(this.$refs['CateGoryStr'].value)
        //alert(this.$refs['CateGoryStr'].currentLabels)
        var file_upload=document.getElementById("file_upload").value.trimall()
        if (file_upload==""){
          this.$message({
            message: this.$t('publish.plsselectfile')+'(.mp3.ogg.mp4.mkv.avi)',
            type: 'warning'
          });
          return;
        }
        var extname="."+file_upload.toLowerCase().split('.').splice(-1)
        if (extname==".mp4" || extname==".mp3" || extname==".ogg" || extname==".wav" || extname==".flac" || extname==".ape" || extname==".mkv" || extname==".avi" || extname==".rmvb") {
        var fileM=document.getElementById("file_upload")
        var fileObj = fileM.files[0];
        var ts=this;
        var formData = new FormData();
        formData.append('upload', fileObj);
        formData.append("Csrf", document.getElementById("_csrf").value)
            axios.post('/bbs/upload',formData)
            .then(function (response) {
              if(response.data.uploaded){
                document.getElementById("mediaurl").value=response.data.url
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
          } else {
            this.$message({
                message: this.$t('publish.wftm')+'('+this.$t('publish.mustbe')+'.mp3.ogg.mp4.mkv.avi.wav.falc.ape.rmvb)',
                type: 'warning'
            });
          }
      },
      selectFiles(formName) {
              document.getElementById('file_upload').click()
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }

    }
  }

var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
var data;
      //CKEDITOR.plugins.addExternal( 'html5video', '/static/ckeditor/build/plugins/html5video', 'plugin.js' );
      //ClassicEditor.builtinPlugins = [ html5video ];
    ClassicEditor.create(document.getElementById('ck-content'), {
      language: {{.CkeditLanguage}},
      //toolbar: ["heading", "|", "alignment:left", "alignment:center", "alignment:right", "alignment:adjust", "|", "bold", "italic", "blockQuote", "link", "|", "bulletedList", "numberedList", "imageUpload", "|", "undo", "redo"],
      //extraPlugins : 'html5video',
      //extraPlugins: [ MyCustomUploadAdapterPlugin ]
      //to set different lang include <script src="/public/js/ckeditor/build/translations/{lang}.js"></scrip> along with core ckeditor script
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
