/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "users/findpassword"}}
{{template "header" .}}
              <div id="app" class="mycenter">
                <h5 class="card-title">{{.Title}}</h5>
              <el-form :model="commonForm" :rules="rules" @submit.native.prevent ref="commonForm" label-width="100px" id="myForm" class="demo-commonForm">
                <el-form-item :label="$t('signup.username')" prop="myregusername" >
                  <el-input v-model="commonForm.myregusername"  :placeholder="$t('signup.peu')"  class="mysigninput"></el-input>
                </el-form-item>
                <el-form-item :label="$t('signup.mailaddress')" prop="myemail" >
                  <el-input v-model="commonForm.myemail" :placeholder="$t('signup.pema')"  class="mysigninput"></el-input>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" native-type="submit" @click="submitForm('commonForm')">[[[$t('common.confirm')]]]</el-button>
                  <el-button @click="resetForm('commonForm')">[[[$t('common.reset')]]]</el-button>
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
          myemail: '',
        },
        rules: {
          myregusername: [
            { required: true, message: this.$t('signup.peu'), trigger: 'blur' },
            {pattern: /^[a-zA-Z]\w{4,14}$/,message:  this.$t('signup.swll5')}
          ],
          myemail: [
            { required: true, message: this.$t('signup.pema'), trigger: 'blur' },
            {type: 'email',message: this.$t('signup.petcma'),trigger: ['blur', 'change']}
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
            formData.append("Csrf", document.getElementById("_csrf").value)
            for (const key of Object.keys(this.commonForm)) {
                formData.append(key, this.commonForm[key])
              }
                  axios.post('/bbs/findpasswordpost',formData)
                  .then(function (response) {
                   if(response.data.info=="ok"){
                      mess=ts.$t('signin.finepasswordinfo');
                      dg1.dialogVisible1 = true;
                      setTimeout(function(){ location.href='/signin'; },5000);
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
      }
    }
  }

var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
</script>
{{template "footer" .}}
{{end}}
