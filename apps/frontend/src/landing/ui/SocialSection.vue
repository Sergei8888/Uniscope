<script setup lang="ts">
import TgImg from '@/landing/assets/social/tg.png';
import YouTubeImg from '@/landing/assets/social/youtube.png';
import VkImg from '@/landing/assets/social/vk.png';
</script>

<template>
    <section class="social">
        <h2 class="social__header">Следите за новостями</h2>
        <ul class="social__list">
            <li class="social__item">
                <a href="https://t.me/uniscope_official">
                    <img :src="TgImg" alt="Telegram logo" />
                </a>
            </li>
            <li class="social__item">
                <a
                    href="https://www.youtube.com/channel/UCLP6cpXmocCC096IB9gEXKA"
                >
                    <img :src="YouTubeImg" alt="YouTube logo" />
                </a>
            </li>
            <li class="social__item">
                <a href="https://vk.com/public211021203">
                    <img :src="VkImg" alt="Vk logo" />
                </a>
            </li>
        </ul>
    </section>
</template>

<style scoped lang="scss">
.social {
    &__list {
        display: flex;
        justify-content: space-around;
        align-items: center;
        list-style: none;
    }

    &__item {
        width: 100px;

        img {
            width: 100%;

            object-fit: contain;
        }
    }
}
</style>
