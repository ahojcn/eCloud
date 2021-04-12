import request from '@/util/requests'

// 登录
export function login(data) {
    return request({
        url: 'session',
        method: 'post',
        data
    })
}

// 登出
export function logout() {
    return request({
        url: 'session',
        method: 'delete',
    })
}

// 判断是否登录
export function is_login() {
    return request({
        url: 'session',
        method: 'get',
    })
}
