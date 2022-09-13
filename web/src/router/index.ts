import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import layout from '../components/Layout/index.vue';
import { UserType } from '../api/base'

const routes: RouteRecordRaw[] = [
    {
        path: '/login',
        component: () => import('../views/login/index.vue'),
        meta: { hidden: true }
    },
    {
        path: '/',
        component: layout,
        children: [
            {
                name: '主页',
                path: '/',
                meta: {
                },
                component: () => import('../views/home/index.vue')
            }
        ]
    },
    {
        path: '/user',
        component: layout,
        name: '用户管理',
        children: [
            {
                name: '用户列表',
                path: 'list',
                meta: { icon: 'User', role: [UserType.admin] },
                component: () => import('../views/user/index.vue')
            }
        ]
    },
    {
        path: '/profile',
        component: layout,
        meta: { hidden: true },
        children: [
            {
                name: "个人信息",
                path: "index",
                component: () => import('../views/profile/index.vue')
            }
        ]
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
});

export { routes };
export default router