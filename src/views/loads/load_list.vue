<template>
	<div>
		<el-button type="primary" style="float:left;margin:10px 0 10px 10px" round @click="getAllKeepJobs"
			size="mini">
			刷新负载</el-button>
		<el-table ref="multipleTable" :data="tableData" tooltip-effect="dark" style="width: 100%">
			<el-table-column type="selection" width="55">
			</el-table-column>
			<el-table-column label="apiVersion"  prop="apiVersion" width="120"></el-table-column>
			<el-table-column prop="kind" label="任务类型" width="120"> </el-table-column>
			<el-table-column prop="spec.job_name" label="任务名" width="120"></el-table-column>
			<el-table-column prop="spec.namespace" label="命名空间" width="120"></el-table-column>
			<el-table-column prop="spec.replica" label="并行任务数" width="120"></el-table-column>
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
		this.getAllKeepJobs()
	},
	methods: {
		getAllKeepJobs() {
			this.$axios.GetAllKeepJobs().then(res => {
				this.tableData = res['data']['items']
				console.log(this.tableData)
			})
		},
	}
};
</script>

<style>

</style>
