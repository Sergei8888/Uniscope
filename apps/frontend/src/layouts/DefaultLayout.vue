<script lang="ts" setup>
import TheFooter from '@/shared/ui/footer/TheFooter.vue';
import TheHeader from '@/shared/ui/header/TheHeader.vue';
import TheMain from '@/shared/ui/main/TheMain.vue';
</script>

<template>
    <ElContainer direction="vertical" class="layout">
        <TheHeader />
        <TheMain />
        <TheFooter />
    </ElContainer>
</template>

<style lang="scss" scoped></style>
