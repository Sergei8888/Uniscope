<script setup lang="ts">
defineProps<{
    nickname: string;
}>();
</script>

<template>
    <div class="user-badge">
        <div class="user-badge__nickname">{{ nickname }}</div>
    </div>
</template>

<style scoped lang="scss">
.user-badge {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 10px 15px;
    border: 1px solid var(--el-border-color);
    border-radius: 7px;
    background-color: var(--el-fill-color-light);

    &__avatar {
        width: 35px;
        height: 35px;
    }
}
</style>
