import request from '@/util/requests'

// 获取服务树
export function apiGetMetricsData(params) {
    return request({
        url: 'metrics',
        method: 'get',
        params: params
    })
}
