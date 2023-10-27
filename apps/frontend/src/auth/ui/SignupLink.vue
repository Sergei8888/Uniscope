<script setup lang="ts">
withDefaults(
    defineProps<{
        size?: 'small' | 'big';
    }>(),
    {
        size: 'small',
    }
);
</script>

<template>
    <router-link
        :to="{ name: 'signup' }"
        class="signup-link"
        :class="{
            'signup-link--big': size === 'big',
            'signup-link--small': size === 'small',
        }"
    >
        Присоединиться
    </router-link>
</template>

<style scoped lang="scss">
.signup-link {
    display: block;
    background-color: var(--el-color-primary);
    border-radius: var(--el-border-radius-base);
    font-size: var(--el-font-size-base);
    font-weight: 500;
    text-transform: uppercase;
    text-decoration: none;

    transition: 0.2s background-color;

    &:hover {
        background-color: var(--el-color-primary-light-3);
    }

    &--small {
        font-size: 14px;
        padding: 8px 15px;
    }

    &--big {
        font-size: 18px;
        padding: 12px 18px;
    }
}
</style>
