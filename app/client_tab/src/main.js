import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './reset.scss';
import '@/assets/styles/fonts.scss';

createApp(App).use(store).use(router).mount('#app');
