import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '@/view/Login'
import Register from '@/view/Register'
import Index from '@/view/Index'
import Service from '@/view/Index/Service'
import Resource from '@/view/Index/Resource'
import ICode from "@/view/ICode";

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
                },
                {
                    name: 'Resource',
                    path: '/resource',
                    component: Resource,
                    meta: {title: '资源'},
                },
                {
                    name: "ICode",
                    path: "/icode",
                    component: ICode,
                    meta: {title: '注册'},
                }
            ]
        },
    ]
})

router.beforeEach((to, from, next) => {
    window.document.title = to.meta.title + ' | 一站式云平台';
    next()
})

export default router;