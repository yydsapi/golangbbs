/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "personalreply/list"}}
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

    <el-table-column prop="replycontent" :label="$t('common.content')" :formatter="formatter" min-width="50%"  ></el-table-column>
        <el-table-column prop="replytime" :label="$t('common.time')"  sortable  v-if=!{{.IsMobile}}  min-width="20%"></el-table-column>

    <el-table-column :label="$t('common.action')"  min-width="30%" >
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
            // Total number of pieces, get the data length according to the interface (Note: it cannot be empty here)
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
    handleEdit(index, row) {
        window.open("/"+row.link+"/editreplyitem/"+row.id);
      },
      handleDelete(index, row) {
        messtype="confirmdeletereply"
        mess=this.$t('tips.ondelreply')
        messid=row.id.toString();
        dg2.dialogVisible2 = true;
      },
      openDetails(row) {
        window.open("/"+row.link+"/showitem/"+row.pid);
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
        return row.replycontent;
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
        handleSizeChange(val) {
          var urlresult=delQueStr(location.href,"pagesize")
          if(urlresult.indexOf('?')==-1){
            location.href=urlresult+"?pagesize="+val
          }else{
            location.href=urlresult+"&pagesize="+val
          }
            this.PageSize=val
            this.currentPage=1  
        },
        handleCurrentChange(val) {
            var urlresult=delQueStr(location.href,"currentpage")
            this.currentPage=val
          if(urlresult.indexOf('?')==-1){
            location.href=urlresult+"?currentpage="+val
          }else{
            location.href=urlresult+"&currentpage="+val
          }
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