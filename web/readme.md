首次使用需要先安装依赖:
```
yarn
```

启动本地调试
```
yarn run dev
```

构建
```
yarn run build
```
构建地址默认是`/dist/web`，可以在`vite.config.ts`中修改

项目结构
```
├─api           api请求封装
├─assets        资源文件
├─components    通用组件
│  └─Layout     布局组件
├─dto           api请求参数与响应类型
├─router        路由配置
├─store         状态管理
├─utils         工具库
└─views         视图
    ├─home      首页
    ├─login     登录页
    ├─profile   个人页
    │  └─components
    └─user
        └─components
```