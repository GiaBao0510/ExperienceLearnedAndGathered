import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import { router } from './router/index'

/*
    Vue-router là một plugin:
    - Sử dụng vue-router như là 1 ổ căm, truy cập bên trong bất kỳ thành phần nào bằng [this.$router].
    - route gần có thể được truy cập bằng [this.$route]
*/

createApp(App)
    .use(router)
    .mount('#app');
