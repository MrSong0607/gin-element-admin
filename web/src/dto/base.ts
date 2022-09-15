export enum UserType {
    customer = "customer",
    admin = "admin",
}

export function UserTypeName(u: UserType): string {
    switch (u) {
        case UserType.admin:
            return '管理员'
        case UserType.customer:
            return '客户'
        default:
            return '游客'
    }
}

export declare type ApiResponse<T> = {
    code: number;
    msg: string;
    data: T;
}

export declare type ArrayResonse<T> = {
    list: T[],
    count: number,
}

export declare type QueryBase = {
    size: number,
    page: number
}

export declare type QueryTimeBase = {
    start?: number,
    end?: number
}