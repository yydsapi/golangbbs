(function (global, factory) {
  if (typeof define === "function" && define.amd) {
    define('element/locale/en', ['module', 'exports'], factory);
  } else if (typeof exports !== "undefined") {
    factory(module, exports);
  } else {
    var mod = {
      exports: {}
    };
    factory(mod, mod.exports);
    global.ELEMENT.lang = global.ELEMENT.lang || {}; 
    global.ELEMENT.lang.en = mod.exports;
  }
})(this, function (module, exports) {
  'use strict';

  exports.__esModule = true;
  exports.default = {
      common: {
        info: 'info',
        tip: 'tips',
        confirm: 'confirm',
        cancel: 'cancel',
        home: 'Home',
        signin:'sign in',
        signup:'sign up',
        my:'my',
        publish:'publish',
        myPublish :'myPublished',
        myReply :'myReplied',
        reply :'reply',
        myBlogConfig:'BlogConfig',
        myMenuConfig:'MenuConfig',
        myProfile:'Profile',
        logout :'logout',
        pik :'please input keyword',
        search:'search',
        google:'google',
        clear:'clear',
        title:'title',
        author:'author',
        star:'star',
        reset:'Reset',
        man:'male',
        woman:'female',
        sort:'sort',
        upload:'upload',
        share:'share',
        attachment:'attachment',
        content:'content',
        select:'select',
        document :'document',
        documents :'documents',
        serverError:'Server internal error',
        open:'open',
        edit:'edit',
        delete:'delete',
        action:'action',
        content:'content',
        action:'action',
        modify:'modify',
        time:'time',
        submit:'submit',
        inputerr:'input error',
        plif:'Please login in first.',
        nopernission:"no permission",
        ftudb:'Failed to operate the database.',
        opfailed:'operation failed.',
        crpe:'Content read permission error, no login?',
        opsuccess:'operation success, the url will be redirected...',
        opfailednologin:'operation failed,need log in?',
        unknownerr:'unknown error,Please contact the administrator.',
      },
      manage:{
        usersmanage:'users manage',
      },
      tips:{
        ondel:'Confirm to delete this article?',
        ondelreply:'Confirm to delete this reply?',
        ondelnode:'Confirm to delete this node? (After deletion, the Related articles will be deleted together)',
        morethannrde:'More than the number of retries or data errors',
        comfirmdelete:'comfirm delete',
      },
      blog:{
        openblog:'open Blog:',
        enkff:'Enter keywords for filtering',
        atfs:'add  the first-level sort',
        pnode:'Parent node',
        sortid:'sort id',
        sortname:'sort name',
        sortweight:'sort weight',
        addsubitem:'addsubitem',
        plsips:'please enter sort name',
        snlen:'Length between 2 and 15, and containing only characters, numbers and underline and - .',
        plssowe:'please enter sort weight',
        plseaint:'Please enter a positive integer greater than 1 and less than 10000',
      },
      personalinfo:{
        cuscore:'Current score:',
        pcsp:'(Publish + 5 Points, Reply + 1 Points, Score + 0.05 Points)',
        notc:'* do not change password if left blank',
        nlaac:'Note: LV5 and above can upload pictures and attachments',
      },
      signup: {
        ucbssp :'username cannot be the same as password',
        username: 'username',
        peu :'please enter your username',
        sex: 'gender',
        mailaddress: 'mailaddress',
        password: 'password',
        peyp:'Please enter your password',
        repeatPassword:'repeatPassword',
        peypa:'Please enter your password again',
        ata:'agree this agreement?',
        agreement:'agreement',
        agree:'agree',
        clickAndRead :'click and Read',
        prtr:'*Please read the registration agreement before you continue to register:',
        prtrcon:'Welcome to join [your website],In order to maintain online public order and social stability, please consciously abide by the following provisions: ... ...<br><br><br>',
        headIcon:'headIcon',
        psti:'please select the icon',
        createImmediately :'Create immediately',
        modifyImmediately :'Modify immediately',
        pema:'Please enter mail address',
        plc :'Password length is 6 - 18 characters',
        tipm:'Twice Input Password Mismatch',
        swll5:'Starting with letters, 5 to 15 in length, and containing only characters, numbers and underline and .-!',
        pss:'Please select sex',
        pema:'Please enter the mailbox address',
        petcma:'Please enter the correct mailbox address. The mailbox address cannot be modified',
        swll8:'the password must starting with letters, 8 to 15 in length, and containing only characters, numbers and underline and .-!',
        ambau:'Agreement must be agreed',
        regsuccessinfo:'Register success, the url will be redirected...',
        permission:'permission',
        permissiontip:'0:disable,more than 0:enable,100:supperadmin',
        score:'score',
        cnbe:'Can not be empty',
        mbai:'Must be an integer',
      },
      signin: {
        signin: 'sign in',
        forgotpassword :'forgot password?',
        signinsuccessinfo:'sign in success, the url will be redirected...',
        confirmlogout:'confirm logout?',
        finepasswordinfo:'The temporary password has been sent to your mailbox. Please login and modify it in time. now will jump to the login interface.',
        fineerrumn:'Failed to retrieve, username and mailbox do not match',
        lfuperr:'Login failure, user name or password error',
        lfud:'Login failed, username disabled',
        uerep:'User name or email address has been registered',
      },
      publish: {
        entertitle: 'Please enter the title',
        selectsort :'Please select the sort',
        isprivate:'isprivate',
        clicktoupload:'click to upload file',
        publishImmediately :'publish Immediately',
        selectrestr1 :'limit upload three documents one time, this time selected ',
        selectrestr2 :' documents, a total of selected ',
        confirmremove :'Confirm remove ',
        contentsnbl :'Content should not be less than 8 characters!',
        publishsucc :'publish success (+ 5 points), the url will be redirected...',
        replysnbl:'Reply content should not be less than 4 characters!',
        plsselectfile:'please select file ',
        wftm:'Wrong file type ',
        mustbe:'Must be',
        replysuccess:'Reply success (+ 1 point), please refresh manually!',  
        pfsecuss:'Score success ( + 0.05 point)!',   
        titlerep:'title repeate',
        sttsu:'Shared to the specified user,and not displayed in the homepage list, separated by, like:,a, or ,a,b,c,',
      },
      about: {
        content:'<BR><b>[golangbbs] Structure and Features:</b><BR><BR>\
        Main project structures: golang + gin + vue + element ui + i18n.<BR>\
                            Main plug-ins: ckedit + html5player + blueimp Gallery (music and photo manager);<BR>\
                            Database: mysql or sqlite.<BR><BR>\
                            1. localization of all resources.<BR>\
                            2. prevent script injections attacks, text filtering.<BR>\
                            3. Open and compact portable design.<BR>\
                            4. This system function was more comprehensive and useful,There are relatively complete sharing functions and attachments when publishing,multimedia capabilities,Can be used as knowledge base and personal electronic notepad and personal media center and simple bbs.<BR>\
                            5. You can set up a private blog.<BR>\
                            6. Menu and blog categories online modification.<BR>\
                            7. User information online modification.<BR>\
                            8. Article and reply online modification.<BR>\
                            9. Basic support for all kinds of mobile browsers. <BR><BR>\
                            Readme:<BR>\
                            1. Users level LV5 (score > 2999) and above can publish pictures, attachments, media, and default management user name:limon,password:password.<BR>\
                            2. Multiple administrators can be specified, and those with user allow > 100 are administrators.<BR>\
                            3. You can adjust some parameters yourself in bbs_config_main_i18n.json file.<BR>\
                            4. BbsUploadPath must have read and write permission.<BR>\
                            5. The image directory you want to display should be in BbsUploadPath+"/Picture/photos", the directory beginning with my is not in the list, but it can still be displayed in the URL.<BR>\
                            6. Known compatible versions go 1.10 or above,gin v1.4 or above,mysql 5.7 or above or sqlite 3.<BR><BR>\
                            Special thanks: @fhst, @kdhly, and all function modules and plug-ins used in the project structure. and other function module not listed.<BR><BR>\
                            More<BR>\
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
        test: 'test'
      },
      appdialog1: {
        confirm: 'confirm',
        clear: 'clear'
      },
      colorpicker: {
        confirm: 'OK',
        clear: 'Clear'
      },
      datepicker: {
        now: 'Now',
        today: 'Today',
        cancel: 'Cancel',
        clear: 'Clear',
        confirm: 'OK',
        selectDate: 'Select date',
        selectTime: 'Select time',
        startDate: 'Start Date',
        startTime: 'Start Time',
        endDate: 'End Date',
        endTime: 'End Time',
        prevYear: 'Previous Year',
        nextYear: 'Next Year',
        prevMonth: 'Previous Month',
        nextMonth: 'Next Month',
        year: '',
        month1: 'January',
        month2: 'February',
        month3: 'March',
        month4: 'April',
        month5: 'May',
        month6: 'June',
        month7: 'July',
        month8: 'August',
        month9: 'September',
        month10: 'October',
        month11: 'November',
        month12: 'December',
        week: 'week',
        weeks: {
          sun: 'Sun',
          mon: 'Mon',
          tue: 'Tue',
          wed: 'Wed',
          thu: 'Thu',
          fri: 'Fri',
          sat: 'Sat'
        },
        months: {
          jan: 'Jan',
          feb: 'Feb',
          mar: 'Mar',
          apr: 'Apr',
          may: 'May',
          jun: 'Jun',
          jul: 'Jul',
          aug: 'Aug',
          sep: 'Sep',
          oct: 'Oct',
          nov: 'Nov',
          dec: 'Dec'
        }
      },
      select: {
        loading: 'Loading',
        noMatch: 'No matching data',
        noData: 'No data',
        placeholder: 'Select'
      },
      cascader: {
        noMatch: 'No matching data',
        loading: 'Loading',
        placeholder: 'Select',
        noData: 'No data'
      },
      pagination: {
        goto: 'Go to',
        pagesize: '/page',
        total: 'Total {total}',
        pageClassifier: ''
      },
      messagebox: {
        title: 'Message',
        confirm: 'OK',
        cancel: 'Cancel',
        error: 'Illegal input'
      },
      upload: {
        deleteTip: 'press delete to remove',
        delete: 'Delete',
        preview: 'Preview',
        continue: 'Continue'
      },
      table: {
        emptyText: 'No Data',
        confirmFilter: 'Confirm',
        resetFilter: 'Reset',
        clearFilter: 'All',
        sumText: 'Sum'
      },
      tree: {
        emptyText: 'No Data'
      },
      transfer: {
        noMatch: 'No matching data',
        noData: 'No data',
        titles: ['List 1', 'List 2'], // to be translated
        filterPlaceholder: 'Enter keyword', // to be translated
        noCheckedFormat: '{total} items', // to be translated
        hasCheckedFormat: '{checked}/{total} checked' // to be translated
      },
      image: {
        error: 'FAILED'
      },

      pageHeader: {
        title: 'Back' // to be translated
      }
    }
  };
  module.exports = exports['default'];
});