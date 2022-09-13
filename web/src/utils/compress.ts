interface BlobCallback {
    (blob: Blob | null): void;
}

export function compressFromDataUrl(dataUrl: string, callback: BlobCallback): void {
    let img = new Image();
    img.src = dataUrl
    img.onload = () => {
        let canvas = document.createElement('canvas')
        let ctx = canvas.getContext('2d');

        let width = img.width
        let height = img.height
        //如果图片大于四百万像素，计算压缩比并将大小压至400万以下
        let ratio
        if ((ratio = (width * height) / 4000000) > 1) {
            // console.log("大于400万像素");
            ratio = Math.sqrt(ratio)
            width /= ratio
            height /= ratio
        } else {
            ratio = 1
        }
        canvas.width = width
        canvas.height = height
        console.log(width, height)

        ctx?.drawImage(img, 0, 0, width, height)
        canvas.toBlob(callback, "image/jpeg", 0.5)
    }
}