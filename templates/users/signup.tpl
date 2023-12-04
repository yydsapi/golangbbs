/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "users/signupget"}}
{{template "header" .}}
<div id="app"  class="mycenter">
                <div  class="webtitle"><h5>{{.PageTitle}}</h5></div>
              <el-form :model="commonForm" @submit.native.prevent :rules="rules" ref="commonForm" label-width="130px">
                <el-form-item :label="$t('signup.username')" prop="myregusername"  class="mysigninput">
                  <el-input v-model="commonForm.myregusername"  :placeholder="$t('signup.peu')"></el-input>
                </el-form-item>

                  <el-form-item :label="$t('signup.sex')" prop="mysex">
                  <el-radio-group v-model="commonForm.mysex">
                    <el-radio :label="$t('common.man')" :value="$t('common.man')"></el-radio>
                    <el-radio :label="$t('common.woman')" :value="$t('common.woman')"></el-radio>
                  </el-radio-group>
                </el-form-item>

                <el-form-item :label="$t('signup.mailaddress')" prop="myemail" >
                  <el-input v-model="commonForm.myemail"  :placeholder="$t('signup.pema')"  class="mysigninput"></el-input>
                </el-form-item>

                <el-form-item :label="$t('signup.password')" prop="myregpassword">
                  <el-input type="password" v-model="commonForm.myregpassword" :placeholder="$t('signup.peyp')"  class="mysigninput" prefix-icon="el-icon-lock" show-password clearable ></el-input>
                </el-form-item>

                <el-form-item :label="$t('signup.repeatPassword')" prop="repassword">
                  <el-input type="password" v-model="commonForm.repassword"  :placeholder="$t('signup.peypa')"  class="mysigninput" prefix-icon="el-icon-lock" show-password clearable ></el-input>
                </el-form-item>
                <el-row>
                   <el-col :span="8">
                    <el-form-item :label="$t('signup.ata')" prop="type">
                    <el-checkbox-group v-model="commonForm.type">
                      <el-checkbox :label="$t('signup.agree')" name="type"></el-checkbox>
                    </el-checkbox-group>
                  </el-form-item>
                   </el-col>
                   <el-col :span="8">
                    <el-link onclick="showhide('xydiv')" class="xydiv">[[[$t('signup.clickAndRead')]]]</el-link>
                   </el-col>
                </el-row>
          <div>
              
              <div id="xydiv" style="display:none;margin:10 auto"><BR>
                        <b>[[[$t('signup.prtr')]]]</b>
                        <span v-html="commonForm.prtrcon"></span>                           
                        </div>
                <el-form-item :label="$t('signup.headIcon')" prop="myicon">
                <el-select v-model="commonForm.avatars" :placeholder="$t('signup.psti')" @change="changeSelection" ref="selecticon" id="myFormavatars">
                  <el-option
                    v-for="avatar in avatars"
                    :key="avatar.Id"
                    :label="avatar.src"
                    :value="avatar.Id"
                  >
                    <img class="avatar" :src="avatar.src" style="height:36px">
                  </el-option>
                </el-select>
              </el-form-item>
                <el-form-item>
                  <el-button type="primary" native-type="submit" @click.stop="submitForm('commonForm')">[[[$t('signup.createImmediately')]]]</el-button>
                  <el-button @click="resetForm('commonForm')">[[[$t('common.reset')]]]</el-button>
                </el-form-item>
              </el-form>
              </div>
          </div>
          </div>
  <script>
    var ts=this;
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
         var validatePass2 = (rule, value, callback) => {
                  if (value === '') {
                  callback(new Error(this.$t('signup.peypa')));
                  } else if (value !== this.commonForm.myregpassword) {
                  callback(new Error(this.$t('signup.tipm')));
                  } else {
                  callback();
                  }
              };
      return {
                avatars: [
          {
            src:"/static/headicon/1.png",
            Id:'/static/headicon/1.png',
          },{
            src:"/static/headicon/2.png",
            Id:'/static/headicon/2.png',
          },{
            src:"/static/headicon/3.png",
            Id:'/static/headicon/3.png',
          },{
            src:"/static/headicon/4.png",
            Id:'/static/headicon/4.png',
          },{
            src:"/static/headicon/5.png",
            Id:'/static/headicon/5.png',
          },{
            src:"/static/headicon/6.png",
            Id:'/static/headicon/6.png',
          },{
            src:"/static/headicon/7.png",
            Id:'/static/headicon/7.png',
          },{
            src:"/static/headicon/8.png",
            Id:'/static/headicon/8.png',
          },{
            src:"/static/headicon/9.png",
            Id:'/static/headicon/9.png',
          },{
            src:"/static/headicon/10.png",
            Id:'/static/headicon/10.png',
          },{
            src:"/static/headicon/11.png",
            Id:'/static/headicon/11.png',
          },{
            src:"/static/headicon/12.png",
            Id:'/static/headicon/12.png',
          },{
            src:"/static/headicon/13.png",
            Id:'/static/headicon/13.png',
          },{
            src:"/static/headicon/14.png",
            Id:'/static/headicon/14.png',
          },{
            src:"/static/headicon/15.png",
            Id:'/static/headicon/15.png',
          },{
            src:"/static/headicon/16.png",
            Id:'/static/headicon/16.png',
          },
        ],
        value: '',
        commonForm: {
          myregusername: '',
          mysex: this.$t('common.man'),
          prtrcon:this.$t('signup.prtrcon'),
          avatars: '/static/headicon/1.png',
          myemail: '',
          myregpassword: '',
          repassword: '',
          type: [],
        },
        rules: {
          myregusername: [
            { required: true, message: this.$t('signup.peu'), trigger: 'blur' },
            {pattern: /^[a-zA-Z][a-zA-Z0-9_.!-]{4,14}$/,message: this.$t('signup.swll5')}
          ],
          mysex: [{ required: true, message: this.$t('signup.pss'), trigger: 'change' }],
          myemail: [
            { required: true, message: this.$t('signup.pema'), trigger: 'blur' },
            {type: 'email',message: this.$t('signup.petcma'),trigger: ['blur', 'change']}
          ],
          myregpassword: [
          { required: true, validator: validatePass, trigger: 'blur' },
          {pattern: /^[a-zA-Z][a-zA-Z0-9_.!-]{7,14}$/,message: this.$t('signup.swll8')}
          ],
          repassword: [{ required: true, validator: validatePass2, trigger: 'blur' }],
          type: [{ type: "array", required: true, message: this.$t('signup.ambau'), trigger: "change"}],
        }
      }
    },
    methods: {
        changeSelection(){ 
          this.$nextTick(function(){
            var tss=this
            setTimeout(function(){ 
                let path=tss.$refs.selecticon.selectedLabel.replace("headicon","headicon/mini")
                //alert(path)
                tss.$refs.selecticon.$el.children[0].children[1].setAttribute('style','background:url('+path+') no-repeat;');
              },500);
           })
          
        },
      submitForm(formName) {
        var ts=this
        this.$refs[formName].validate((valid) => {
          if (valid) {
            if(this.commonForm.myregusername==this.commonForm.myregpassword){
               this.$message({
               message: this.$t('signup.ucbssp'),
               type: 'warning'
                });
               return false;
            }
var formData = new FormData();
                formData.append("myusername", this.commonForm.myregusername)
                formData.append("Csrf", document.getElementById("_csrf").value)
                formData.append("myemail", this.commonForm.myemail)
                formData.append("mysex", this.commonForm.mysex)
                formData.append("mypassword", this.commonForm.myregpassword)
                formData.append("myicon", this.commonForm.avatars)
                  axios.post('/signuppost',formData)
                  .then(function (response) {
                   if(response.data.info=="ok"){
                      mess=ts.$t('signup.regsuccessinfo');
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
      }
    }
  }

var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
            setTimeout(function(){ 
                document.getElementById("myFormavatars").nextElementSibling.setAttribute('style','background:url('+'/static/headicon/mini/1.png'+') no-repeat;');
              },1000);
</script>
{{template "footer" .}}
{{end}}
