import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '@/view/Login'
import Register from '@/view/Register'
import Index from '@/view/Index'
import IndexService from '@/view/Index/Service'
import IndexResource from '@/view/Index/Resource'

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
                    component: IndexService,
                    meta: {title: '服务'},
                },
                {
                    name: 'Resource',
                    path: '/resource',
                    component: IndexResource,
                    meta: {title: '资源'},
                }
            ]
        }
    ]
})

router.beforeEach((to, from, next) => {
    window.document.title = to.meta.title + ' | 一站式云平台';
    next()
})

export default router;