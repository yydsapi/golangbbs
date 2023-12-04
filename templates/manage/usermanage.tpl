/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "manage/usersmanage"}}
{{template "header" .}}
{{template "navbar" .}}
<div class="mycenter" role="main">
<div id="app">                
<template>
  <el-table ref="filterTable" :data="tableData" style="width: 100%" :row-class-name="tableRowClassName" @row-click="openDetails">
    <el-table-column prop="headicon" label=" " align="center" min-width="10%" padding="0px">
      <template slot-scope="tableData"><img :src="tableData.row.headicon" alt="" style="width: 26px;height:26px"></template>
      </el-table-column>
    <el-table-column prop="myusername" :label="$t('signup.username')" :formatter="formatter" min-width="45%" >
            <template slot-scope="tableData">
                  <span v-html="tableData.row.myusername"></span> <img :src="tableData.row.levelicon" alt="" style="width: 16px;height:16px"><img src='/static/img/space.png' width='5px'><font color='#0956D9'><span v-html="tableData.row.score"></span></font><img src='/static/img/space.png' width='5px'> <font color='#C0C4CC'><br><span v-html="tableData.row.regtime"></span></font>
                </template>
    </el-table-column>
        
    <el-table-column prop="allow" :label="$t('signup.sex')" min-width="10%" sortable>
      <template slot-scope="tableData">
        <span v-html="tableData.row.sex"></span></span>
      </template>
    </el-table-column>
        <el-table-column prop="allow" :label="$t('signup.permission')" min-width="10%" sortable>
      <template slot-scope="tableData">
        <span v-html="tableData.row.allow"></span>
      </template>
    </el-table-column>
    <el-table-column :label="$t('common.action')" min-width="30%">
      <template slot-scope="scope">
        <el-button
          size="mini"
          @click.stop="handleEdit(scope.$index, scope.row)">[[[$t('common.edit')]]]</el-button>
      </template>
    </el-table-column>
  </el-table>
         <div class="tabListPage">
            <el-pagination @size-change="handleSizeChange" 
                           @current-change="handleCurrentChange"
                           :current-page="currentPage" 
                           :page-sizes="pageSizes" 
                           :page-size="PageSize" layout="total, sizes, prev, pager, next, jumper" 
                           :total="totalCount">
              </el-pagination>
        </div>
</template>
</div>
<div id="appdialog3"><el-dialog :visible.sync="dialogVisible3" width="520px" @close="closeDialog('commonForm')" v-if="dialogVisible3" class="DialogForm">
  <el-form :model="commonForm" :rules="rules" ref="commonForm" label-width="120px" id="myForm" size="small" class="demo-commonForm">
                <el-form-item label="cid" prop="cid" style="display: none">
                  <el-input v-model="commonForm.cid"  :placeholder="$t('blog.sortid')" :disabled="true"  id="cid"></el-input>
                </el-form-item>
                <el-form-item label="username" prop="username">
                  <el-input v-model="commonForm.username" :disabled="true"  id="username"></el-input>
                </el-form-item>
                <div style="margin-left:15px;margin-top:-15px;color:#ff7575;font-size: 14px;vertical-align:middle;line-height: 40px;">[[[$t('personalinfo.notc')]]]</div>
                <el-form-item :label="$t('signup.password')" prop="myregpassword">
                  <el-input type="password" v-model="commonForm.myregpassword" placeholder="" prefix-icon="el-icon-lock" show-password clearable ></el-input>
                </el-form-item>
                <el-form-item :label="$t('signup.repeatPassword')" prop="repassword">
                  <el-input type="password" v-model="commonForm.repassword"  placeholder="" prefix-icon="el-icon-lock" show-password clearable ></el-input>
                </el-form-item>
                <el-form-item :label="$t('signup.score')" prop="score" >
                  <el-input v-model="commonForm.score" id="score"></el-input>
                </el-form-item>
                <div style="margin-left:15px;margin-top:-15px;color:#ff7575;font-size: 14px;vertical-align:middle;line-height: 40px;">[[[$t('signup.permissiontip')]]]</div>
                <el-form-item :label="$t('signup.permission')" prop="allow" >
                  <el-input v-model="commonForm.allow" id="allow"></el-input>
                </el-form-item>

                <el-form-item>
                  <el-button type="primary" @click="submitForm('commonForm')">[[[$t('common.submit')]]]</el-button>
                </el-form-item>
    </el-form>
</el-dialog></div>
<style>
  .el-table .warning-row {
    background: oldlace;
  }
  .el-table .success-row {
    background: #f0f9eb;
  }
  .el-table .cell {
    /*text-align: center;*/
    white-space: pre-line;/*Retain line breaks*/
  }
</style>
  <script>
    
