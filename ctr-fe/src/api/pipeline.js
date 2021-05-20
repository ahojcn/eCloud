import request from '@/util/requests'

// 获取流水线列表
export function apiGetPipeLineList(params) {
    return request({
        url: 'pipeline/list',
        method: 'get',
        params: params
    })
}

// 获取流水线状态信息
export function apiGetPipeLineStatus(params) {
    return request({
        url: 'pipeline/status',
        method: 'get',
        params: params
    })
}

// 重置流水线
export function apiResetPipeLine(params) {
    return request({
        url: 'pipeline/reset',
        method: 'get',
        params: params
    })
}

// 删除流水线
export function apiDeletePipeLine(params) {
    return request({
        url: 'pipeline/delete',
        method: 'get',
        params: params
    })
}

// 创建流水线
export function apiCreatePipeLine(data) {
    return request({
        url: 'pipeline/create',
        method: 'post',
        data
    })
}

// 执行流水线
export function apiRunPipeLine(data) {
    return request({
        url: 'pipeline/run',
        method: 'post',
        data
    })
}