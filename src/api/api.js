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
	// =================================================================  NODES =============================================
	GetAllVirtualNodes:()=>{
		return Get('virtual_node/all')
	},
	// =================================================================  KEEPJOB ============================================
	GetAllKeepJobs:()=>{
		return Get('keep_job/all')
	},
	GetAllKeepJobTypes:()=>{
		return Get('keep_job/types/all')
	},
	// =================================================================  CORE ===============================================
	GetAllNamespace:()=>{
		return Get('core/namespace/all')
	}
}
