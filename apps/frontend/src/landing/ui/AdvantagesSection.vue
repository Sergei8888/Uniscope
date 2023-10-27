<script setup lang="ts">
withDefaults(
    defineProps<{
        advantagesHeader: string;
        advantages: Array<string>;
        aligned?: 'left' | 'right';
    }>(),
    {
        aligned: 'left',
    }
);
</script>

<template>
    <section
        class="advantages"
        :class="{
            'advantages--left': aligned === 'left',
            'advantages--right': aligned === 'right',
        }"
    >
        <h2 class="advantages__header">
            {{ advantagesHeader }}
        </h2>
        <ul class="advantages__list">
            <li
                v-for="(advantage, i) in advantages"
                :key="i"
                class="advantages__item"
            >
                {{ advantage }}
            </li>
        </ul>
    </section>
</template>

<style scoped lang="scss">
.advantages {
    &__header {
        margin-bottom: 20px;
    }

    &__list {
        display: flex;
        flex-direction: column;
        gap: 10px;
        margin: 0;
    }

    &__item {
        font-size: 17px;
    }

    &--right {
        text-align: right;

        .advantages__list {
            direction: rtl;
        }
    }
}
</style>
