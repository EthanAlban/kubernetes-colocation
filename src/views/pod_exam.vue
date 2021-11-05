<template>
	<div id="">
		<el-table :data="tableData" stripe>
			<el-table-column prop="metadata.generateName" label="generateName" width="140">
			</el-table-column>
			<el-table-column prop="status.hostIP" label="hostIP" width="120">
			</el-table-column>
			<el-table-column prop="spec.nodeName" label="nodeName" width="120">
			</el-table-column>
			<el-table-column prop="status.phase" label="phase">
			</el-table-column>
			<el-table-column prop="status.podIP" label="podIP">
			</el-table-column>
			<el-table-column prop="status.qosClass" label="qosClass">
			</el-table-column>
			<el-table-column prop="status.startTime" label="startTime">{{  | formatDate}}
			</el-table-column>
		</el-table>
	</div>
</template>

<script>
	export default {
		data() {
			return {
				tableData:[],
			};
		},
		mounted(){
			this.$axios.GetPodByPodname('kubeedge-counter-app-6984dcb4dd-8stc2').then(res=>{
				console.log(res)
				var a = new Array();
				a.push(res.data)
				this.tableData = a
			})
		},
		methods: {
			// 页面导航
			page_navigator(page) {
				this.selfLog(page)
				this.$router.push(page);
			},
		}
	};
</script>

<style>
</style>
