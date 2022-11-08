<template>
	<div>
		<el-button type="primary" style="float:left;margin:10px 0 10px 10px" round @click="getAllVirtualNodes"
			size="mini">
			刷新节点</el-button>
		<el-table ref="multipleTable" :data="tableData" tooltip-effect="dark" style="width: 100%">
			<el-table-column type="selection" width="55">
			</el-table-column>
			<el-table-column label="NodeName" width="120">
				<template slot-scope="scope">
					<el-tooltip class="item" effect="dark" content="scope.row.spec.detail" placement="bottom">
						<i class="el-icon-coin"></i>
					</el-tooltip>
					{{ scope.row.spec.nodename }}
				</template>
			</el-table-column>
			<el-table-column prop="status.created" label="状态" width="120">
				<template slot-scope="scope">
					<i v-if="scope.row.status.created" class="el-icon-success" style="color: green;"></i>
					<i v-else="scope.row.status.created === false" class="el-icon-error" style="color: red;"></i>
				</template>
			</el-table-column>
			<el-table-column prop="status.cpu_usage" label="CPU用量" width="120">
			</el-table-column>
			<el-table-column prop="status.mem_usage" label="内存用量" width="120">
			</el-table-column>
			<el-table-column prop="status.disk_usage" label="磁盘IO" width="120">
			</el-table-column>
		</el-table>
	</div>
</template>

<script>
export default {
	data() {
		return {
			tableData: [],
		};
	},
	mounted() {
		this.getAllVirtualNodes()
	},
	methods: {
		// 页面导航
		page_navigator(page) {
			this.selfLog(page)
			this.$router.push(page);
		},
		getAllVirtualNodes() {
			this.$axios.GetAllVirtualNodes().then(res => {
				this.tableData = res['data']['items']
				console.log(this.tableData)
			})
		},
	}
};
</script>

<style>

</style>
