import request from '@/util/requests'

// 获取服务树
export function get_metrics_data(query) {
    return request({
        url: 'metrics',
        method: 'get',
        params: query
    })
}
