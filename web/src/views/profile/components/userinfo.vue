<script lang="ts" setup>
import { FormRules, ElMessage, FormInstance } from 'element-plus'
import { UserInfo, UpdateInfo } from '@/api/auth'
import { usePermissStore } from '@/store/permiss'

const formRef = ref<FormInstance>()
const store = usePermissStore()
const form = ref<UserInfo>({ alias: store.info?.alias || '' })
const loading = ref(false)
const rules = reactive<FormRules>({
    alias: [{ required: true, message: '名称是必填项' }],
})

const onSubmit = async (el: FormInstance | undefined) => {
    if (!el) return

    try {
        loading.value = true
        await el.validate()
        await UpdateInfo(form.value)
        ElMessage.success('修改信息成功')
        if (store.info) {
            store.info.alias = form.value.alias
        }
    } catch (error) {
        console.log(error)
    } finally {
        loading.value = false
    }
}
</script>
<template>
    <el-form :model="form" :rules="rules" ref="formRef" label-width="5rem">
        <el-form-item label="名称" prop="alias">
            <el-input v-model="form.alias" placeholder="名称" />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" :loading="loading" @click="onSubmit(formRef)">提交</el-button>
        </el-form-item>
    </el-form>
</template>