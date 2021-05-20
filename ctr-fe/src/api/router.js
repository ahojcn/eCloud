import request from '@/util/requests'

// 获取接入层指标（un, uri）
export function apiGetRouterMetrics(params) {
    return request({
        url: 'router/metrics',
        method: 'get',
        params: params
    })
}

// 获取接入层监控信息
export function apiQueryRouterMetrics(params) {
    return request({
        url: 'router/query',
        method: 'get',
        params: params
    })
}
