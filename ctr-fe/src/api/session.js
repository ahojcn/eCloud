import request from '../util/requests'

// 登录
export function login(data) {
    return request({
        url: 'session',
        method: 'post',
        data
    })
}


export function getYuQueOAuth(query) {
    return request({
        url: 'user',
        method: 'get',
        params: query
    })
}
