import request from '@/util/requests'

// 获取接入层信息
export function apiGetRouterInfo(params) {
    return request({
        url: 'm_router',
        method: 'get',
        params: params
    })
}

// 创建接入层
export function apiRouterDeploy(data) {
    return request({
        url: 'm_router',
        method: 'post',
        data
    })
}

// 重新部署接入层
export function apiRouterRedo(data) {
    return request({
        url: 'm_router/redo',
        method: 'post',
        data
    })
}

// 获取接入层状态
export function apiGetRouterStatus(params) {
    return request({
        url: 'm_router/status',
        method: 'get',
        params: params
    })
}