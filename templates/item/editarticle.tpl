/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "bbs/personal/editarticle"}}
{{template "header" .}}
<div id="app" class="mycenter">
<div  class="webtitle"><h5>{{.PageTitle}}</h5></div> 
<el-form :model="commonForm" @submit.native.prevent :rules="rules" ref="commonForm" label-width="100px" id="myForm" class="demo-commonForm" size="medium">
<el-form-item :label="$t('common.title')" prop="mytitle" >
     <el-col :span="12"><el-input v-model="commonForm.mytitle"  :placeholder="$t('publish.entertitle')"></el-input></el-col>
</el-form-item>
<el-row>
   <el-col :span="6">
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
  <el-col :span="8">
        <el-form-item :label="$t('publish.isprivate')" prop="isprivate">                        
          <el-switch
            v-model="commonForm.isprivate"
            active-color="#13ce66"
            inactive-color="#DEDEDE">
          </el-switch>
      </el-form-item>
  </el-col>
</el-row>

    {{if ge .IntAdmin 5}}
    <el-form-item label="Media" prop="ck-mediaurl" >
      <input type="file" name="image" id="file_upload">
      <el-button type="primary" @click="submitMedia('commonForm')">[[[$t('common.upload')]]]</el-button>
      <el-input v-model="commonForm.mediaurl" placeholder="" id="mediaurl" style="width:40%;"></el-input>
    </el-form-item>
    <div style="margin-left:95px;margin-top:-15px;color:#ff7575;font-size: 12px;vertical-align:middle;line-height: 30px;">[[[$t('publish.sttsu')]]]</div>
      <el-form-item :label="$t('common.share')" prop="strshare">                        
    <el-input v-model="commonForm.strshare" id="strshare" style="width:40%;margin-left:10px"></el-input>
    </el-form-item>
     {{end}} {{if gt .UserLevel 3}}
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
            :data="CsrfStr" 
            :on-exceed="handleExceed"
            :file-list="fileList">
            <el-button size="small" type="primary">[[[$t('publish.clicktoupload')]]]</el-button>
            <!--<div slot="tip" class="el-upload__tip">Only upload JPG / png files, and no more than 500 KB</div>-->
      </el-upload></div>
     </el-form-item>
                {{end}}
      <el-form-item :label="$t('common.content')" prop="ck-content" >
        <textarea name="content" id="ck-content" style="height:400px">{{.Content}}</textarea>
      </el-form-item>
      <el-form-item :label="$t('common.sort')" prop="categorytmp" style="display: none;">
          <el-radio-group v-model="commonForm.categorytmp">
          <el-radio :label="$t('common.man')" :value="$t('common.man')"></el-radio>
          <el-radio :label="$t('common.woman')" :value="$t('common.woman')"></el-radio>
        </el-radio-group>
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
  var flist=JSON.parse({{.FileList}})
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
          mytitle: '{{.Title}}',
          category: {{.Categoryenall}}.split(","),
          mediaurl: '{{.Media}}',
          isprivate: {{.isprivate}},
          strshare:'{{.ReaderStr}}',
        },
        rules: {
          mytitle: [
            { required: true, message: this.$t('publish.entertitle'), trigger: 'blur' },
          ],
          category: [{ required: true, message:this.$t('publish.selectsort'), trigger: 'change' }],
         
        },
        fileList: JSON.parse({{.FileList}}),
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
        return this.$confirm(this.$t('tips.comfirmdelete')+`${ file.name } ?`);
      },
      handleChangeUpload(file, fileList) {
        this.fileList = fileList.slice(-3);
      },
      handleChange(value) {
        console.log(value);

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
        if (this.$refs['CateGoryStr'].value[1]==null) {
            this.$message({
                                  message: this.$t('publish.selectsort'),
                                  type: 'success'
                                });
            return;
        }
        this.$refs[formName].validate((valid) => {
          if (valid) {
            
                        var flisttmp=""
            try{
                for(var i=0;i<flist.length;i++){
                  if (flisttmp==""){
                    if (flist[i].response){
                      flisttmp=flist[i].response.url
                    }else{
                      flisttmp=flist[i].url
                    }
                    
                  }else{
                    if (flist[i].response){
                      flisttmp+="?"+flist[i].response.url
                    }else{
                      flisttmp+="?"+flist[i].url
                    }
                  }
                }
      }catch(e){flisttmp=""}
            var form = document.getElementById("myForm");
            var formData = new FormData();
            for (const key of Object.keys(this.commonForm)) {
                formData.append(key, this.commonForm[key])
              }
             var categoryens=this.$refs['CateGoryStr'].value
            formData.append("attachment",flisttmp);
            try{formData.append("strshare", document.getElementById("strshare").value)}catch(e){formData.append("strshare", "")}
            formData.append("categoryen",categoryens[categoryens.length-1]);
            formData.append("categoryenall",this.$refs['CateGoryStr'].value);
            formData.append("categorycnall",this.$refs['CateGoryStr'].currentLabels);
            formData.append("content",editor.getData());
            formData.append("Csrf", document.getElementById("_csrf").value)
            formData.append("id",{{.Id}});
            try{
              var mediaurl=document.getElementById("mediaurl").value
              formData.append("mediaurltmp",mediaurl);}catch(e){}
                  axios.post('/bbs/updatearticlepost',formData)
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
                      type: 'success'
                    });
          return;
        }
        var ts=this
        var extname=file_upload.toLowerCase().split('.').splice(-1)
        if (extname=="mp3" || extname=="ogg" || extname=="mp4" || extname=="mkv" || extname=="avi") {
          var fileM=document.getElementById("file_upload")
          var fileObj = fileM.files[0];
          var formData = new FormData();
          formData.append("Csrf", document.getElementById("_csrf").value)
          formData.append('upload', fileObj);
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
                      type: 'success'
                    });
          }
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
</script>
{{template "footer" .}}
{{end}}
