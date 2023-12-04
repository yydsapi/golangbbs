/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "users/usersinfo"}}
{{template "header" .}}
        <h4 class="card-title">{{.Title}}</h4>
              <div id="app"  class="mycenter">
              <el-form :model="commonForm" ref="commonForm" @submit.native.prevent label-width="130px">
                <el-form-item :label="$t('signup.username')" prop="myregusername" >
          {{.Userid}} <img src="{{.UserLevel}}" alt="" style="width: 26px;height:26px;border:0px;position:relative;top:7px;"> <span style="position:relative;left:17px;">[[[$t('personalinfo.cuscore')]]]{{.UsersScore}} [[[$t('personalinfo.pcsp')]]]<br><div style="margin-left:1px;margin-top:-5px;color:#ff7575;font-size: 14px;vertical-align:middle;line-height: 40px;">[[[$t('personalinfo.nlaac')]]]</div><img src="/static/img/level/1.png" alt="lv1" style="width: 22px;height:22px;border:0px;position:relative;top:5px;">:0 <img src="/static/img/level/2.png" alt="lv2" style="width: 22px;height:22px;border:0px;position:relative;top:5px;">:300 <img src="/static/img/level/3.png" alt="lv3" style="width: 22px;height:22px;border:0px;position:relative;top:5px;">:900 <img src="/static/img/level/4.png" alt="lv4" style="width: 22px;height:22px;border:0px;position:relative;top:5px;">:1800 <img src="/static/img/level/5.png" alt="lv5" style="width: 22px;height:22px;border:0px;position:relative;top:5px;">:3000 <img src="/static/img/level/6.png" alt="lv6" style="width: 22px;height:22px;border:0px;position:relative;top:5px;">:5000 <img src="/static/img/level/7.png" alt="lv7" style="width: 22px;height:22px;border:0px;position:relative;top:5px;">:7000 <img src="/static/img/level/8.png" alt="lv8" style="width: 22px;height:22px;border:0px;position:relative;top:5px;">:10000 <img src="/static/img/level/9.png" alt="lv9" style="width: 22px;height:22px;border:0px;position:relative;top:5px;">:40000 </span>
                </el-form-item>
                <el-form-item :label="$t('signup.mailaddress')" prop="myemail" >
                    {{.Email}}
                </el-form-item>
                  <el-form-item :label="$t('signup.sex')" prop="mysex">
                  <el-radio-group v-model="commonForm.mysex">
                    <el-radio :label="$t('common.man')" :value="$t('common.man')"></el-radio>
                    <el-radio :label="$t('common.woman')" :value="$t('common.woman')"></el-radio>
                  </el-radio-group>
                </el-form-item>
                <div style="margin-left:15px;margin-top:-15px;color:#ff7575;font-size: 14px;vertical-align:middle;line-height: 40px;">[[[$t('personalinfo.notc')]]]</div>
                <el-form-item :label="$t('signup.password')" prop="myregpassword">
                  <el-input type="password" v-model="commonForm.myregpassword" placeholder="" class="mysigninput" prefix-icon="el-icon-lock" show-password clearable ></el-input>
                </el-form-item>
                <el-form-item :label="$t('signup.repeatPassword')" prop="repassword">
                  <el-input type="password" v-model="commonForm.repassword"  placeholder="" class="mysigninput" prefix-icon="el-icon-lock" show-password clearable ></el-input>
                </el-form-item>


<div >

                <el-form-item :label="$t('signup.headIcon')" prop="avatars">
                <el-select v-model="commonForm.avatars" :placeholder="$t('signup.psti')" @change="changeSelection" ref="selecticon">
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
                  <el-button type="primary" native-type="submit" @click.stop="submitForm('commonForm')">[[[$t('signup.modifyImmediately')]]]</el-button>
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

          } else if (value.toString().length < 6 || value.toString().length > 15) {
            callback(new Error(this.$t('signup.plc')))
          } else {
            //callback();
          }
        };
         var validatePass2 = (rule, value, callback) => {
                  if (value === '') {
                  } else if (value !== this.commonForm.myregpassword) {
                  callback(new Error(this.$t('signup.tipm')));
                  } else {
                  //callback();
                  }
              };
      return {
        commonForm: {
          mysex: '{{.Sex}}',
          myusername:'{{.Userid}}',
          avatars: '{{.Headicon}}',
        },
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
        rules: {
          myregpassword: [
          { required: false, validator: validatePass, trigger: 'blur' },
          {pattern: /^[a-zA-Z][a-zA-Z0-9_.!-]{7,14}$/,message: this.$t('signup.swll8')}
          ],
          repassword: [{ required: false, validator: validatePass2, trigger: 'blur' }],
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
         var ts=this;
        var mypwd1=""
        var mypwd2=""
        var intcheck=0
        if(typeof(this.commonForm.myregpassword)!="undefined"){
          mypwd1=this.commonForm.myregpassword
        }
        if(typeof(this.commonForm.repassword)!="undefined"){
          mypwd2=this.commonForm.repassword
        }
        if (mypwd1!="" || mypwd2!=""){
          var zz=/^[a-zA-Z][a-zA-Z0-9_.!-]{7,14}$/
          if(zz.test(mypwd1)&&zz.test(mypwd2)){
            intcheck=1
          }else{
            this.$message({
                message: this.$t('signup.swll8'),
                type: 'warning'
              });
          }
        }else{
          intcheck=1
        }
        if(intcheck==1){
            if(this.commonForm.myusername==mypwd1){
               this.$message({
               message: this.$t('signup.ucbssp'),
               type: 'warning'
                });
               return false;
            }
            if(mypwd1!=mypwd2){
                     this.$message({
                         message: this.$t('signup.tipm'),
                         type: 'warning'
                      });
                     return false;
            }
          var formData = new FormData();
            formData.append("mysex",this.commonForm.mysex);
            formData.append("mypassword",mypwd1);
            formData.append("Csrf", document.getElementById("_csrf").value)
            formData.append("myicon",this.commonForm.avatars);
                  axios.post('/bbs/personalinfopost',formData)
                  .then(function (response) {
                    if(response.data.info=="ok"){
                      mess=ts.$t('common.opsuccess')
                      dg1.dialogVisible1 = true
                      setTimeout(function(){ location.href="/"; },1000);
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
        }
      },
      resetForm(formName) {
        location.href=location.href;
      }
    }
  }

var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
</script>
{{template "footer" .}}
{{end}}
