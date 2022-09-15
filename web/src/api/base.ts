import axios from 'axios';
import { ElMessage } from 'element-plus';
import 'element-plus/theme-chalk/src/message.scss';
import { usePermissStore } from '../store/permiss';

const service = axios.create({
    baseURL: import.meta.env.VITE_APP_BASE_API, // url = base url + request url
    // withCredentials: true, // send cookies when cross-domain requests
    timeout: 30000 // request timeout
})

service.interceptors.response.use(response => {
    const { code, msg } = response.data
    const session = usePermissStore()

    if (code !== 2000) {
        ElMessage.error(msg)

        if (code == 4001) {
            session.clear()
            location.reload()
        }
        return Promise.reject(new Error(msg))
    }
    return response
}, error => {
    ElMessage.error(error)
    return Promise.reject(error)
})

service.interceptors.request.use(config => {
    let session = usePermissStore()
    if (config.headers && session.isLogined) {
        config.headers["X-SESSION"] = session.sessionKey || ''
    }

    return config
})

export default service