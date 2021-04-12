import request from '@/util/requests'

// 注册
export function register(data) {
    return request({
        url: 'user',
        method: 'post',
        data
    })
}