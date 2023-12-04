    var dg1 = new Vue({
 delimiters: ['[[[', ']]]'],
      el: '#appdialog1',
      data: function() {
        return { dialogVisible1: false }
      },
      methods: {
            open () {
                this.$nextTick(() => document.getElementById("dialogcontent1").innerHTML=mess)
            }
        }
    })
    var dg2 = new Vue({
       delimiters: ['[[[', ']]]'],
      el: '#appdialog2',
      data: function() {
        return { dialogVisible2: false }
      },
      methods: {
            open () {
                this.$nextTick(() => document.getElementById("dialogcontent2").innerHTML=mess)
            },
            appdialog2Confirm () {
                doConfirm(messtype,messid)
            }
        }
    })
  try{
        var appdivider = new Vue({
          delimiters: ['[[[', ']]]'],
      el: '#appdivider',
      data: function() {
        return {}
      },
      methods: {
            open () {
                //this.$nextTick(() => document.getElementById("dialogcontent1").innerHTML=mess)
            }
        }
    })}catch(e){}
    try{
        var apprate = new Vue({
          delimiters: ['[[[', ']]]'],
      el: '#apprate',
      data: function() {
        return {value5: 3.7}
      }
    })}catch(e){}


    try{
        var vmmenu = new Vue({
          delimiters: ['[[[', ']]]'],
    el: '#appmenu',
    data() {
      return {
        activeIndex: '1'
      };
    },
    methods: {
      handleSelect(key, keyPath) {
        console.log(key, keyPath);
      }
    }
})
    }catch(e){}
    //try{new Vue().$mount('#headmenu')}catch(e){}
        try{
         // new Vue().$mount('#headphoto')
         new Vue({
        el: '#appcarousel',
        data () {
            return {
              BannerImg: [{
                  item: '/static/upload/bannerpic/banner1.jpg',
                  index: 0
                },{
                  item: '/static/upload/bannerpic/banner2.jpg',
                  index: 1
                },{
                  item: '/static/upload/bannerpic/banner3.jpg',
                  index: 2
                },{
                  item: '/static/upload/bannerpic/banner4.jpg',
                  index: 3
                }]
            }
          },
        mounted () {
            //this.setSize();
            //const that = this;
            window.addEventListener('resize', function() {
              //that.screenWidth = window.width();
             // that.setSize();
            }, false);
        },
        methods:{
          setSize: function () {
            this.bannerHeight = 740 / 2560 * this.screenWidth
            if(this.bannerHeight > 740) this.bannerHeight = 740
            if(this.bannerHeight < 360) this.bannerHeight = 360
          }
        }
    })
 }catch(e){}
 
