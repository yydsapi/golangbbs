(function (global, factory) {
  if (typeof define === "function" && define.amd) {
    define('element/locale/zh-CN', ['module', 'exports'], factory);
  } else if (typeof exports !== "undefined") {
    factory(module, exports);
  } else {
    var mod = {
      exports: {}
    };
    factory(mod, mod.exports);
    global.ELEMENT.lang = global.ELEMENT.lang || {}; 
    global.ELEMENT.lang.zhCN = mod.exports;
  }
})(this, function (module, exports) {
  'use strict';

  exports.__esModule = true;
  exports.default = {    
    common: {
        info: '信息',
        tip: '提示',
        confirm: '确定',
        cancel: '取消',
        home: '首页',
        signin:'登陆',
        signup:'注册',
        my:'我的',
        publish:'发布',
        myPublish :'我发布的',
        myReply :'我的回复',
        reply :'回复',
        myBlogConfig:'Blog配置',
        myMenuConfig:'菜单配置',
        myProfile:'个人资料',
        logout :'登出',
        pik :'请输入关键字',
        search:'查找',
        google:'百度',
        clear:'清除',
        title:'标题',
        author:'作者',
        star:'星级',
        reset:'重置',
        man:'男',
        woman:'女',
        sort:'类别',
        upload:'上传',
        share:'分享',
        attachment:'附件',
        content:'内容',
        select:'选择',
        document :'文档',
        documents :' 个文档',
        serverError:'服务器内部错误',
        open:'打开',
        edit:'编缉',
        delete:'删除',
        action:'操作',
        modify:'修改',
        content:'内容',
        time:'时间',
        action:'操作',
        submit:'提交',
        inputerr:'输入错误',
        plif:'请先登陆',
        nopernission:"没有权限",
        ftudb:'更新数据库失败。',
        opfailed:'操作失败。',
        crpe:'内容读取权限错误，没有登录？',
        opsuccess:'操作成功，即将跳转......',
        opfailednologin:'操作失败，没有登陆吗？',
        unknownerr:'未知错误,请联系管理员.',
        plif:'请先登陆',
      },
      manage:{
        usersmanage:'用户管理',
      },
      tips:{
        ondel:'确认删除此文章？',
        ondelreply:'确认删除此回复？',
        ondelnode:'确认删除此节点？(删除后其中文章将一并删除)',
        morethannrde:'超过重试次数或数据错误',
        comfirmdelete:'确认删除',
      },
      blog:{
        openblog:'开启Blog：',
        enkff:'输入关键字进行过滤',
        atfs:'增加一级分类',
        pnode:'父节点',
        sortid:'分类 id',
        sortname:'分类名称',
        sortweight:'排序权重',
        addsubitem:'添加子项',
        plsips:'请输入分类名称',
        snlen:'长度在2~15之间，只能包含中文、字符、数字和下划线和- .',
        plssowe:'请输入分类权重',
        plseaint:'请输入大于1小于10000的正整数',
        
      },
      personalinfo:{
        cuscore:'当前积分：',
        pcsp:'(发布+5分，评论+1分，评分+0.05分)',
        notc:'*密码留空时默认不修改密码',   
        nlaac:'注：Lv5级及以上可上传图片及附件',     
      },
      signup: {
        ucbssp :'用户名不能和密码一样',
        username: '用户名',
        peu :'请输入用户名',
        sex: '性别',
        mailaddress: '邮件地址',
        password: '密码',
        peyp:'请输入密码',
        repeatPassword:'重复密码',
        peypa:'请再次输入密码',
        ata:'同意协议？',
        agreement:'协议',
        agree:'同意',
        clickAndRead :'点击阅读',
        prtr:'＊继续注册前请先阅读注册协议：',
        prtrcon:'<p><BR>欢迎您加入【<b>golangbbs.com</b>】，为维护网上公共秩序和社会稳定，请您自觉遵守以下条款：<BR><BR>\
                            一、不得利用本站危害国家安全、泄露国家秘密，不得侵犯国家社会集体的和公民的合法权益，不得利用本站制作、复制和传播下列信息： <BR>\
                            <BR>\
                            （一）煽动抗拒、破坏宪法和法律、行政法规实施的；<BR>\
                            （二）煽动颠覆国家政权，推翻社会主义制度的；<BR>\
                            （三）煽动分裂国家、破坏国家统一的；<BR>\
                            （四）煽动民族仇恨、民族歧视，破坏民族团结的；<BR>\
                            （五）捏造或者歪曲事实，散布谣言，扰乱社会秩序的；<BR>\
                            （六）宣扬封建迷信、淫秽、色情、赌博、暴力、凶杀、恐怖、教唆犯罪的；<BR>\
                            （七）公然侮辱他人或者捏造事实诽谤他人的，或者进行其他恶意攻击的；<BR>\
                            （八）损害国家机关信誉的；<BR>\
                            （九）其他违反宪法和法律行政法规的；<BR>\
                            （十）进行商业广告行为的。 <BR>\
                            <BR>\
                            二、互相尊重，对自己的言论和行为负责。</p><BR>',
       
        headIcon:'图标',
        psti:'请选择头像',
        createImmediately :'立即创建',
        modifyImmediately :'立即修改',
        pema:'请输入邮件地址',
        plc :'密码长度为6 - 18个字符',
        tipm:'两次输入密码不一致!',
        swll5:'以字母开头，长度在5~15之间，只能包含字符、数字和下划线和.-!',
        pss:'请选择性别',
        pema:'请填写邮件地址',
        petcma:'请输入正确的邮箱地址，邮件地址无法修改',
        swll8:'密码必须以字母开头，长度在8~15之间，只能包含字符、数字和下划线和.-!',
        ambau:"必须同意协议",
        regsuccessinfo:'注册成功，即将跳转...',
        permission:'权限',
        permissiontip:'0：禁用,大于0：启用,100：超级管理员',
        score:'积分：',
        cnbe:'不能为空',
        mbai:'必须为整数',
      },
      signin: {
        signin: '登陆',
        forgotpassword :'忘记密码?',
        signinsuccessinfo:'登陆成功，即将跳转...',
        confirmlogout:'确认登出吗?',
        finepasswordinfo:'临时密码已发送到您的邮箱,请及时登陆并修改,既将跳转到登陆界面...',
        fineerrumn:'找回失败，用户名和邮箱不匹配',
        lfuperr:'登陆失败,用户名或密码错误',
        lfud:'登陆失败,用户名被禁用',
        uerep:'用户名或邮件地址已被注册',
      },
      publish: {
        entertitle: '请输入标题',
        selectsort :'请选择类别',
        isprivate:'是否私有? ',
        clicktoupload:'点击上传',
        publishImmediately :'立即发布',
        selectrestr1 :'当前限制上传三个文档, 已选择',
        selectrestr2 :'文档, 总共选择 ',
        confirmremove :'确定移除 ',
        contentsnbl :'内容不能少于8个字节！',
        publishsucc :'发布成功(LV +5分),既将跳转...',
        replysnbl:'评论不能少于4个字节！',
        plsselectfile:'请选择文件 ',
        wftm:'文件类型不对 ',
        mustbe:'必须是',
        replysuccess:'评论成功(LV +1分),请手动刷新！',
        pfsecuss:'评分成功(LV +0.05分)！',
        titlerep:'标题重复',
        sttsu:'共享给指定的用户,并且不会在主页列表显示,前后用,隔开，例：,a, 或 ,a,b,c,',
      },
      about: {
        content:'<BR><b>[golangbbs] 结构和功能:</b><BR><BR>\
        主要项目结构:golang + gin + vue + element ui + i18n;<BR>\
                            主要插件:ckedit + html5player + blueimp Gallery(music and photo manager)<BR>\
                            数据库:mysql or sqlite;<BR><BR>\
                            特性:<BR>\
                            (一) 所有资源本地化;<BR>\
                            (二) 防注入,文本过滤;<BR>\
                            (三) 功能全面,发布时有比较完善的分享功能和附件,多媒体功能,可以做为知识库,个人电子记事本,个人媒体中心或简单的bbs;<BR>\
                            (四) 可以建立私有博客;<BR>\
                            (五) 菜单和博客分类在线修改;<BR>\
                            (六) 用户资料在线修改;<BR>\
                            (七) 文章和回复在线修改。 <BR>\
                            (八) 基本支持各类移动浏览器 <BR><BR>\
                            使用说明:<BR>\
                            (一) 用户LV5(score>2999)及以上能发布图片,附件,媒体,默认管理用户名:limon,密码:password;<BR>\
                            (二) 可指定多个管理员,用户allow>100的为管理员;<BR>\
                            (三) 已知兼容的版本go 1.10以上,gin v1.4或以上,mysql 5.7 or sqlite 3;<BR><BR>\
                            特别感谢: @fhst, @kdhly, 以及本项目中所使用到的插件和功能模块;以及其它一些未列出的功能模块;<BR><BR>\
                            更多:<BR>\
                            <a href="https://github.com/kdhly/golangbbs/blob/master/READMECN.md"  target=_blank>goto GitHub>></a><BR><BR>'
      },
      pgheader: {
        test: 'test'
      },
      pgcontent: {
        test: 'test'
      },
      pgfooter: {
        test: 'test'
      },
      el: {
      col: {
        test: '确定'
      },
      appdialog1: {
        confirm: '确定',
        clear: '清空'
      },
            colorpicker: {
        confirm: '确定',
        clear: '清空'
      },
      datepicker: {
        now: '此刻',
        today: '今天',
        cancel: '取消',
        clear: '清空',
        confirm: '确定',
        selectDate: '选择日期',
        selectTime: '选择时间',
        startDate: '开始日期',
        startTime: '开始时间',
        endDate: '结束日期',
        endTime: '结束时间',
        prevYear: '前一年',
        nextYear: '后一年',
        prevMonth: '上个月',
        nextMonth: '下个月',
        year: '年',
        month1: '1 月',
        month2: '2 月',
        month3: '3 月',
        month4: '4 月',
        month5: '5 月',
        month6: '6 月',
        month7: '7 月',
        month8: '8 月',
        month9: '9 月',
        month10: '10 月',
        month11: '11 月',
        month12: '12 月',
        // week: '周次',
        weeks: {
          sun: '日',
          mon: '一',
          tue: '二',
          wed: '三',
          thu: '四',
          fri: '五',
          sat: '六'
        },
        months: {
          jan: '一月',
          feb: '二月',
          mar: '三月',
          apr: '四月',
          may: '五月',
          jun: '六月',
          jul: '七月',
          aug: '八月',
          sep: '九月',
          oct: '十月',
          nov: '十一月',
          dec: '十二月'
        }
      },
      select: {
        loading: '加载中',
        noMatch: '无匹配数据',
        noData: '无数据',
        placeholder: '请选择'
      },
      cascader: {
        noMatch: '无匹配数据',
        loading: '加载中',
        placeholder: '请选择',
        noData: '暂无数据'
      },
      pagination: {
        goto: '前往',
        pagesize: '条/页',
        total: '共 {total} 条',
        pageClassifier: '页'
      },
      messagebox: {
        title: '提示',
        confirm: '确定',
        cancel: '取消',
        error: '输入的数据不合法!'
      },
      upload: {
        deleteTip: '按 delete 键可删除',
        delete: '删除',
        preview: '查看图片',
        continue: '继续上传'
      },
      table: {
        emptyText: '暂无数据',
        confirmFilter: '筛选',
        resetFilter: '重置',
        clearFilter: '全部',
        sumText: '合计'
      },
      tree: {
        emptyText: '暂无数据'
      },
      transfer: {
        noMatch: '无匹配数据',
        noData: '无数据',
        titles: ['列表 1', '列表 2'],
        filterPlaceholder: '请输入搜索内容',
        noCheckedFormat: '共 {total} 项',
        hasCheckedFormat: '已选 {checked}/{total} 项'
      },
      image: {
        error: '加载失败'
      },
      pageHeader: {
        title: '返回'
      }
    }
  };
  module.exports = exports['default'];
});