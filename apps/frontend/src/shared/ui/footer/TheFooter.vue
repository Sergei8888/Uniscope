<script setup lang="ts"></script>

<template>
    <ElFooter class="footer">
        <div class="container footer__content">
            <div class="footer__left-part">
                <span class="footer__organization">
                    2021 - 2023 Uniscope "ООО Ксенофиум"
                </span>
                <router-link class="el-link" :to="{ name: 'contacts' }">
                    Контакты
                </router-link>
            </div>
            <ElLink
                :underline="false"
                class="footer__email"
                href="mailto:uniscope.space@gmail.com"
            >
                uniscope.space@gmail.com
            </ElLink>
        </div>
    </ElFooter>
</template>

<style scoped lang="scss">
.footer {
    height: auto;

    padding-top: 12px;
    padding-bottom: 12px;

    background-color: var(--el-fill-color-dark);

    &__content {
        display: flex;
        justify-content: space-between;
    }

    &__organization {
        font-size: var(--el-font-size-base);
        font-weight: var(--el-font-weight-primary);
        color: var(--el-text-color-regular);
    }

    &__left-part {
        display: flex;
        gap: 20px;
    }

    &__email {
        transition: color 0.2s;
    }
}
</style>
