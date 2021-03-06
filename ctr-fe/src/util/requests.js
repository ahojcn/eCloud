import axios from 'axios'
import qs from 'qs'
import {LoadingBar, Notice} from 'view-design'

Notice.config({
    top: 100,
    duration: 3
})

// axios 实例
const service = axios.create({
    baseURL: 'http://127.0.0.1:10001/',
    //baseURL: 'http://192.168.0.101:10001/',
    withCredentials: true, // send cookies when cross-domain requests
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    },
    // timeout: 60000 * 10,
});

// request 拦截器
service.interceptors.request.use(
    config => {
        LoadingBar.start();
        // 默认 get 请求 query array 是 col[]=aaa&col[]=bbb
        // 引入 qs 插件转为 col=aaa&col=bbb
        if (config.method === 'get') {
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
        if (err && err.response) {
            // 服务器返回的错误信息
            let data = err.response.data
            Notice.error({
                title: data.msg,
                desc: JSON.stringify(data.data),
            })
            return data;
        } else {
            Notice.error({
                title: '出错了！',
                desc: err,
            })
            return err;
        }
    }
);

export default service;
