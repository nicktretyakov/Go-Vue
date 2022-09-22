import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import TimelineView from '../views/TimelineView.vue';
import SettingsView from '../views/SettingsView.vue';

const routes = [
    {
        path: '/',
        name: 'home',
        component: HomeView,
    },
    {
        path: '/timeline',
        name: 'timeline',
        component: TimelineView,
    },
    {
        path: '/settings',
        name: 'settings',
        component: SettingsView,
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

export default router;
