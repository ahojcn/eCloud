import request from '@/util/requests'

// 获取服务树
export function apiGetTreeInfo(params) {
    return request({
        url: 'tree',
        method: 'get',
        params: params
    })
}

// 添加服务树节点
export function apiAddTreeNode(data) {
    return request({
        url: 'tree',
        method: 'post',
        data
    })
}

// 标记删除节点
export function apiDeleteTree(params) {
    return request({
        url: 'tree',
        method: 'delete',
        params: params
    })
}

// 给用户添加服务树节点权限
export function apiAddUserTree(data) {
    return request({
        url: 'user_tree',
        method: 'post',
        data
    })
}

// 删除用户的服务树节点权限
export function apiDeleteUserTree(params) {
    return request({
        url: 'user_tree',
        method: 'delete',
        params: params
    })
}
