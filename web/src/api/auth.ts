import { ApiResponse, UserType, default as request } from './base'

export declare type SessionInfo = {
    id: number,
    acct: string,
    account_status: string,
    alias: string,
    user_type: UserType
}

export async function Login(data: any): Promise<ApiResponse<string>> {
    let resp = await request({
        url: '/auth/login',
        method: 'post',
        data
    })

    return resp.data
}

export async function Info(): Promise<ApiResponse<SessionInfo>> {
    let resp = await request({
        url: '/auth/info',
        method: 'get',
    })

    return resp.data
}

export declare type UserInfo = {
    alias: string,
}

export async function UpdateInfo(data: UserInfo): Promise<ApiResponse<any>> {
    let resp = await request({
        url: '/auth/info',
        method: 'post',
        data
    })

    return resp.data
}

export declare type SecurityInfo = {
    passwd: string,
    old_passwd: string
}

export async function Security(data: SecurityInfo): Promise<ApiResponse<any>> {
    let resp = await request({
        url: '/auth/security',
        method: 'post',
        data
    })

    return resp.data
}