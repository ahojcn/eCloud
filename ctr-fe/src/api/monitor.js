import request from '@/util/requests'

// 获取服务树
export function get_metrics_data(params) {
    return request({
        url: 'metrics',
        method: 'get',
        params: params
    })
}
