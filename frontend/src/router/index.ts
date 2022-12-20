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
        },
        {
            path: '/items',
            name: 'items',
            component: () => import('../views/Items.vue'),
            children: [
                {
                    path: 'new',
                    name: 'items.new',
                    component: () => import('../components/Items/New.vue')
                },
                {
                    path: ':id',
                    name: 'items.detail',
                    component: () => import('../components/Items/Detail.vue')
                },
                {
                    path: ':id/edit',
                    name: 'items.edit',
                    component: () => import('../components/Items/Edit.vue')
                },
                {
                    path: '',
                    name: 'items.list',
                    component: () => import('../components/Items/List.vue')
                }

            ]
        },
        {
            path: '/collections',
            name: 'collections',
            component: () => import('../views/Collections.vue'),
            children: [
                {
                    path: 'new',
                    name: 'collections.new',
                    component: () => import('../components/Collections/New.vue')
                },
                {
                    path: ':id',
                    name: 'collections.detail',
                    component: () => import('../components/Collections/Detail.vue')
                },
                {
                    path: ':id/edit',
                    name: 'collections.edit',
                    component: () => import('../components/Collections/Edit.vue')
                },
                {
                    path: '',
                    name: 'collections.list',
                    component: () => import('../components/Collections/List.vue')
                }
            ]
        },
        {
            name: '404',
            path: '/:pathMatch(.*)*',
            component: () => import('../views/404.vue')
        }
    ]
})