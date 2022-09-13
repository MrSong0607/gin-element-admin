import { ApiResponse, ArrayResonse, default as request } from './base'

export async function ListUser(params: any): Promise<ApiResponse<ArrayResonse<any>>> {
    let resp = await request({
        url: '/user/list',
        method: 'get',
        params
    })

    return resp.data
}

export async function CreateUser(data: { acct: string, pwd: string, alias: string }): Promise<ApiResponse<any>> {
    let resp = await request({
        url: '/user/new',
        method: 'post',
        data
    })

    return resp.data
}