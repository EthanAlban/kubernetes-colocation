<template>
	<div>
		{{ this.active }}
		<!-- <el-drawer title="启动任务" :visible.sync="showProcess" :direction="direction" :before-close="handleClose" size="60%"> -->
		<el-steps :active="active" finish-status="success" simple style="margin-top: 20px;">
			<el-step title="步骤 1"></el-step>
			<el-step title="步骤 2">b</el-step>
			<el-step title="步骤 3">c</el-step>
			<el-step title="步骤 4">d</el-step>
			<el-step title="步骤 5">e</el-step>
		</el-steps>
		<div v-if="active === 0">
			<el-form ref="form" :model="form" label-width="80px" style="margin-top: 20px">
				<el-form-item label="负载名" style="width: 20vw;">
					<el-input v-model="form.keepJObName"></el-input>
				</el-form-item>
				<el-form-item label="命名空间" style="width: 20vw;">
					<el-select v-model="selectedNamespace" placeholder="请选择">
						<el-option v-for="ns in namespaces" :key="ns.metadata.name" :label="ns.metadata.name"
							:value="ns.metadata.name">
						</el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="负载镜像" style="width: 35vw;">
					<el-input v-model="form.Image"></el-input>
				</el-form-item>
				<el-form-item label="负载类型" style="width: 20vw;">
					<el-select v-model="selectedType" placeholder="请选择">
						<el-option v-for="typed in Jobtypes" :key="typed" :label="typed"
							:value="typed">
						</el-option>
					</el-select>
				</el-form-item>
				<!-- <el-form-item>
					<el-button type="primary" >立即创建</el-button>
					<el-button>取消</el-button>
				</el-form-item> -->
			</el-form>
		</div>
		<el-button style="margin-top: 12px;" @click="pre">上一步</el-button>
		<el-button style="margin-top: 12px;" @click="next">下一步</el-button>
		<!-- </el-drawer> -->
	</div>
</template>
<script>
export default {
	data() {
		return {
			namespaces: [],
			selectedNamespace: "",
			Jobtypes: [],
			selectedType:"",
			showProcess: false,
			direction: 'btt',
			active: 0,
			form: {
				keepJObName: '',
				Namespace: this.selectedNamespace,
				Image: '',
				date2: '',
				delivery: false,
				type: [],
				resource: '',
				desc: ''
			},
		};
	},
	mounted() {
		this.$axios.GetAllNamespace().then(res => {
			console.log(res['data']['items'])
			this.namespaces = res['data']['items']
		})
		this.$axios.GetAllKeepJobTypes().then(res => {
			console.log(res)
			this.Jobtypes = res['data']
		})
	},
	methods: {
		createNewjob() {
			this.showProcess = true
		},
		handleClose(done) {
			this.$confirm('确认关闭？')
				.then(_ => {
					done();
				})
				.catch(_ => { });
		},
		next() {
			console.log(this.active)
			if (this.active++ > 6) this.active = 0;
		},
		pre() {
			if (this.active > 0) {
				this.active--
			}
		}
	}
};
</script>

<style>

</style>
