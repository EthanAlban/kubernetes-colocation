import {
	Get,
	Post
} from '@/api/request'
import {
	selfLog,
	UrlEncode
} from '../utils'
import axios from 'axios'
export default {
	NewCaptId: () => {
		return Get('new')
	},
	// 鉴别验证码
	VerifyCaptcha: params=>{
		return Get('verify',params)
	},
	GetAllUpdates:subject=>{
		return Get('get_all_updates?subject='+subject)
	},
	InsertNewUpdate:params=>{
		console.log(params)
		return Post('insert_new_update',params)
	}
}
