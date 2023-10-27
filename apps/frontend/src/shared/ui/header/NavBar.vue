<script setup lang="ts"></script>

<template>
    <nav class="navigation">
        <ul class="navigation__list">
            <li class="navigation__item">
                <router-link
                    :to="{ name: 'telescope-catalog' }"
                    class="navigation__link"
                >
                    Каталог телескопов
                </router-link>
            </li>
            <li class="navigation__item">
                <router-link :to="{ name: 'gallery' }" class="navigation__link">
                    Галерея
                </router-link>
            </li>
        </ul>
    </nav>
</template>

<style scoped lang="scss">
.navigation {
    margin-left: auto;

    &__list {
        display: flex;
        justify-content: space-between;
        gap: 40px;
        padding: 0;
        list-style: none;
    }

    &__link {
        text-decoration: none;
        padding: 2px 3px;
        transition: 0.2s color;

        &:hover {
            color: var(--el-color-primary);
        }
    }
}
</style>
