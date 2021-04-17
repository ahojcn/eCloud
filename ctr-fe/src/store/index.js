import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        user_info: {},
        menu_active_name: 'service',
        tree_node_type: {
            "4": {title: "Namespace", color: "#f56a00"},
            "3": {title: "ProductLine", color: "#7265e6"},
            "2": {title: "Group", color: "#ffbf00"},
            "1": {title: "Service", color: "#87d068"},
            "0": {title: "Cluster", color: "#00a2ae"}
        },
        user_tree_rights: [
            {value: 0, label: "n 无权限"},
            {value: 1, label: "r 只读"},
            {value: 2, label: "w 只写"},
            {value: 3, label: "rw 可读写"},
            {value: 4, label: "c 可新增"},
            {value: 5, label: "d 可删除"},
            {value: 6, label: "a 管理"},
        ],
        tree_node_type_selector: [
            {value: 0, label: 'cluster'},
            {value: 1, label: 'service'},
            {value: 2, label: 'group'},
            {value: 3, label: 'pdl'},
            {value: 4, label: 'namespace'},
        ],
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