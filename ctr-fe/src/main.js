import Vue from 'vue'
import App from './App.vue'

import animated from 'animate.css'
import ViewUI from 'view-design';
import 'view-design/dist/styles/iview.css';

import router from './router'
import store from './store'

Vue.use(animated);
Vue.use(ViewUI);

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  router,
  store,
}).$mount('#app')
