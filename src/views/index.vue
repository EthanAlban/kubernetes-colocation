<template>
	<div id="">
		<el-container style="height: 100%;height: 100%; border: 1px solid #eee">
			<el-aside width="200px" style="background-color: rgb(238, 241, 246)">
				<el-menu :default-openeds="['1', '3']">
					<el-submenu index="1">
						<template slot="title"><i class="el-icon-message"></i>题目订正</template>
						<el-menu-item-group>
							<template slot="title">科目</template>
							<el-menu-item index="1-1" @click="page_navigator('/home/os_updates')">操作系统</el-menu-item>
							<el-menu-item index="1-2" @click="page_navigator('/home/ds_updates')">数据结构</el-menu-item>
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
					<span>电科cs龙王</span>
				</el-header>
				<el-button type="primary" style="position: absolute;bottom: 60px;right: 20px;"  @click="showNewDialog()" icon="el-icon-circle-plus-outline" circle></el-button>

				<el-main style="padding:0">
					<div id="main">
						<router-view />
					</div>
				</el-main>
			</el-container>
		</el-container>

		<el-dialog title="提交更正" :visible.sync="visible">
			<el-form :model="form">
				<el-form-item label="科目" :label-width="formLabelWidth">
					<el-select v-model="form.subject" placeholder="请选择科目">
						<el-option label="操作系统" value="os"></el-option>
						<el-option label="数据结构" value="ds"></el-option>
					</el-select>
				</el-form-item>

				<el-form-item label="位置" :label-width="formLabelWidth">
					<el-input v-model="form.position" autocomplete="off"></el-input>
				</el-form-item>

				<el-form-item label="原始内容" :label-width="formLabelWidth">
					<el-input v-model="form.oringin_content" autocomplete="off"></el-input>
				</el-form-item>

				<el-form-item label="更正内容" :label-width="formLabelWidth">
					<el-input v-model="form.corrected" autocomplete="off"></el-input>
				</el-form-item>
				<el-form-item label="备注" :label-width="formLabelWidth">
					<el-input v-model="form.comment" autocomplete="off"></el-input>
				</el-form-item>
				<el-form-item label="QQ" :label-width="formLabelWidth">
					<el-input v-model="form.qq" autocomplete="off"></el-input>
				</el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button @click="dialogFormVisible = false">取 消</el-button>
				<el-button type="primary" @click="uploadForm">确 定</el-button>
			</div>
		</el-dialog>

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
			uploadForm(){
				this.$axios.InsertNewUpdate({
					'Position':this.form.position,
					'OringinContent':this.form.oringin_content,
					'Corrected':this.form.corrected,
					'Comment':this.form.comment,
					'Subject':this.form.subject,
					'QQ':this.form.qq,
				}).then(res=>{
					console.log(res)
					this.visible = false
					 this.$message({
					          message: '提交成功，审核后将会展示，感谢支持~',
					          type: 'success'
					        })
				})
			}
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
