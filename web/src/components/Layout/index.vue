<script setup lang="ts">
import { usePermissStore } from '../../store/permiss';
import sidebar from './Sidebar.vue';

const route = useRoute()
const collapse = ref(false)
const store = usePermissStore()

const handleCollapse = () => {
    collapse.value = !collapse.value
}

const paths = computed(() => {
    return route.matched
})

const logout = () => {
    store.clear()
    location.reload()
}
// console.log(route)
</script>
<template>
    <el-container style="height: 100%">
        <transition name="collapse" mode="out-in">
            <el-aside v-if="collapse" width="64px" class="aside">
                <div style="height:60px" class="logo">Unicom</div>
                <sidebar :collapse="collapse" />
            </el-aside>
            <el-aside v-else width="200px" class="aside">
                <div style="height:60px" class="logo">Unicom</div>
                <sidebar :collapse="collapse" />
            </el-aside>
        </transition>

        <el-container>
            <el-header class="header">
                <div class="collapse" @click="handleCollapse">
                    <el-icon>
                        <Fold />
                    </el-icon>
                </div>
                <el-breadcrumb separator="/">
                    <el-breadcrumb-item v-for="val in paths">{{ val.name }}</el-breadcrumb-item>
                </el-breadcrumb>
                <div class="right">
                    <el-dropdown>
                        <span class="dropdown-link">
                            {{ store.info?.alias }}
                            <el-icon class="el-icon--right">
                                <arrow-down />
                            </el-icon>
                        </span>
                        <template #dropdown>
                            <el-dropdown-menu>
                                <el-dropdown-item>
                                    <router-link to="/profile/index">个人信息</router-link>
                                </el-dropdown-item>
                                <el-dropdown-item divided>
                                    <el-button size="small" type="danger" @click="logout" text>退出</el-button>
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>
                </div>
            </el-header>
            <el-main>
                <div class="content-box" :class="{ 'content-collapse': '' }">
                    <div class="content">
                        <router-view v-slot="{ Component, route }">
                            <transition name="slide-fade" mode="out-in">
                                <div :key="route.path">
                                    <keep-alive>
                                        <div v-if="route.meta.nocard">
                                            <component :is="Component" :key="route.path"></component>
                                        </div>
                                        <el-card v-else>
                                            <component :is="Component" :key="route.path"></component>
                                        </el-card>
                                    </keep-alive>
                                </div>
                            </transition>
                        </router-view>
                    </div>
                </div>
            </el-main>
        </el-container>
    </el-container>
</template>
<style lang="scss" scoped>
.slide-fade-enter-active {
    transition: all 0.3s ease-out;
}

.slide-fade-leave-active {
    transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
    transform: translateX(20px);
    opacity: 0;
}

.collapse-enter-active,
.collapse-leave-active {
    transition: all 0.3s linear;
    transform: scaleX(0);
}

.collapse-leave-to,
.collapse-enter {
    transform: scaleX(100%);
}

.aside {
    background: fixed;
    background-color: #324157;
}

.header {
    background-color: #ffffff;
    display: flex;
    align-items: center;
    padding: 0;
    box-shadow: 0 1px 4px rgb(0 21 41 / 8%)
}

.logo {
    background-color: #2b2f3a;
    color: #ffffff;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: bold;
    will-change: filter;
    height: 60px;
}

.logo:hover {
    filter: drop-shadow(0 0 2em #e5e5e9aa);
}

.collapse {
    font-size: 1.3rem;
    padding: 17px 20px;
    cursor: pointer;

    :hover {
        background: rgba(0, 0, 0, 0.025);
    }
}

.right {
    margin-left: auto;
    order: 2;
    padding: 20px;
    font-weight: bold;
}
</style>