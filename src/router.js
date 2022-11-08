import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
	mode: 'history',
	base: process.env.BASE_URL,
	routes: [
		{
			path: '/',
			redirect: '/home/node_list'
		},

		{
			path: '/home',
			name: 'home',
			component: () => import('@/views/index.vue'),
			children: [{
				path: '/home/node_list',
				name: 'node_list',
				component(resolve) {
					require(['@/views/node_list.vue'], resolve)
				},
				meta: {
					// 页面标题title
					title: '集群节点'
				}
				// component: () => import('@/views/Home/Nav1.vue'),
			}, 
			{
				path: '/home/loads/load_list',
				name: 'load_list',
				component(resolve) {
					require(['@/views/loads/load_list.vue'], resolve)
				},
				meta: {
					// 页面标题title
					title: '负载列表'
				},
			},
			{
				path: '/home/loads/create_load',
				name: 'create_load',
				component(resolve) {
					require(['@/views/loads/create_load.vue'], resolve)
				},
				meta: {
					// 页面标题title
					title: '创建负载'
				},
			},
		]
		}
	]
})
