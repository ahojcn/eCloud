import request from '@/util/requests'

// 获取接入层信息
export function apiGetPipeLineList(params) {
    return request({
        url: 'pipeline/list',
        method: 'get',
        params: params
    })
}

// 创建接入层
export function apiCreatePipeLine(data) {
    return request({
        url: 'pipeline/create',
        method: 'post',
        data
    })
}