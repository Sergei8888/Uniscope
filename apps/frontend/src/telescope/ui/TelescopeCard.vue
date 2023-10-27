<script setup lang="ts">
import { TelescopeDtoI } from '@uniscope/shared/vite';

import { ImgUrlResolver } from '@/shared/lib/img/img-url-resolver';
import TelescopeStatusIndicator from '@/telescope/ui/TelescopeStatusIndicator.vue';

defineProps<{
    telescope: TelescopeDtoI;
}>();
</script>

<template>
    <router-link
        class="telescope-card-link"
        :to="{ name: 'telescope-control', query: { id: telescope.id } }"
    >
        <section class="telescope-card">
            <img
                class="telescope-card__img"
                :src="ImgUrlResolver.resolveBackendPath(telescope.imgUrl)"
                alt="Telescope image"
            />
            <div class="telescope-card__content">
                <h2 class="telescope-card__name">
                    {{ telescope.name }}
                </h2>
                <p class="telescope-card__description">
                    {{ telescope.description }}
                </p>
                <div class="telescope-card__info-bar">
                    <span>lat: {{ telescope.latitude }}</span>
                    <span>lon: {{ telescope.longitude }}</span>
                    <TelescopeStatusIndicator
                        class="telescope-card__status-indicator"
                        :status="telescope.status"
                    />
                </div>
            </div>
        </section>
    </router-link>
</template>

<style scoped lang="scss">
.telescope-card-link {
    text-decoration: none;
}

.telescope-card {
    width: 500px;
    height: 250px;
    display: flex;
    gap: 15px;
    padding: 15px;
    border-radius: 7px;
    background-color: var(--el-fill-color-dark);
    transition: 0.2s background-color;

    &:hover {
        background-color: var(--el-fill-color-darker);
    }

    &__img {
        border-radius: 4px;
        width: 45%;
        object-fit: cover;
    }

    &__content {
        flex-grow: 1;
        display: flex;
        flex-direction: column;
    }

    &__name {
        font-weight: 500;
        margin-bottom: 10px;
    }

    &__info-bar {
        display: flex;
        gap: 10px;
        margin-top: auto;
    }

    &__status-indicator {
        margin-left: auto;
    }
}
</style>
