 
Vue.use(VueI18n) // 通过插件的形式挂载
 
const i18n = new VueI18n({
    locale: 'zh-CN',    // 语言标识
    //this.$i18n.locale // 通过切换locale的值来实现语言切换
    messages: {
      'zh-CN': require('/static/langs/locale/zh.json'),   // 中文语言包
      'en-US': require('/static/langs/locale/en.json')    // 英文语言包
    }
})
 
/* eslint-disable no-new */
new Vue({
  el: '#app',
  i18n,  // 不要忘记
  store,
  router,
  template: '<App/>',
  components: { App }
})