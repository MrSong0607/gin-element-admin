/// <reference types="vite/client" />

interface ImportMetaEnv {
    readonly VITE_APP_PAGE: string,
    readonly VITE_APP_BASE_API: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}