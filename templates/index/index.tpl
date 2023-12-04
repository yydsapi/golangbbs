/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "index/list"}}
{{template "header" .}}
{{template "navbar" .}}
<div class="mycenter" role="main">
<div id="app">  
  <el-form ref="commonForm" :model="commonForm" @submit.native.prevent label-width="100px" id="myForm" class="demo-commonForm" size="small">
    <el-form-item label="">
        <el-input ref="searchKey"  v-model="commonForm.searchKey" :placeholder="$t('common.pik')" id="searchKey" class="mysearch"><i slot="prefix" class="el-input__icon el-icon-search"></i></el-input>
        <el-button type="primary" native-type="submit" @click="submitForm('commonForm')">[[[$t('common.search')]]]</el-button>
        <el-button native-type="submit" @click="submitMySearch('commonForm')">[[[$t('common.google')]]]</el-button>
        <el-button @click="resetForm('commonForm')" v-if=!{{.IsMobile}} >[[[$t('common.clear')]]]</el-button>
        <img src="/static/img/write.png" alt="write" style="vertical-align: middle;cursor: pointer; margin-left:15px;width: 23px;height:23px" @click="
        publish('commonForm')" >
    </el-form-item>
  </el-form>                
<template>
  <!--<el-button @click="resetDateFilter">Clear Date Filter</el-button>:show-overflow-tooltip = "true"
  <el-button @click="clearFilter">Clear all filters</el-button> -->
  <el-table ref="filterTable" :data="tableData" style="width: 100%" :row-class-name="tableRowClassName" @row-click="openDetails">
    <el-table-column prop="img" label=" " align="center" min-width="10%" padding="0px">
      <template slot-scope="tableData"><img :src="tableData.row.img" alt="" style="width: 16px;height:16px"><img :src="tableData.row.imgraise" alt="" style="width: 16px;height:16px"></template>
      </el-table-column>
    <el-table-column prop="title" :label="$t('common.title')" :formatter="formatter" min-width="45%" >
      <template slot-scope="tableData">
                  <span v-html="tableData.row.title"></span>
                </template>
    </el-table-column>
        <el-table-column prop="author" :label="$t('common.author')"  sortable min-width="25%" >
          <template slot-scope="tableData">
            <a slot="reference" > </a>
<!--<span v-if="true" >[[[tableData.row.myname]]]</span><span v-else style="color: #37B328" >[[[tableData.row.mytime]]]</span>-->
                  <img :src="tableData.row.userheadicon" alt="" style="width: 26px;height:26px;border:0px"><span>[[[tableData.row.author]]]</span> <img :src="tableData.row.userlevel" alt="" style="width: 20px;height:20px;border:0px"><img :src="tableData.row.privateimg" alt="" style="width: 16px;height:16px;border:0px"><br>
                  <span class="graw" >[[[tableData.row.add_time]]]</span>
                </template>
        </el-table-column>
    <el-table-column prop="raise_count" :label="$t('common.star')" min-width="20%" v-if=!{{.IsMobile}} sortable>
      <template slot-scope="tableData">
        <el-rate  :value="tableData.row.raise_count" disabled  text-color="#ff9900" ></el-rate> 
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
            // Total number of bars, get the length of data according to the interface (note: this can not be empty)
            totalCount:{{.ItemCounts}},
            // Number selector (modifiable)
            pageSizes:[30,60,90,120],
            // Default number of items displayed per page (modifiable)
            PageSize:{{.PageSize}},
                    commonForm: {
          searchKey: '{{.Searchkey}}',
        },
      }
    },
    methods: {
      resetForm(formName){
        try{document.getElementById("searchKey").value="";}catch(e){}
      },
      publish(formName){
        location.href="/bbs/addarticle?category="+getQueryString('category')
      },
      submitMySearch(formName) {
          var searchkey=document.getElementById("searchKey").value.trimall()
          if (searchkey!="") {
                 //location.href="http://www.baidu.com/s?wd="+searchkey
                 if({{.LanguageStr}}.indexOf("cn")==-1){
                  window.open("https://www.google.com/search?q="+searchkey)
                 }else{
                  window.open("https://www.baidu.com/s?wd="+searchkey)
                 }
                 
          } else {
            this.$message({
                      message: this.$t('common.pik'),
                      type: 'success'
                    });
          }
      },
      submitForm(formName) {
          var searchkey=document.getElementById("searchKey").value.trimall()
          if (searchkey!="") {
                 location.href="/?searchkey="+searchkey
          } else {
            this.$message({
                      message: this.$t('common.pik'),
                      type: 'success'
                    });
          }
      },
      openDetails(row) {
        //window.open(row.link+"/showitem/"+row.id);
        location.href=row.link+"/showitem/"+row.id+window.location.search
        },
      resetDateFilter() {
        this.$refs.filterTable.clearFilter('mytime');
      },
      clearFilter() {
        this.$refs.filterTable.clearFilter();
      },
      formattername(row, column) {
        //return row.myname+'\n'+row.mytime;
      },
      formatter(row, column) {
        return row.title
        //return row.title.replaceAll("&lt;","<").replaceAll("&gt;",">");
      },
      filterTag(value, row) {
        return row.tag === value;
      },
      filterHandler(value, row, column) {
        const property = column['property'];
        return row[property] === value;
      },
      tableRowClassName({row, rowIndex}) {
        if (rowIndex === 1) {
          //return 'warning-row';
        } else if (rowIndex === 3) {
          //return 'success-row';
        }
        return '';
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
            // Note: When changing the number of bars displayed on each page, the page number should be displayed on the first page
            this.currentPage=1  
        },
          // Which page is displayed?
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
  </script>
</div>
{{template "footer" .}}
{{end}}