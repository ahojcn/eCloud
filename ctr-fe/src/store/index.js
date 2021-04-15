import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        user_info: {},
        menu_active_name: 'service',
    },
    mutations: {
        setUserInfo(state, user_info) {
            state.user_info = user_info
        },
        setMenuActiveName(state, name) {
            state.menu_active_name = name
        },
    }
})