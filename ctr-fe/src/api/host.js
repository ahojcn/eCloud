import request from '@/util/requests'

// 添加主机
export function apiAddHost(data) {
    return request({
        url: 'host',
        method: 'post',
        data
    })
}

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

// 给一个用户添加主机使用权限
export function apiAddHostUser(data) {
    return request({
        url: 'host_user',
        method: 'post',
        data
    })
}

// Shell 执行命令
export function apiRunCommand(data) {
    return request({
        url: 'command',
        method: 'post',
        data
    })
}
