<script lang="ts" setup>
import { QueryBase } from '../../api/base';
import { ListUser } from '../../api/user';
import Pagination from '../../components/Pagination.vue';
import { timeFormat } from '../../utils/index';
import newUser from './components/newUser.vue';

const data = ref<any[]>([])
const query = ref<QueryBase & { acct?: string }>({ size: 20, page: 1, })
const total = ref(0)
const showCreate = ref(false)

const fetchData = async () => {
    const { data: res } = await ListUser(query.value)
    data.value = res.list
    total.value = res.count
}

onMounted(() => {
    fetchData()
})
</script>
<template>
    <el-form :inline="true" :model="query" size="small" style="display:flex">
        <el-form-item label="单号">
            <el-input v-model="query.acct" placeholder="账号" clearable />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="fetchData">查询</el-button>
            <el-button type="primary" @click="showCreate = true">创建</el-button>
        </el-form-item>
    </el-form>

    <el-table :data="data" size="small" stripe border>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="alias" label="名称" min-width="80" show-overflow-tooltip />
        <el-table-column prop="acct" label="账号" min-width="80" show-overflow-tooltip />
        <el-table-column prop="chn_code" label="推广码" min-width="150" show-overflow-tooltip />
        <el-table-column label="创建时间" width="130">
            <template #default="{ row }">
                <span>{{ timeFormat(row.create_at) }}</span>
            </template>
        </el-table-column>
        <el-table-column label="操作" min-width="100" />
    </el-table>
    <Pagination v-model:page-size="query.size" v-model:current-page="query.page" :total="total" @fetch="fetchData" />

    <el-dialog v-model="showCreate" title="创建用户" width="50%" destroy-on-close center>
        <newUser @onSubmit="showCreate = false" @onCancel="showCreate = false" />
    </el-dialog>
</template>