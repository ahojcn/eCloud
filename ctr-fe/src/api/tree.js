import request from '@/util/requests'

// 获取服务树
export function get_tree() {
    return request({
        url: 'tree',
        method: 'get',
    })
}
