<script lang="ts" setup>
import 'element-plus/theme-chalk/dark/css-vars.css';
import 'element-plus/es/components/notification/style/css'; // this is only needed if the page also used ElMessage

import { useAuthStore } from '@/auth/model/auth.store';

const authStore = useAuthStore();
authStore.authByTokens();
</script>

<template>
    <router-view v-if="authStore.primaryAuthFinished" />
    <div v-else>Loading...</div>
</template>

<style lang="scss">
@import 'normalize.css';
@import 'scss/css_variables.scss';
@import 'scss/global.scss';
</style>
