import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
	mode: 'history',
	base: process.env.BASE_URL,
	routes: [{
			path: '/test',
			name: '/test',
			component(resolve) {
				require(['@/views/test.vue'], resolve)
			},
			meta: {
				// 页面标题title
				title: '电科820订正'
			}
			// component: () => import('@/views/Login.vue')
		},
		{
			path: '/',
			redirect: '/home/os_updates'
		},

		{
			path: '/home',
			name: 'home',
			component: () => import('@/views/index.vue'),
			children: [{
				path: '/home/os_updates',
				name: 'os_updates',
				component(resolve) {
					require(['@/views/os_updates.vue'], resolve)
				},
				meta: {
					// 页面标题title
					title: 'os订正'
				}
				// component: () => import('@/views/Home/Nav1.vue'),
			}, {
				path: '/home/ds_updates',
				name: 'ds_updates',
				component(resolve) {
					require(['@/views/ds_updates.vue'], resolve)
				},
				meta: {
					// 页面标题title
					title: 'ds订正'
				}
				// component: () => import('@/views/Home/Nav1.vue'),
			}, ]
		}
	]
})
