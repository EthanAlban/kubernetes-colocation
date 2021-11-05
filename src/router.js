import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
	mode: 'history',
	base: process.env.BASE_URL,
	routes: [
		{
			path: '/',
			redirect: '/home/pod_exam'
		},

		{
			path: '/home',
			name: 'home',
			component: () => import('@/views/index.vue'),
			children: [{
				path: '/home/pod_exam',
				name: 'pod_exam',
				component(resolve) {
					require(['@/views/pod_exam.vue'], resolve)
				},
				meta: {
					// 页面标题title
					title: 'pod例子'
				}
				// component: () => import('@/views/Home/Nav1.vue'),
			}, {
				path: '/home/device_handler',
				name: 'device_handler',
				component(resolve) {
					require(['@/views/device_handler.vue'], resolve)
				},
				meta: {
					// 页面标题title
					title: 'device操作'
				}
				// component: () => import('@/views/Home/Nav1.vue'),
			}, ]
		}
	]
})
