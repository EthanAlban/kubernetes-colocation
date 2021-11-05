<template>
	<div id="">
		<el-container style="height: 100%;height: 100%; border: 1px solid #eee">
			<el-aside width="200px" style="background-color: rgb(238, 241, 246)">
				<el-menu :default-openeds="['1', '3']">
					<el-submenu index="1">
						<template slot="title"><i class="el-icon-message"></i>k8sdemo</template>
						<el-menu-item-group>
							<template slot="title">类别</template>
							<el-menu-item index="1-1" @click="page_navigator('/home/pod_exam')">pod例子</el-menu-item>
							<el-menu-item index="1-2" @click="page_navigator('/home/device_handler')">device操作</el-menu-item>
						</el-menu-item-group>
					</el-submenu>
					<!-- <el-submenu index="2">
						<template slot="title"><i class="el-icon-menu"></i>导航二</template>
						<el-menu-item-group>
							<template slot="title">分组一</template>
							<el-menu-item index="2-1">选项1</el-menu-item>
						</el-menu-item-group>
					</el-submenu> -->
				</el-menu>
			</el-aside>

			<el-container>
				<el-header style="text-align: right; font-size: 12px">
					<span>电科KEEP</span>
				</el-header>
				<el-button type="primary" style="position: absolute;bottom: 60px;right: 20px;"  @click="showNewDialog()" icon="el-icon-circle-plus-outline" circle></el-button>

				<el-main style="padding:0">
					<div id="main">
						<router-view />
					</div>
				</el-main>
			</el-container>
		</el-container>

		

	</div>
</template>

<script>
	export default {
		data() {
			return {
				visible: false,
				form: {
					subject: '',
					position: '',
					oringin_content: '',
					corrected: '',
					comment: '',
					qq:''
				},
				formLabelWidth: '120px'
			}
		},
		methods: {
			// 页面导航
			page_navigator(page) {
				this.selfLog(page)
				this.$router.push(page)
			},
			showNewDialog() {
				this.visible = true
			},
			
		}
	}
</script>


<style>
	.el-header {
		background-color: #B3C0D1;
		color: #333;
		line-height: 60px;
	}

	.el-aside {
		color: #333;
	}
</style>
