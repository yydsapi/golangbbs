/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "personalpublish/list"}}
{{template "header" .}}
{{template "navbar" .}}
<div class="mycenter" role="main">
<div id="app">  
  <el-form ref="commonForm" :model="commonForm" label-width="100px" id="myForm" class="demo-commonForm" size="small">
    <el-form-item label="">

    </el-form-item>
  </el-form>                
<template>
  <el-table ref="filterTable" :data="tableData" style="width: 100%" :row-class-name="tableRowClassName" @row-click="openDetails">
    <el-table-column prop="img" label=" " align="center" min-width="10%" padding="0px">
      <template slot-scope="tableData"><img :src="tableData.row.img" alt="" style="width: 16px;height:16px"><img :src="tableData.row.imgraise" alt="" style="width: 16px;height:16px"></template>
      </el-table-column>
    <el-table-column prop="title" :label="$t('common.title')" :formatter="formatter" min-width="45%" >
            <template slot-scope="tableData">
                  <span v-html="tableData.row.title"></span>
                </template>
    </el-table-column>
        
    <el-table-column prop="raise_count" :label="$t('common.star')" min-width="20%"  v-if=!{{.IsMobile}} sortable>
      <template slot-scope="tableData">
        <el-rate  :value="tableData.row.raise_count" disabled  text-color="#ff9900" ></el-rate> 
      </template>
    </el-table-column>
    <el-table-column :label="$t('common.action')" min-width="30%">
      <template slot-scope="scope">
        <el-button
          size="mini"
          @click.stop="handleEdit(scope.$index, scope.row)">[[[$t('common.edit')]]]</el-button>
        <el-button
          size="mini"
          type="danger"
          @click.stop="handleDelete(scope.$index, scope.row)">[[[$t('common.delete')]]]</el-button>
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
        window.open("/"+row.link+"/edititem/"+row.id);
      },
      handleDelete(index, row) {
        var ts=this  
        messtype="confirmdeletearticle"
        mess=ts.$t('tips.ondel')
        messid=row.id.toString();
        dg2.dialogVisible2 = true;
      },
      openDetails(row) {
        window.open("/"+row.link+"/showitem/"+row.id);
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
  </script>
</div>
{{template "footer" .}}
{{end}}