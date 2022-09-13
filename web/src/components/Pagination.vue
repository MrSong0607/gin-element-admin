<script setup lang="ts">
const props = defineProps<{ background?: boolean, small?: boolean, pageSizes?: number[], pageSize: number, currentPage: number, total: number }>()
const emit = defineEmits<{
    (e: 'fetch', page: number, size: number): void,
    (e: 'update:currentPage', page: number): void,
    (e: 'update:pageSize', page: number): void,
}>()

const currentPageVal = computed({
    get() {
        return props.currentPage
    },
    set(val) {
        emit('update:currentPage', val)
        emit('fetch', val, pageSizeVal.value)
    }
})

const pageSizeVal = computed({
    get() {
        return props.pageSize
    },
    set(val) {
        emit('update:pageSize', val)
        emit('fetch', currentPageVal.value, val)
    }
})
</script>
<template>
    <el-pagination v-model:current-page="currentPageVal" @update:current-page="(val:number) => currentPageVal = val"
        v-model:page-size="pageSizeVal" @update:page-size="(val:number) => pageSizeVal = val"
        :page-sizes="pageSizes || [10, 20, 50, 100]" :small="small" :background="background"
        layout="total, prev, pager, next, sizes" :total="total" class="pagination">
    </el-pagination>
</template>
<style scoped>
.pagination {
    margin: 10px 0 0 0;
}

:deep(.el-pagination__sizes) {
    margin: 0 0 0 10px;
}
</style>