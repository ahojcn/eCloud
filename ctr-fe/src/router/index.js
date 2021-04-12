import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '@/view/Login'
import Register from '@/view/Register'
import Index from '@/view/Index'

Vue.use(VueRouter);

const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
};

export default new VueRouter({
    routes: [
        {
            path: "/login",
            component: Login
        },
        {
            path: "/register",
            component: Register
        },
        {
            path: "/",
            component: Index
        }
    ]
})
