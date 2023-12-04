/*{{/*! Copyright 2019 golangbbs Core Team.  All rights reserved.
license : use of this source code is governed by AGPL-3.0.
license that can be found in the LICENSE file.*/}}
{{define "users/editblogtree"}}
{{template "header" .}}

<div class="container" role="main" style="justify-content: center;display: flex;margin-top: 30px">
<div id="app" class="appmenutree">
  <el-row><el-col :span=6>[[[$t('blog.openblog')]]]<el-switch
  v-model="switchvalue"
  active-color="#13ce66"
  @change="changeswitchvalue(switchvalue)"
  inactive-color="#ff4949">
</el-switch>
</el-col>
                   <el-col :span=6>
      <div>
      <el-input
        :placeholder="$t('blog.enkff')"
        v-model="filterText"><i slot="prefix" class="el-input__icon el-icon-search"></i>
      </el-input>
    </div></el-col>
<el-col :span=8><el-button type="primary" @click="additem()" style="margin-left: 10px">[[[$t('blog.atfs')]]]</el-button></el-col></el-row>
<div id="treeView" class="menutree">
    <el-tree
      :data="data"
      show-checkbox
      node-key="en"
      default-expand-all
      :props="defaultProps" 
      :expand-on-click-node="false"
      :filter-node-method="filterNode" 
      :render-content="renderContent"
      ref="tree">
    </el-tree>
</div>

</div>
<div id="appdialog3"><el-dialog :visible.sync="dialogVisible3" width="460px" @close="closeDialog('commonForm')" v-if="dialogVisible3" class="DialogForm">
  <el-form :model="commonForm" :rules="rules" ref="commonForm" label-width="100px" id="myForm" size="small" class="demo-commonForm">
            <el-form-item :label="$t('blog.pnode')" prop="pid"  style="display: none">
                  <el-input v-model="commonForm.pid" :placeholder="$t('blog.pnode')" id="pid" :disabled="true"></el-input>
                </el-form-item>
                <el-form-item label="cid" prop="cid" style="display: none">
                  <el-input v-model="commonForm.cid"  :placeholder="$t('blog.sortid')" :disabled="true"  id="cid"></el-input>
                </el-form-item>
                        <el-form-item :label="$t('blog.sortname')" prop="categoryname" >
                  <el-input v-model="commonForm.categoryname" :placeholder="$t('blog.sortname')"  id="categoryname"></el-input>
                </el-form-item>
                        <el-form-item :label="$t('blog.sortweight')" prop="weight" >
                  <el-input v-model="commonForm.weight" :placeholder="$t('blog.sortweight')" id="weight"></el-input>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="submitForm('commonForm')">[[[$t('common.submit')]]]</el-button>
                </el-form-item>
    </el-form>
