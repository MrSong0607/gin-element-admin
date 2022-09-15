<script lang="ts" setup>
import { usePermissStore } from '../../store/permiss';
const store = usePermissStore()
const props = defineProps<{ collapse: boolean }>()
const emit = defineEmits<{ (e: 'onCollapse', collapse: boolean): void }>()
const route = useRoute()

const paths = computed(() => {
    return route.matched
})

const handleCollapse = () => {
    emit('onCollapse', !props.collapse)
}

const logout = () => {
    store.clear()
    location.reload()
}
</script>
<template>
    <div class="collapse" @click="handleCollapse">
        <el-icon>
            <Expand v-if="props.collapse" />
            <Fold v-else />
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
</template>

<style lang="scss" scoped>
.header {
    background-color: #ffffff;
    display: flex;
    align-items: center;
    padding: 0;
    box-shadow: 0 1px 4px rgb(0 21 41 / 8%)
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