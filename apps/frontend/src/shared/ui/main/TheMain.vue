<script setup lang="ts"></script>

<template>
    <ElMain style="--el-main-padding: 0" class="main container">
        <router-view />
    </ElMain>
</template>

<style scoped lang="scss">
.main {
    padding-top: 60px;
    padding-bottom: 60px;
}
</style>
