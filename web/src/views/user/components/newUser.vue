<script lang="ts" setup>
import { FormInstance, FormRules } from 'element-plus';
import { CreateUser } from '../../../api/user';
import {ElMessage} from 'element-plus'

type CreateUserForm = {
    acct: string, alias: string, pwd: string
}

const emit = defineEmits<{
    (e: 'onSubmit', param: CreateUserForm): void,
    (e: 'onCancel'): void
}>()

const formRef = ref<FormInstance>()
const form = ref({ acct: '', alias: '', pwd: '' })
const loading = ref(false)

const rules = reactive<FormRules>({
    acct: [{ required: true, message: '账号是必填项' }],
    alias: [{ required: true, message: '名称是必填项' }],
    pwd: [{ required: true, message: '密码是必填项' }],
})

const onSubmit = async (el: FormInstance | undefined) => {
    if (!el) return

    try {
        loading.value = true
        await el.validate()
        await CreateUser(form.value)
        ElMessage.success('创建用户成功')
        emit('onSubmit', form.value)
    } catch (error) {
        console.log(error)
        return
    } finally {
        loading.value = false
    }
}
</script>
<template>
    <el-form :model="form" :rules="rules" ref="formRef" label-width="5rem">
        <el-form-item label="账号" prop="acct">
            <el-input v-model="form.acct" placeholder="账号" />
        </el-form-item>
        <el-form-item label="名称" prop="alias">
            <el-input v-model="form.alias" placeholder="名称" />
        </el-form-item>
        <el-form-item label="密码" prop="pwd">
            <el-input v-model="form.pwd" type="password" placeholder="密码" show-password />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" :loading="loading" @click="onSubmit(formRef)">提交</el-button>
            <el-button @click="emit('onCancel')">取消</el-button>
        </el-form-item>
    </el-form>
</template>