import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        user_info: {},
    },
    mutations: {
        set_user_info(state, user_info) {
            state.user_info = user_info
        }
    }
})