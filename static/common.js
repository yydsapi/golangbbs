Vue.config.delimiters = ['[[[', ']]]']
var ts=this
var mess=""
var messtype=""
var messid=""
var messeffct=""
function confirmLogout(){
    mess=this.Vue.tc('signin.confirmlogout')
    messtype="confirmlogout"
    dg2.dialogVisible2 = true
}
function doConfirm(x,y) {
  var ts=this
    switch (x) {
    case "confirmlogout":
        axios.get('/logout', {
            params: {}
        }).then(function(response) {
            if (response.data.info == "ok") {
                mess=this.Vue.tc('common.opsuccess')
                setTimeout(function() {
                    location.href = "/";
                },
                1000);
            }
            dg1.dialogVisible1 = true

        }).catch(function(error) {
            mess=this.Vue.tc('common.serverError')
            dg1.dialogVisible1 = true;
            console.log(error);
        });
        break;
    case "confirmdeletearticle":
        var formData = new FormData();
        formData.append("id", y)
        formData.append("Csrf", document.getElementById("_csrf").value)
          axios.post('/bbs/deletearticlepost',formData)
          .then(function (response) {
            if(response.data.info=="ok"){
              mess=this.Vue.tc('common.opsuccess')
              dg1.dialogVisible1 = true;
              setTimeout(function(){ location.href=location.href; },1000);
            }
            else{
              mess=this.Vue.tc('common.opfailednologin')
              dg1.dialogVisible1 = true;
            }
          })
          .catch(function (error) {
            mess=this.Vue.tc('common.serverError')
            dg1.dialogVisible1 = true;
            console.log(error);
          });
        break;
    case "confirmdeletereply":
        var formData = new FormData();
        formData.append("id", y)
        formData.append("Csrf", document.getElementById("_csrf").value)
          axios.post('/bbs/deletereplypost',formData)
          .then(function (response) {
            if(response.data.info=="ok"){
              mess=this.Vue.tc('common.opsuccess')
              dg1.dialogVisible1 = true;
              setTimeout(function(){ location.href=location.href; },1000);
            }
            else{
              mess=this.Vue.tc('common.opfailednologin')
              dg1.dialogVisible1 = true;
            }
          })
          .catch(function (error) {
            mess=this.Vue.tc('common.serverError')
            dg1.dialogVisible1 = true;
            console.log(error);
          });
        break;
     case "confirmdeletethismenuitem":
          var formData = new FormData();
          formData.append("id", y)
          formData.append("dbid", messeffct)
          formData.append("Csrf", document.getElementById("_csrf").value)
            axios.post('/bbs/personaldeletemenutree',formData)
            .then(function (response) {
              if(response.data.info=="ok"){
                mess=this.Vue.tc('common.opsuccess')
                dg1.dialogVisible1 = true;
                setTimeout(function(){ location.href=location.href; },1000);
              }
              else{
                mess=this.Vue.tc('common.opfailednologin')
                dg1.dialogVisible1 = true;
              }
            })
            .catch(function (error) {
              mess=this.Vue.tc('common.serverError')
              dg1.dialogVisible1 = true;
              console.log(error);
            });
        break;
    default:
        mess=this.Vue.tc('common.unknownerr')
        dg1.dialogVisible1 = true;
        break;
    }
}
function showhide(x) {
	var objdiv=document.getElementById(x);
  if (objdiv.style.display=="none"){
  	   objdiv.style.display="block";
  	}else{
  	   objdiv.style.display="none";
  	}
}
String.prototype.replaceAll = function(s1,s2){ 
  return this.replace(new RegExp(s1,"gm"),s2); 
}

String.prototype.trimall = function(){
  return this.replace(/^[ 　]*(.*?)[ 　]*$/g,'$1');
}

function setCookie(name,value){
  var Days = 1000;
  var exp = new Date(); 
  exp.setTime(exp.getTime() + Days*24*60*60*1000);
  document.cookie = name + "="+ escape (value) + ";expires=" + exp.toGMTString();
}

function getCookie(name){
  var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
  if(arr=document.cookie.match(reg)) return unescape(arr[2]);
  else return "";
}

function delCookie(name){
  var exp = new Date();
  exp.setTime(exp.getTime() - 1);
  var cval=getCookie(name);
  if(cval!=null) document.cookie= name + "="+cval+";expires="+exp.toGMTString();
}

function getQueryString(paras) {
  var url = window.location.search;
    var paraString = url.substring(url.indexOf("?") + 1, url.length).split("&");
    var paraObj = {}
    for (i = 0; j = paraString[i]; i++) {
        paraObj[j.substring(0, j.indexOf("=")).toLowerCase()] = j.substring(j.indexOf("=") + 1, j.length);
    }
    var returnValue = paraObj[paras.toLowerCase()];
    if (typeof (returnValue) == "undefined") {
        return "";
    } else {
        return returnValue;
    }
}
function delQueStr(url, ref) //Delete parameter
{
  var str = "";

  if (url.indexOf('?') != -1)
      str = url.substr(url.indexOf('?') + 1);
  else
      return url;
  var arr = "";
  var returnurl = "";
  var setparam = "";
  if (str.indexOf('&') != -1) {
    arr = str.split('&');
    for (i in arr) {
        if (arr[i].split('=')[0] != ref) {
            returnurl = returnurl + arr[i].split('=')[0] + "=" + arr[i].split('=')[1] + "&";
        }
    }
    return url.substr(0, url.indexOf('?')) + "?" + returnurl.substr(0, returnurl.length - 1);
  }
  else {
    arr = str.split('=');
    if (arr[0] == ref)
        return url.substr(0, url.indexOf('?'));
    else
        return url;
  }
}
function switchbbsvalue(url,sort,id)
{
  var formData = new FormData();
  formData.append("id", id)
  formData.append("sort", sort)
  formData.append("Csrf", document.getElementById("_csrf").value)
    axios.post(url,formData)
    .then(function (response) {
      if(response.data.info=="ok"){
        mess=this.Vue.tc('common.opsuccess')
        dg1.dialogVisible1 = true;
        setTimeout(function(){ location.href=location.href; },1000);
      }
      else{
        mess=this.Vue.tc('common.opfailednologin')
        dg1.dialogVisible1 = true;
      }
    })
    .catch(function (error) {
      mess=this.Vue.tc('common.serverError')
      dg1.dialogVisible1 = true;
      console.log(error);
    });
}
/*
function gotourl(url,sort)
{
    if (sort=="1"){
        location.href=url;
    }else if(sort=="2"){
      window.open(url)  
    }else{
        location.href=location.href
    }
}
*/
