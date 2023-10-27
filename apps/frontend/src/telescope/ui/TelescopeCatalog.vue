<script setup lang="ts">
import { ref } from 'vue';

import { useTelescopeStore } from '@/telescope/model/telescope.store';
import TelescopeCard from '@/telescope/ui/TelescopeCard.vue';

const telescopeStore = useTelescopeStore();
const telescopesLoading = ref(true);
telescopeStore.fetchTelescopes().then(() => {
    telescopesLoading.value = false;
});
</script>

<template>
    <ul v-if="!telescopesLoading" class="telescope-catalog">
        <li
            v-for="telescope in telescopeStore.telescopes"
            :key="telescope.id"
            class="telescope-catalog__item"
        >
            <TelescopeCard :telescope="telescope" />
        </li>
    </ul>
    <ElSkeleton
        v-else
        class="telescope-catalog telescope-catalog--skeleton"
        animated
        loading
    >
        <template #template>
            <ElSkeletonItem
                class="telescope-catalog__item telescope-catalog__item--skeleton"
                variant="image"
            />
            <ElSkeletonItem
                class="telescope-catalog__item telescope-catalog__item--skeleton"
                variant="image"
            />
            <ElSkeletonItem
                class="telescope-catalog__item telescope-catalog__item--skeleton"
                variant="image"
            />
            <ElSkeletonItem
                class="telescope-catalog__item telescope-catalog__item--skeleton"
                variant="image"
            />
        </template>
    </ElSkeleton>
</template>

<style scoped lang="scss">
@use 'sass:map';
@import '@/app/ui/scss/scss_variables.scss';

.telescope-catalog {
    display: flex;
    flex-wrap: wrap;
    gap: 40px;
    list-style: none;
    margin: 0;
    padding: 0;

    &__item {
        &--skeleton {
            width: 500px;
            height: 300px;
        }
    }
}

@media (max-width: map.get($breakpoints, xl)) {
    .telescope-catalog {
        &__item {
            &--skeleton {
                width: 400px;
                height: 300px;
            }
        }
    }
}
</style>
