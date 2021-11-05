let env = 'dev' //设置全局的console是否生效的条件,开发情况下允许console
export function selfLog(x) {
	if (env === 'dev') {
		console.log(x)
	}
}



export function UrlEncode(str) {
	var res = ''
	var strSpec = '!"#$%&\'()*+,/:;<=>?[]^`{|}~%'
	var text = ''

	for (var i = 0; i < str.length; i++) {
		var chr = str.charAt(i)
		let val = chr
		chr = ('0' + val.charCodeAt(0).toString(16)).slice(-2)
		var c = chr
		text += chr + ':' + c + 'n'
		if (parseInt('0x' + c) > 0x7f) {
			res += '%' + c.slice(0, 2) + '%' + c.slice(-2)
		} else {
			if (chr === ' ') res += '+'
			else if (strSpec.indexOf(chr) != -1) res += '%' + c.toString(16)
			else res += chr
		}
	}
	console.log(res, 'res')
	return res
}

export function str2ASCII(val) {
	return ('0' + val.charCodeAt(0).toString(16)).slice(-2)
}

export function encodeUtf8(text) {
	const code = encodeURIComponent(text)
	const bytes = []
	for (var i = 0; i < code.length; i++) {
		const c = code.charAt(i)
		if (c === '%') {
			const hex = code.charAt(i + 1) + code.charAt(i + 2)
			const hexVal = parseInt(hex, 16)
			bytes.push(hexVal)
			i += 2
		} else bytes.push(c.charCodeAt(0))
	}
	return bytes
}

export function decodeUtf8(bytes) {
	var encoded = ''
	for (var i = 0; i < bytes.length; i++) {
		encoded += '%' + bytes[i].toString(16)
	}
	return decodeURIComponent(encoded)
}
