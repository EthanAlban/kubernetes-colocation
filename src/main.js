import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import dataV from '@jiaminghi/data-view'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import { vueBaberrage } from 'vue-baberrage'


import Api from './api/api'


import VueCookies from 'vue-cookies'
import 'default-passive-events'
// 引入懒加载
import VueLazyload from 'vue-lazyload'


// import more codemirror resource...

// you can set default global options and events when Vue.use


Vue.use(VueCookies)

Vue.prototype.$axios = Api

// 全局的this.selfLog替换
// import selfLog from './utils/selfLog';
import { selfLog } from './utils'
import { decodeUtf8 } from './utils'
import { encodeUtf8 } from './utils'

Vue.prototype.selfLog = selfLog
Vue.prototype.decodeUtf8 = decodeUtf8
Vue.prototype.encodeUtf8 = encodeUtf8


import VueClipboard from 'vue-clipboard2'
import contribution from 'vue-contribution'
import 'vue-contribution/dist/vue-contribution.css'
Vue.use(contribution)
Vue.use(VueClipboard)
Vue.use(vueBaberrage)
Vue.use(dataV)
Vue.use(ElementUI)

Vue.config.productionTip = false

import Moment from 'moment'

// 定义全局时间戳过滤器
Vue.filter('formatDate', function(value) {
	return Moment(value).format('YYYY-MM-DD HH:mm:ss')
})

router.beforeEach((to, from, next) => {
  /* 路由发生变化修改页面title */
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})

new Vue({
	
	router,
	store,
	data: function() {
		return {
			sys_status: 1,
			setInterval_id: -1,
			user_profile: null,
		}
	},


	render: h => h(App)
}).$mount('#app')
