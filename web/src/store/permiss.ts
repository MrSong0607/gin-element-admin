import { defineStore } from 'pinia';
import { Login, SessionInfo, Info, } from '../api/auth';
import { UserType } from '../dto/base'

declare type sessionStore = { info: SessionInfo | null, sessionKey: string | null }

export const usePermissStore = defineStore('permiss', {
    state(): sessionStore {
        const ss: sessionStore = { info: null, sessionKey: null }
        let str = localStorage.getItem("session")
        if (str) {
            ss.sessionKey = str
        }
        return ss
    },
    getters: {
        isLogined: (state): boolean => state.sessionKey != null,
        userType: (state): UserType | null => state.info?.user_type || null,
        isAdmin: (state): boolean => state.info?.user_type == UserType.admin
    },
    actions: {
        async login(param: { acct: string, passwd: string, code?: string }): Promise<SessionInfo> {
            const { data } = await Login(param)
            localStorage.setItem('session', data)
            this.sessionKey = data

            return await this.getAccountInfo()
        },
        clear() {
            this.info = null
            this.sessionKey = null
            localStorage.removeItem('session')
        },
        async getAccountInfo(): Promise<SessionInfo> {
            if (this.info) return this.info

            const { data } = await Info()
            this.info = data
            return data
        },
        hasPermission(...userType: UserType[]): boolean {
            if (!this.isLogined || !this.userType) {
                return false
            }

            if (userType.length == 0) {
                return true
            }

            if (this.isAdmin) return true

            return userType.includes(this.userType)
        }
    }
})