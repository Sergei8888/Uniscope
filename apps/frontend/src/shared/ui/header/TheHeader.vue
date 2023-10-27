<script setup lang="ts">
import LogoWithName from '@/shared/ui/header/LogoWithName.vue';
import NavBar from '@/shared/ui/header/NavBar.vue';
import SignupLink from '@/auth/ui/SignupLink.vue';
import UserBadge from '@/user/ui/UserBadge.vue';
import { useUserStore } from '@/user/model/user.store';

const userStore = useUserStore();
</script>

<template>
    <ElHeader class="header">
        <div class="header__content container">
            <router-link class="brand-link" :to="{ name: 'home' }">
                <LogoWithName class="header__logo" />
            </router-link>
            <NavBar />
            <UserBadge
                v-if="userStore.user"
                :nickname="userStore.user.nickname"
                class="header__auth-part"
            />
            <SignupLink v-else class="header__auth-part" />
        </div>
    </ElHeader>
</template>

<style scoped lang="scss">
.header {
    height: auto;
    background-color: var(--el-fill-color-dark);

    &__content {
        height: 100%;
        display: flex;
        align-items: center;

        padding-top: 10px;
        padding-bottom: 10px;
    }

    &__auth-part {
        margin-left: 40px;
    }
}

.brand-link {
    text-decoration: none;
}
</style>
