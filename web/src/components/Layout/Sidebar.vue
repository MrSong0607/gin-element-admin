<script setup lang="ts">
import { routes } from '../../router/index';
import { usePermissStore } from '../../store/permiss';
import { build_path } from '../../utils/index';
import { UserType } from '../../dto/base'

const permiss = usePermissStore()
const props = defineProps<{ collapse: boolean }>()
interface MenuItem {
    name?: string | symbol,
    children?: MenuSubItem[],
    path: string
    icon?: string
}

interface MenuSubItem {
    name?: string | symbol,
    path: string;
    icon?: string
}

const menus: MenuItem[] = []

routes.forEach(t => {
    if (!(t.meta?.hidden || false) && permiss.hasPermission(...((t.meta?.role || []) as UserType[]))) {
        let m: MenuItem = Object.assign({ icon: '' }, t)
        m.icon = t.meta?.icon ? String(t.meta?.icon) : undefined

        m.children = t.children?.filter(r => {
            return !(r.meta?.hidden || false) && permiss.hasPermission(...((r.meta?.role || []) as UserType[]))
        }).map(x => {
            let c: MenuSubItem = Object.assign({}, x)
            c.icon = x.meta?.icon ? String(x.meta?.icon) : undefined
            return c
        })

        if ((t.children?.length || 0) > 0 && m.children?.length == 0) {
            return
        }

        menus.push(m)
    }
})

const warpRoute = (item: MenuItem) => {
    if (item.children) {
        if (item.children.length == 1) {
            return build_path(item.path, item.children[0].path)
        }
    } else {
        return item.path
    }
    return ''
}

const activeIndex = ref(useRoute().fullPath)
// console.log(menus)
</script>
<template>
    <el-menu class="menu" :router="true" :default-active="activeIndex" :val="props.collapse" :collapse="props.collapse"
        mode="vertical" background-color="#324157" text-color="#bfcbd9" active-text-color="#20a0ff">
        <el-menu-item index="0" class="logo">
            <el-icon v-if="props.collapse" :size="50">
                <ElementPlus />
            </el-icon>
            <div v-else>演示系统</div>
        </el-menu-item>
        <template v-for="(val) in menus">
            <el-sub-menu :index="`${val.path}`" v-if="(val.children?.length || 0) > 1">
                <template #title>
                    <el-icon v-if="val.icon">
                        <component :is="val.icon"></component>
                    </el-icon>
                    <span>{{ val.name }}</span>
                </template>
                <el-menu-item v-for="(sub) in val.children" :path="`${val.path}/${sub.path}`"
                    :index="build_path(val.path, sub.path)">{{ sub.name }}
                </el-menu-item>
            </el-sub-menu>
            <el-menu-item :index="warpRoute(val)" v-else-if="(val.children?.length || 0) == 1">
                <el-icon v-if="val.icon">
                    <component :is="val.icon"></component>
                </el-icon>
                <template #title> {{ val.children && val.children[0].name }}</template>
            </el-menu-item>
            <el-menu-item :index="`${val.path}`" v-else>
                <template #title>
                    <el-icon v-if="val.icon">
                        <component :is="val.icon"></component>
                    </el-icon>
                    <span>{{ val.name }}</span>
                </template>
            </el-menu-item>
        </template>
    </el-menu>
</template>
<style lang="scss" >
.el-menu {
    border-right: 0;
}

.el-sub-menu .el-menu-item {
    background-color: #1f2d3d;
}

.menu:not(.el-menu--collapse) {
    width: 200px;
}

.logo {
    padding: 0;
    justify-content: center;
    background-color: #1f2d3d;
    height: 50px;

    .div {
        width: 100%;
    }
}
</style>