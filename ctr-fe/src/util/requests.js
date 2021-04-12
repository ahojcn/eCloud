import axios from 'axios'
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
        return config;
    },
    err => {
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