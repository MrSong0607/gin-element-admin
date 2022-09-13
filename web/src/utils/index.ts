import qrcode from 'qrcode'

export function timeFormat(unix: number): string {
    let dt = new Date(unix * 1000)
    return dt.toLocaleDateString() + ' ' + dt.toLocaleTimeString()
}

export function build_path(...args: string[]): string {
    return args.map((part, i) => {
        if (i === 0) {
            return part.trim()
        } else {
            return part.trim().replace(/(^[\/]*|[\/]*$)/g, '')
        }
    }).filter(x => x.length).join('/')
}

export async function generateQR(txt: string): Promise<string> {
    return await qrcode.toDataURL(txt, {
        width: 400
    })
}
