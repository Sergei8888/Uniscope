<script lang="ts" setup>
import SignupForm from '@/auth/ui/SignupForm.vue';
</script>

<template>
    <div class="page-wrapper">
        <SignupForm />
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
