/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "users/signin"}}
{{template "header" .}}
<div id="app" class="mycenter"> 
              <div  class="webtitle"><h5>{{.PageTitle}}</h5></div>               
              <el-form :model="commonForm" :rules="rules" ref="commonForm" label-width="130px" id="myForm" @submit.native.prevent class="demo-commonForm">
                <el-form-item :label="$t('signup.username')" prop="myregusername" >
                  <el-input v-model="commonForm.myregusername" :placeholder="$t('signup.peu')"  class="mysigninput"></el-input>
                </el-form-item>
                <el-form-item :label="$t('signup.password')" prop="myregpassword">
                  <el-input type="password" v-model="commonForm.myregpassword" :placeholder="$t('signup.peyp')" class="mysigninput" prefix-icon="el-icon-lock" show-password clearable ></el-input>
                </el-form-item>
                <el-form-item>
                  <el-button  native-type="submit" type="primary" @click="submitForm('commonForm')">[[[$t('signin.signin')]]]</el-button>
                  <el-button @click="gotourl('/bbs/findpassword')">[[[$t('signin.forgotpassword')]]]</el-button>
                </el-form-item>
              </el-form>
              </div>
  <script>
  var Main = {
    delimiters: ['[[[', ']]]'],
    data() {      
      var validatePass = (rule, value, callback) => {
          if (!value) {
            callback(new Error(this.$t('signup.peyp')));
          } else if (value.toString().length < 6 || value.toString().length > 15) {
            callback(new Error(this.$t('signup.plc')))
          } else {
            callback();
          }
        };
      return {
        commonForm: {
          myregusername: '',
          myregpassword: '',
        },
        rules: {
          myregusername: [
            { required: true, message: this.$t('signup.peu'), trigger: 'blur' },
            {pattern: /^[a-zA-Z][a-zA-Z0-9_.!-]{4,14}$/,message: this.$t('signup.swll5')}
          ],
          myregpassword: [
          { required: true, validator: validatePass, trigger: 'blur' },
          {pattern: /^[a-zA-Z][a-zA-Z0-9_.!-]{7,14}$/,message: this.$t('signup.swll8')}
          ],
        }
      }
    },
    methods: {
      submitForm(formName) {
        var ts = this
        this.$refs[formName].validate((valid) => {
          if (valid) {
            var form = document.getElementById("myForm");
            var formData = new FormData();
            for (const key of Object.keys(this.commonForm)) {
                formData.append(key, this.commonForm[key])
              }
              formData.append("Csrf", document.getElementById("_csrf").value)
                  axios.post('/signinpost?return='+getQueryString("return"),formData)
                  .then(function (response) {
                   if(response.data.info=="ok"){
                      mess=ts.$t('signin.signinsuccessinfo');
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
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
            gotourl(url) {
        location.href=url;
      }
    }
  }

var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
</script>
{{template "footer" .}}
{{end}}
