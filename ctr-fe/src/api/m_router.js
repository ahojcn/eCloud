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