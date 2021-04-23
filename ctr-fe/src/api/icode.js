import request from "@/util/requests";

// 创建icode
export function apiCreateICode(data) {
    return request({
        url: 'icode',
        method: 'post',
        data
    })
}

// 获取icdoe开发机列表
export function apiGetICodeList(params) {
    return request({
        url: 'icode',
        method: 'get',
        params: params
    })
}

// 删除icdoe开发机
export function apiDeleteICode(params) {
    return request({
        url: 'icode',
        method: 'delete',
        params: params
    })
}
