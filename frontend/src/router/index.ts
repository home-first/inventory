import { createRouter, createWebHistory } from 'vue-router'

export const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/label/:label',
            name: 'label',
            component: () => import('../views/Label.vue')
        },
        {
            path: '/',
            name: 'home',
            component: () => import('../views/Home.vue')
        }
    ]
})