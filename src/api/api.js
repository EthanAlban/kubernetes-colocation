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
	GetPodByPodname:podname=>{
		return Get('pod/get_pod_by_podname?podname='+podname)
	},
	// InsertNewUpdate:params=>{
	// 	console.log(params)
	// 	return Post('insert_new_update',params)
	// }
}
