<template>
	<el-container>
		<el-header style="text-align: right; font-size: 12px;">
			<div><img src="../assets/template.png" style="height: 50px;position: absolute;left: 15px;top: 15px;"></img>
			</div>
			<!-- <el-dropdown>
				<i class="el-icon-setting" style="margin-right: 15px"></i> -->
			<!-- <el-dropdown-menu slot="dropdown">
					<el-dropdown-item>查看</el-dropdown-item>
					<el-dropdown-item>新增</el-dropdown-item>
					<el-dropdown-item>删除</el-dropdown-item>
				</el-dropdown-menu>
			</el-dropdown>
			<span>王小虎</span> -->
		</el-header>

		<el-container style=" border: 1px solid #eee">
			<el-aside width="200px" style="height:100vh;background-color:#545c64">
				<el-menu default-active="2" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose"
					background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
					<el-submenu index="1">
						<template slot="title">
							<i class="el-icon-location"></i>
							<span>集群节点</span>
						</template>
						<el-menu-item index="1-1" @click="page_navigator('/home/node_list')">节点列表</el-menu-item>
					</el-submenu>
					<el-submenu index="2">
						<template slot="title">
							<i class="el-icon-menu"></i>
							<span>集群负载</span>
						</template>
						<el-menu-item index="2-1" @click="page_navigator('/home/loads/load_list')">负载列表</el-menu-item>
						<el-menu-item index="2-2" @click="page_navigator('/home/loads/create_load')">创建负载</el-menu-item>
					</el-submenu>
					<el-menu-item index="3" disabled>
						<i class="el-icon-document"></i>
						<span slot="title">导航三</span>
					</el-menu-item>
					<el-menu-item index="4">
						<i class="el-icon-setting"></i>
						<span slot="title">导航四</span>
					</el-menu-item>
				</el-menu>
			</el-aside>

			<el-main>
				<div id="main">
					<router-view />
				</div>
			</el-main>
		</el-container>
	</el-container>
</template>
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
  
<script>
export default {
	data() {
		return {
			tableData: [],
		}
	},
	methods: {
		handleOpen(key, keyPath) {
			console.log(key, keyPath);
		},
		handleClose(key, keyPath) {
			console.log(key, keyPath);
		},

		// 页面导航
		page_navigator(page) {
			this.selfLog(page)
			this.$router.push(page);
		}
	}
};
</script>