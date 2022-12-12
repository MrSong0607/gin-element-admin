<script setup lang="ts">
import sidebar from './Sidebar.vue';
import Header from './Header.vue'

const collapse = ref(false)
</script>
<template>
    <el-container style="height: 100%">
        <el-aside width="auto" class="aside">
            <sidebar :collapse="collapse" />
        </el-aside>
        <el-container>
            <el-header height="50px" class="header">
                <Header :collapse="collapse" @on-collapse="c => collapse = c" />
            </el-header>
            <el-main>
                <div class="content-box" :class="{ 'content-collapse': '' }">
                    <div class="content">
                        <router-view v-slot="{ Component, route }">
                            <transition name="slide-fade" mode="out-in">
                                <div :key="route.path">
                                    <keep-alive>
                                        <template v-if="route.meta.nocard">
                                            <component :is="Component" :key="route.path"></component>
                                        </template>
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
.aside {
    background: fixed;
    border-right: 1px solid var(--el-border-color)
}

.header {
    display: flex;
    align-items: center;
    padding: 0;
    box-shadow: 0 1px 4px rgb(0 21 41 / 8%)
}

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
</style>