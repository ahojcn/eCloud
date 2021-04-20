import request from '@/util/requests'

// 获取主机列表
export function apiGetHostList(params) {
    return request({
        url: 'host',
        method: 'get',
        params: params
    })
}

// 删除一台主机
export function apiDeleteHost(host_id) {
    return request({
        url: 'host/' + host_id,
        method: 'delete',
    })
}

// 删除一台主机的用户权限
export function apiDeleteHostUser(params) {
    return request({
        url: 'host_user',
        method: 'delete',
        params: params
    })
}
