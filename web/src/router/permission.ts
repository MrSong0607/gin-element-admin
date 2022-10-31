import { usePermissStore } from '../store/permiss'
import router from './index'

const whiteList = ['/login']

router.beforeEach(async (to, from, next) => {
    const store = usePermissStore()

    // console.log(to.fullPath, store.isLogined)
    if (!store.isLogined) {
        if (whiteList.indexOf(to.path) !== -1) {
            next()
        } else {
            next(`/login?redirect=${to.path}`)
        }
        return
    }

   await store.getAccountInfo()

    if (to.path === '/login') {
        next({ path: '/' })
        return
    }

    next()
})
