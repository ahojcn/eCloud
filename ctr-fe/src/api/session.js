import request from '@/util/requests'

// 登录
export function apiLogin(data) {
    return request({
        url: 'session',
        method: 'post',
        data
    })
}

// 登出
export function apiLogout() {
    return request({
        url: 'session',
        method: 'delete',
    })
}

// 判断是否登录
export function apiIsLogin() {
    return request({
        url: 'session',
        method: 'get',
    })
}
