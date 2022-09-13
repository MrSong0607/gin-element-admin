<script lang="ts" setup>
import UserInfo from './components/userinfo.vue'
import Security from './components/Security.vue';

</script>
<template>
    <el-tabs type="border-card" class="demo-tabs">
        <el-tab-pane>
            <template #label>
                <span class="custom-tabs-label">
                    <el-icon>
                        <calendar />
                    </el-icon>
                    <span>个人信息</span>
                </span>
            </template>
            <UserInfo />
        </el-tab-pane>
        <el-tab-pane label="账户安全">
            <Security />
        </el-tab-pane>
    </el-tabs>
</template>