var Main = {
  delimiters: ['[[[', ']]]'],
    data() {
      return {
        tableData: JSON.parse({{.Content}}),
          // Which page is displayed by default
            currentPage:{{.CurrentPage}},
            // Total number, get the length of data according to the interface (note: this can not be empty)
            totalCount:{{.ItemCounts}},
            // Number selector (modifiable)
            pageSizes:[30,60,90,120],
            // Default number of items displayed per page (modifiable)
            PageSize:{{.PageSize}},
                    commonForm: {
          searchKey: '',
        },
      }
    },
    methods: {
            submitForm(formName) {

      },
    handleEdit(index, row) {
         dg3.dialogVisible3 = true;
                  setTimeout(function(){ 
                      dg3.$set(dg3.commonForm,'cid', row.id)
                      dg3.$set(dg3.commonForm,'username', row.myusername)
                      dg3.$set(dg3.commonForm,'myregpassword', '')
                      dg3.$set(dg3.commonForm,'regpassword', '')
                      dg3.$set(dg3.commonForm,'allow', row.allow)
                      dg3.$set(dg3.commonForm,'score', row.score)
                },800);
      },
      handleDelete(index, row) {
      },
      openDetails(row) {
        //window.open("/"+row.link+"/showitem/"+row.id);
        },
      resetDateFilter() {
        //this.$refs.filterTable.clearFilter('mytime');
      },
      clearFilter() {
        //this.$refs.filterTable.clearFilter();
      },
      formattername(row, column) {
        //return row.myname+'\n'+row.mytime;
      },
      formatter(row, column) {
        return row.title;
      },
      filterTag(value, row) {
        //return row.tag === value;
      },
      filterHandler(value, row, column) {
        //const property = column['property'];
        //return row[property] === value;
      },
      tableRowClassName({row, rowIndex}) {
        //if (rowIndex === 1) {
          //return 'warning-row';
        //} else if (rowIndex === 3) {
          //return 'success-row';
        //}
        //return '';
      },
        // Number of items displayed per page
        handleSizeChange(val) {
          var urlresult=delQueStr(location.href,"pagesize")
          if(urlresult.indexOf('?')==-1){
            location.href=urlresult+"?pagesize="+val
          }else{
            location.href=urlresult+"&pagesize="+val
          }
            // Change the number of items displayed per page
            this.PageSize=val
            // When you click on the number of items displayed on each page, the first page is displayed.
            //this.getData(val,1)
            // Note: When changing the number of bars displayed on each page, the page number should be displayed on the first page.
            this.currentPage=1  
        },
          // Which page is displayed
        handleCurrentChange(val) {
            // Change the default number of pages
            var urlresult=delQueStr(location.href,"currentpage")
            this.currentPage=val
          if(urlresult.indexOf('?')==-1){
            location.href=urlresult+"?currentpage="+val
          }else{
            location.href=urlresult+"&currentpage="+val
          }
            // When switching page numbers, get the number of bars displayed per page
            //this.getData(this.PageSize,(val)*(this.pageSize))
        },
    },
    created:function(){
          //this.getData(this.PageSize,this.currentPage) 
    }
  }
var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
var dg3 = new Vue({
   delimiters: ['[[[', ']]]'],
      el: '#appdialog3',
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
        dialogVisible3: false,
        commonForm: {
          cid: '',
          username:'',
          myregpassword: '',
          repassword:'',
          score:'',
          allow:'',
        },
        rules: {
          allow: [
            { required: true, message: this.$t('signup.cnbe'), trigger: 'blur' },
            {pattern: /^[0-9]*$/,message: this.$t('signup.mbai')}
          ],
          score: [
            { required: true, message: this.$t('signup.cnbe'), trigger: 'blur' },
            {pattern: /^[0-9]*$/,message: this.$t('signup.mbai')}
          ],
        }
      }
    },
    methods: {
      submitForm(formName) {
        var ts=this
        this.$refs[formName].validate((valid) => {
              if (valid) {
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
                  formData.append("cid",this.commonForm.cid);
                  formData.append("mypassword",mypwd1);
                  formData.append("Csrf", document.getElementById("_csrf").value)
                  formData.append("score",this.commonForm.score);
                  formData.append("allow",this.commonForm.allow);
                        axios.post('/bbs/usersmanagepost',formData)
                        .then(function (response) {
                          if(response.data.info=="ok"){
                            mess=ts.$t('common.opsuccess')
                            dg1.dialogVisible1 = true
                            setTimeout(function(){ location.href="/bbs/usersmanage"; },1000);
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
          } else {
            return false;
          }
        });
      },
 closeDialog(formName) {
    this.dialogVisible3 = false;
    this.$refs[formName].resetFields();
},
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }
    }
    })
  </script>
</div>
{{template "footer" .}}
{{end}}