</el-dialog></div>
</div>
<script type="text/javascript">
  var Main = {
    delimiters: ['[[[', ']]]'],
      watch: {
      filterText(val) {
      this.$refs.tree.filter(val)
    }
  },
    data() { 
    try{    
        return {
          filterText: '',
          switchvalue:{{.SessionUseBlog}},
          data: JSON.parse({{.Content}})[0][0].children,
          defaultProps: {
          children: 'children',
          label: 'name',
          //id: 'id'
          }
        }
      }catch(e){
            return {
              filterText: '',
              switchvalue:{{.SessionUseBlog}},
              data: [],
              defaultProps: {},
      }
              //data: JSON.parse(JSON.stringify(data))
      }
    },
    methods: {
      changeswitchvalue(data){
          if(data){
              switchbbsvalue("commonupdatebyid","updateuserblog","1")
          }else{
              switchbbsvalue("commonupdatebyid","updateuserblog","0")
          }
      },
      append(data) {
        const newChild = { en: "en", label: 'ts', children: [] };
        if (!data.children) {
          this.$set(data, 'children', []);
        }
        data.children.push(newChild);
      },
      filterNode(value, data) {
        if (!value) return true;
        return data.name.indexOf(value) !== -1
      },
      remove(node, data) {
        const parent = node.parent;
        const children = parent.data.children || parent.data;
        const index = children.findIndex(d => d.en === data.en);
        children.splice(index, 1);
      },
      additem() {
            dg3.dialogVisible3 = true;
            setTimeout(function(){
                dg3.$set(dg3.commonForm,'cid', "")
                dg3.$set(dg3.commonForm,'pid', "blog")
          },800);
      },

      renderContent(h, { node, data, store }) {
        
          return h('span', {
          style: {
            //color: "red",
          },
          on: {
            //'mouseenter': () => {
              'mouseover': () => {
                
              data.is_show = true;
            },
            //'mouseleave': () => {
              'mouseout': () => {

              data.is_show = false;
            }
          }
        }, [
          h('span', {
                }, node.label),
          h('span', {  class: 'treemenusmallfont',
                }, " ("+ data.sort_weight+")"),
          h('span', {
            style: {position: 'absolute',right: 0,},
          }, [
            h('el-button', {
              props: {
                type: 'text',
                size: 'small',
              },
              style: {marginLeft:"15px",},
              on: {
                click: () => {
                  dg3.dialogVisible3 = true;
                  setTimeout(function(){
                    dg3.$set(dg3.commonForm,'cid', "")
                    dg3.$set(dg3.commonForm,'pid', data.en)
                },1000);
                }
              }
            }, this.$t('blog.addsubitem')),
            h('el-button', {
              props: {
                type: 'text',
                size: 'small',
              },
              style: {

              },
              on: {
                click: () => {
                  dg3.dialogVisible3 = true;
                  setTimeout(function(){ 
                    document.getElementById("pid").value=data.pid
                    document.getElementById("cid").value=data.id
                    dg3.$set(dg3.commonForm,'categoryname', data.name)
                    dg3.$set(dg3.commonForm,'weight', data.sort_weight)
                  //document.getElementById("categoryname").value=data.name
                },1000);
                }
              }
            }, this.$t('common.modify')),
            h('el-button', {
              props: {
                type: 'text',
                size: 'small',
              },
              style: {

              },
              on: {
                click: () => {
                  messtype="confirmdeletethismenuitem"
                  mess=this.$t('tips.ondelnode')
                  dg2.dialogVisible2 = true;
                  messid=data.id
                  messeffct="blog"
                }
              }
            },  this.$t('common.delete')),
                ]),
        ]);

      }
    }
  };
var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
var dg3 = new Vue({
   delimiters: ['[[[', ']]]'],
      el: '#appdialog3',
      data() {
          return {
            dialogVisible3: false,
            commonForm: {
              pid: '',
              categoryname: '',
              weight: '500',
              cid:'',
            },
            rules: {
              categoryname: [
                { required: true, message: this.$t('blog.plsips'), trigger: 'blur' },
                {pattern: /^[A-Za-z0-9_\-\ \.\&\u4e00-\u9fa5]{2,32}$/,message: this.$t('blog.snlen')}
              ],
              weight: [
              { required: true, message: this.$t('blog.plssowe'), trigger: 'blur' },
              {pattern: /^(?!0)(?:[0-9]{1,4}|10000)$/,message: this.$t('blog.plseaint')}
              ],
            }
          }
    },
    methods: {
      submitForm(formName) {
        var ts=this
        this.$refs[formName].validate((valid) => {
          if (valid) {
            var formData = new FormData();
            formData.append("pid", document.getElementById("pid").value)
            formData.append("Csrf", document.getElementById("_csrf").value)
            formData.append("categoryname", document.getElementById("categoryname").value)
            formData.append("weight", document.getElementById("weight").value)
            formData.append("cid", document.getElementById("cid").value)
            formData.append("dbid", "blog")
                  axios.post('/bbs/personalmenutreepost?return='+getQueryString("return"),formData)
                  .then(function (response) {
                   if(response.data.info=="ok"){
                      mess=ts.$t('common.opsuccess')
                      dg1.dialogVisible1 = true;
                      setTimeout(function(){ location.href=delQueStr(location.href,"t")+"?t="+new Date().getTime(); },1000);
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

<style>
  .custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
    padding-right: 8px;
  }
</style>

{{template "footer" .}}
{{end}}