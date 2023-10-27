<script setup lang="ts">
import Logo from '@/shared/ui/header/assets/logo.png';
</script>

<template>
    <div class="brand">
        <ElImage :src="Logo" class="brand__logo" />
        <span class="brand__name">Uniscope</span>
    </div>
</template>

<style scoped lang="scss">
.brand {
    display: flex;
    align-items: center;

    &__logo {
        width: 50px;
        margin-right: 10px;
    }

    &__name {
        font-size: 38px;
        text-transform: uppercase;
    }
}
</style>
