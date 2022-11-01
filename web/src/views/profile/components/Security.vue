<script lang="ts" setup>
import { SecurityInfo, Security } from '@/api/auth'
import { FormRules, FormInstance, ElMessage } from 'element-plus'

const passwdValidate = (rule: any, value: any, callback: any) => {
    if (!/\w+/.test(value)) {
        callback(new Error('必须包含小写字母'))
        return
    }

    if (!/[A-Z]+/.test(value)) {
        callback(new Error('必须包含大写字母'))
        return
    }

    if (!/\d+/.test(value)) {
        callback(new Error('必须包含数字'))
        return
    }

    if ((value as string).length < 8) {
        callback(new Error('密码长度必须大于8'))
        return
    }

    callback()
}

const repeatPasswdValidate = (rule: any, value: any, callback: any) => {
    if (value != form.value.passwd) {
        callback(new Error('两次密码不一致'))
        return
    }
    callback()
}

const form = ref<SecurityInfo & { repeatPasswd: string }>({ passwd: '', old_passwd: '', repeatPasswd: '' })
const loading = ref(false)
const rules = reactive<FormRules>({
    old_passwd: [{ required: true, message: '旧密码是必填项' }],
    passwd: [{ required: true, message: '密码是必填项' }, { validator: passwdValidate, trigger: 'blur' }],
    repeatPasswd: [{ required: true, message: '重复密码是必填项' }, { validator: repeatPasswdValidate, trigger: 'blur' }],
})
const formRef = ref<FormInstance>()


const onSubmit = async (el: FormInstance | undefined) => {
    if (!el) return

    try {
        loading.value = true
        await el.validate()
        await Security(form.value)
        ElMessage.success('修改密码成功')
    } catch (error) {

    } finally {
        loading.value = false
    }
}
</script>
<template>
    <el-form :model="form" :rules="rules" ref="formRef" label-width="5rem">
        <el-form-item label="旧密码" prop="old_passwd">
            <el-input v-model="form.old_passwd" type="password" placeholder="旧密码" />
        </el-form-item>
        <el-form-item label="新密码" prop="passwd">
            <el-input v-model="form.passwd" type="password" placeholder="新密码" />
        </el-form-item>
        <el-form-item label="重复密码" prop="repeatPasswd">
            <el-input v-model="form.repeatPasswd" type="password" placeholder="重复新密码" />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" :loading="loading" @click="onSubmit(formRef)">提交</el-button>
        </el-form-item>
    </el-form>
</template>