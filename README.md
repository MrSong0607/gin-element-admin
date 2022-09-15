# gin-element-admin
本项目使用golang + vue + element-plus实现了一个后台管理系统，具备基础的登录，角色权限验证等功能，拓展方便，结构简单。可以根据需要添加功能。

后端技术栈
- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/go-gorm/gorm)
- [redigo](https://github.com/gomodule/redigo)

前端技术栈
- [vue3](https://github.com/vuejs/vue)
- [element-plus](https://github.com/element-plus/element-plus)
- [pinia](https://github.com/vuejs/pinia)

## [后端](api/readme.md)
### 注册路由的方法:   
`controllers/base.go`中定义了
```golang
type Controller interface {
	RouterRegist(r *gin.Engine)
}
```
使用时，只需要定义一个`struct`
```golang
type User struct{}
```
然后实现RouterRegist方法，并在init()中添加该类型的实例到_router中即可
```golang
func (f User) RouterRegist(r *gin.Engine) {
	g := r.Group("user")
	{
		g.POST("info", ApiWrap(f.Info))
	}
}

func init() {
	_router = append(_router, User{})
}
```

### 内置功能
api已内置基础的中间件，可以直接使用:

- ApiAuth 使用方法如下
```golang
g.GET("info", ApiAuth(), ApiWrap(f.Info))
//ApiAuth接受用户角色可选参数，表示该接口只允许特定角色访问,注:admin用户不受此功能限制
g.GET("info", ApiAuth(models.UserCustomer), ApiWrap(f.Info))
```

- QpsBlock 访问频率限制，使用方法
```golang
g.GET("info", QpsBlock("limit_openapi",100,60), ApiWrap(f.Info)) //表示60秒内最多访问100次
```

- GetSessionInfo方法可以在添加了ApiAuth中间件的接口方法中使用以获取当前登录用户的个人信息
```golang
func (Auth) UpdateInfo(c *gin.Context) (r any, e error) {
    session := GetSessionInfo(c)
}
```

- ApiWrap函数统一接口的返回值类型，该函数接受一个`func(*gin.Context) (any, error)`参数，并返回接口调用者一个json`dto.ApiResponse`,当返回值error不为`nil`时，error会被写到`dto.ApiResponse.Message`中

## [前端](web/readme.md)

### 路由
路由注册在`/src/router/index.ts`中，仿照例子添加即可   
`src/components/Layout/index.vue`是后台默认的布局，需要使用时，在路由配置中设置`component: layout`，然后将view添加到`children`中即可，最终的地址是`{path}/{children.path}`   
路由配置项的meta, 设置`hidden: true`的路由不会显示在左侧的菜单列表中，如果父路由的`hidden`为`true`，则整个children都不会显示   
设置meta.role可以配置不同角色的菜单项，role接受一个`UserType`类型的数组，`UserType`的定义在`/src/api/base.ts`中

### 状态管理
状态在`/src/store`中定义，`permiss.ts`中定义了用户信息和权限相关的状态   
如果需要在`views`中使用按钮和列表粒度的权限控制，可以使用如下方法
```vue
<script lang="ts" setup>
import { usePermissStore } from '../../store/permiss'
const permiss = usePermissStore().hasPermission
const data = ref([])
</script>

<template>
    <el-table :data="data" size="small" stripe border>
        <el-table-column label="操作" min-width="100">
            <template #default="{ row }">
                <el-button size="small" type="primary" v-if="permiss(UserType.admin)" plain>
                    编辑
                </el-button>
            </template>
        </el-table-column>
    </el-table>
</template>
```

## 注意事项
- 由于使用了vue router的WebHistory模式，所以nginx配置需要添加try_files命令，示例如下
```nginx
location / {
        root /home/gin-element-admin/web;
        try_files $uri $uri/ /index.html;
    }
```