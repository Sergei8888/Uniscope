<script lang="ts" setup>
import SigninForm from '@/auth/ui/SigninForm.vue';
</script>

<template>
    <div class="page-wrapper">
        <SigninForm />
    </div>
</template>

<style lang="scss" scoped>
.page-wrapper {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
}
</style>
