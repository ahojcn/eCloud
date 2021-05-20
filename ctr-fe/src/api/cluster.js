import request from '@/util/requests'

// 获取单个集群配置
export function apiClusterRetrieve(params) {
    return request({
        url: 'cluster/one',
        method: 'get',
        params: params
    })
}

// 删除单个集群配置
export function apiClusterDelete(params) {
    return request({
        url: 'cluster/delete',
        method: 'get',
        params: params
    })
}

// 创建集群配置
export function apiClusterCreate(data) {
    return request({
        url: 'cluster/create',
        method: 'post',
        data
    })
}

// /cluster/list
// 获取集群配置列表
export function apiClusterList(params) {
    return request({
        url: 'cluster/list',
        method: 'get',
        params: params
    })
}