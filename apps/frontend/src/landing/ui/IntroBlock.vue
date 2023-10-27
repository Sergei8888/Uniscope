<script setup lang="ts">
import Logo from '@/shared/ui/header/assets/logo.png';
</script>

<template>
    <section class="intro">
        <ElImage :src="Logo" class="intro__logo" />
        <h1 class="intro__header">
            <span class="intro__header-main">Uniscope</span>
            <span class="intro__header-caption">
                Система удалённых астрономических наблюдений
            </span>
        </h1>
    </section>
</template>

<style scoped lang="scss">
.intro {
    display: flex;
    align-items: center;
    justify-content: center;

    &__logo {
        width: 25%;
        margin-right: 60px;
    }

    &__header {
        display: flex;
        flex-direction: column;
        text-align: center;
        font-weight: 500;

        &-main {
            font-size: 100px;
            text-transform: uppercase;
        }

        &-caption {
            font-size: 32px;
            margin-top: 20px;
            max-width: 600px;
        }
    }
}

@media (min-width: 1200px) {
    .intro {
        &__header {
            &-main {
                font-size: 120px;
            }

            &-caption {
                font-size: 35px;
            }
        }
    }
}
</style>
