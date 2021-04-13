import axios from 'axios'
import qs from 'qs'
import {LoadingBar, Notice} from 'view-design'

// axios 实例
const service = axios.create({
    baseURL: 'http://127.0.0.1:10001/',
    withCredentials: true, // send cookies when cross-domain requests
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    },
});

// request 拦截器
service.interceptors.request.use(
    config => {
        LoadingBar.start();
        // 默认 get 请求 query array 是 col[]=aaa&col[]=bbb
        // 引入 qs 插件转为 col=aaa&col=bbb
        if(config.method === 'get') {
            config.paramsSerializer = function (params) {
                return qs.stringify(params, {arrayFormat: 'repeat'})
            }
        }
        return config;
    },
    err => {
        LoadingBar.error();
        return err;
    }
);

// response 拦截器
service.interceptors.response.use(
    res => {
        LoadingBar.finish();

        Notice.success({
            title: res.data.msg
        })
        return res.data
    },
    err => {
        LoadingBar.error();

        let data = err.response.data
        Notice.error({
            title: data.msg,
            desc: JSON.stringify(data.data)
        })
        return data;
    }
);

export default service;