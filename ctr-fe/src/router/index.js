import Vue from 'vue';
import VueRouter from 'vue-router';

import Login from '@/view/Login';
import Register from '@/view/Register';
import Index from '@/view/Index';

import Service from '@/view/Service';
import TreeNodeCreate from "@/view/Service/TreeNodeCreate";
import TreeNodeDetail from "@/view/Service/TreeNodeDetail";

import Resource from '@/view/Resource';
import HostDetail from "@/view/Resource/HostDetail";
import HostAdd from "@/view/Resource/HostAdd";
import HostMonitor from "@/view/Resource/HostMonitor";
import HostRunCmd from "@/view/Resource/HostRunCmd";

import ICode from "@/view/ICode";
import ICodeDetail from "@/view/ICode/ICodeDetail";
import ICodeCreate from "@/view/ICode/ICodeCreate";

import Deploy from "@/view/Deploy";
import RouterInfo from '@/view/Deploy/RouterInfo'

Vue.use(VueRouter);

const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
};

let router = new VueRouter({
    routes: [
        {
            name: 'Login',
            path: "/login",
            component: Login,
            meta: {title: '登录'},
        },
        {
            name: 'Register',
            path: "/register",
            component: Register,
            meta: {title: '注册'},
        },
        {
            path: "/",
            name: "Index",
            redirect: "/service",
            component: Index,
            children: [
                {
                    name: 'Service',
                    path: '/service',
                    component: Service,
                    meta: {title: '服务'},
                    children: [
                        {
                            name: 'TreeNodeCreate',
                            path: '/service/create',
                            component: TreeNodeCreate,
                            meta: {title: '创建节点'}
                        },
                        {
                            name: 'TreeNodeDetail',
                            path: '/service/detail',
                            component: TreeNodeDetail,
                            meta: {title: '节点详情'}
                        },
                    ]
                },
                {
                    name: 'Deploy',
                    path: '/deploy',
                    component: Deploy,
                    meta: {title: '部署'},
                    children: [
                        {name: 'RouterInfo', path: 'deploy/router', component: RouterInfo, meta: {title: '接入层信息'}}
                    ]
                },
                {
                    name: 'Resource',
                    path: '/resource',
                    component: Resource,
                    meta: {title: '资源'},
                    children: [
                        {
                            name: 'HostDetail',
                            path: '/resource/host/detail',
                            component: HostDetail,
                            meta: {title: '主机详情'}
                        },
                        {
                            name: 'HostAdd',
                            path: '/resource/host/add',
                            component: HostAdd,
                            meta: {title: '添加主机'}
                        },
                        {
                            name: 'HostMonitor',
                            path: '/resource/host/monitor',
                            component: HostMonitor,
                            meta: {title: '主机监控'}
                        },
                        {
                            name: 'HostRunCmd',
                            path: '/resource/host/cmd',
                            component: HostRunCmd,
                            meta: {title: '主机监控'}
                        },
                    ]
                },
                {
                    name: "ICode",
                    path: "/icode",
                    component: ICode,
                    meta: {title: '注册'},
                    children: [
                        {
                            name: "ICodeCreate",
                            path: "/icode/create",
                            component: ICodeCreate,
                            meta: {title: "创建开发机"}
                        },
                        {
                            name: "ICodeDetail",
                            path: "/icode/detail",
                            component: ICodeDetail,
                            meta: {title: "详细信息"}
                        },
                    ]
                }
            ]
        },
    ]
})

router.beforeEach((to, from, next) => {
    window.document.title = to.meta.title + ' | 一站式云平台'
    next()
})

export default router;