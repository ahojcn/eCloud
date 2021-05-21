import request from '@/util/requests'

// 获取单个service节点的预案配置
export function apiGetService(params) {
    return request({
        url: 'service/one',
        method: 'get',
        params: params
    })
}

// 删除单个service节点的预案配置
export function apiDeleteService(params) {
    return request({
        url: 'service/delete',
        method: 'get',
        params: params
    })
}

// 创建单个service节点的预案配置
export function apiCreateService(data) {
    return request({
        url: 'service/create',
        method: 'post',
        data
    })
}