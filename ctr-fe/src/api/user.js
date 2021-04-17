import request from '@/util/requests'

// 注册
export function register(data) {
    return request({
        url: 'user',
        method: 'post',
        data
    })
}

// 根据用户名或邮箱获取用户列表
export function apiGetUserListByUsername(params) {
    return request({
        url: 'user',
        method: 'get',
        params: params
    })
}