import request from '@/util/requests'

// 获取主机列表
export function apiGetRouterInfo(params) {
    return request({
        url: 'm_router',
        method: 'get',
        params: params
    })
}