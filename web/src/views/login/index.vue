<script lang="ts" setup>
import { FormInstance, FormRules } from 'element-plus';
import { usePermissStore } from '../../store/permiss';
import { ElMessage } from 'element-plus'

const formRef = ref<FormInstance>()
const loading = ref(false)
const form = ref<{ acct: string, passwd: string, code?: string }>({ acct: '', passwd: '' })
const rules = reactive<FormRules>({
    acct: [{ required: true, message: '账号是必填项' }],
    passwd: [{ required: true, message: '密码是必填项' }],
    code: [{ type: 'number', message: '验证码必须是一个数字' }],
})
const router = useRouter()

const onSubmit = async (el: FormInstance | undefined) => {
    if (!el) return

    try {
        loading.value = true
        await el.validate()
        await usePermissStore().login(form.value)
        ElMessage.success('登录成功')

        router.push({ path: '/' })
    } catch (error) {
        console.log(error)
        return
    } finally {
        loading.value = false
    }
}

</script>
<template>
    <div class="login">
        <div class="box-card">
            <h1 class="title">管理后台</h1>
            <el-form :model="form" :rules="rules" size="large" ref="formRef" label-width="0"
                @keyup.enter.native="onSubmit(formRef)">
                <el-form-item prop="acct">
                    <el-input v-model="form.acct" placeholder="账号">
                        <template #prefix>
                            <el-icon>
                                <User />
                            </el-icon>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item prop="passwd">
                    <el-input v-model="form.passwd" type="password" placeholder="密码">
                        <template #prefix>
                            <el-icon>
                                <Lock />
                            </el-icon>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item prop="code">
                    <el-input v-model="form.code" placeholder="验证码">
                        <template #prefix>
                            <el-icon>
                                <Key />
                            </el-icon>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" :loading="loading" style="width: 100%;" @click="onSubmit(formRef)">登录
                    </el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<style scoped lang="scss">
.login {
    height: 100%;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;

    .box-card {
        margin-top: -15rem;
        width: 28rem;
        height: 20rem;
    }

    .title {
        text-align: center;
        font-size: 3rem;
    }
}
</style>