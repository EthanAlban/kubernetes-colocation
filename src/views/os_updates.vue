<template>
	<div id="">
		<el-table :data="tableData" stripe>
			<el-table-column prop="Subject" label="科目" width="140">
			</el-table-column>
			<el-table-column prop="Position" label="位置" width="120">
			</el-table-column>
			<el-table-column prop="oringinContent" label="原始内容">
			</el-table-column>
			<el-table-column prop="Corrected" label="修正内容">
			</el-table-column>
			<el-table-column prop="Comment" label="备注">
			</el-table-column>
			<el-table-column prop="Time_" label="时间">{{  | formatDate}}
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
		mounted(){
			this.$axios.GetAllUpdates('os').then(res=>{
				console.log(res)
				this.tableData = res.data
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
