const path = require('path')
const CompressionPlugin = require('compression-webpack-plugin')

function resolve(dir) {
	return path.join(__dirname, dir) // path.join(__dirname)设置绝对路径
}

module.exports = {
	// 启动压缩
	configureWebpack: {
		plugins: [
			new CompressionPlugin({
				test: /\.(js|css)(\?.*)?$/i, //需要压缩的文件正则
				threshold: 10240, //文件大小大于这个值时启用压缩
				deleteOriginalAssets: false //压缩后保留原文件
			})
		]
	},
	// 别名的配置
	chainWebpack(config) {
		config.module
			.rule('svg')
			.exclude.add(resolve('src/icons'))
			.end()
		config.module
			.rule('icons').test(/\.svg$/)
			.include.add(resolve('src/icons'))
			.end()
			.use('svg-sprite-loader')
			.loader('svg-sprite-loader')
			.options({ symbolId: 'icon-[name]' })
			.end()
	},
	chainWebpack: config => {
		config.resolve.alias
			.set('@', resolve('./src'))
			.set('components', resolve('./src/components'))
			.set('views', resolve('./src/views'))
			.set('assets', resolve('./src/assets'))
			.set('styles', resolve('./src/styles'))
			.set('api', resolve('./src/api'))
		// set第一个参数：设置的别名，第二个参数：设置的路径
	},
	// 设置网站icon
	pwa: {
		iconPaths: {
			favicon32: './code.png',
			favicon16: './code.png',
			appleTouchIcon: './code.png',
			maskIcon: './code.png',
			msTileImage: './code.png'
		}
	},
	// 配置跨域请求
	devServer: {
		port: 8080,
		host: '0.0.0.0',
		open: true,
		https: false,
		open: true,
		proxy: {
			// oj服务器
			'/api': {
				// target: 'http://192.168.10.63:30001',
				target: 'http://192.168.1.128:6666',
				// target: 'http://192.168.128.0:8080/',
				ws: true,
				changeOrigin: true,
				pathRewrite: { '^/api': '/api' }
			},
			'/captcha': {
				// target: 'http://192.168.10.63:30002',
				target: 'http://192.168.128.0:8666',
				ws: true,
				changeOrigin: true,
				pathRewrite: { '^/captcha': '/captcha' }
			},
			'/open-apis': {
				target: 'https://open.feishu.cn',
				ws: true,
				changeOrigin: true,
				pathRewrite: { '^/open-apis': '/open-apis' }
			},
			// 学校的API http://studyapi.uestc.edu.cn
			'/ckd': {
				target: 'http://studyapi.uestc.edu.cn',
				ws: true,
				changeOrigin: true,
				pathRewrite: { '^/ckd': '/ckd' }
			},
			'uestcapi': {
				target: 'http://feishumiddle.uestc.edu.cn',
				ws: true,
				changeOrigin: true,
				pathRewrite: { '^/uestcapi': '/uestcapi' }
			}
		}
	}
}
