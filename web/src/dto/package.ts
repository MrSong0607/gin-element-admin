export declare type PackageInfo = {
    background: string
    create_at: number
    create_by: number
    enable: boolean
    goods_id: string
    id: number
    imgs: string
    month_fee: number
    pkg_name: string
}

export declare type PackageImgs = {
    header_img: string,
    after_num_img: string[],
    detail_img: string[],
    footer_img: string,
}

export declare type NewPkgParam = {
    pkg_name: string,
    goods_id: string,
    month_fee: number,
    background: string
    imgs: PackageImgs,
}

export declare type UpdatePkgParam = {
    id: number
} & NewPkgParam