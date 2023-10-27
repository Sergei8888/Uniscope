<script setup lang="ts">
import FasiImg from '@/landing/assets/partners/fasi.png';
import AstromodelImg from '@/landing/assets/partners/astromodel.png';
import BfuImg from '@/landing/assets/partners/bfu.png';
import Lyceum23Img from '@/landing/assets/partners/23lyceum.png';
import FutureSchoolImg from '@/landing/assets/partners/future-school.png';
</script>

<template>
    <section class="partners">
        <h2 class="partners__header">Партнёры стартап проекта</h2>
        <ul class="partners__list">
            <li class="partners__partner">
                <a href="https://fasie.ru/" class="partners__link">
                    <img
                        :src="FasiImg"
                        alt="Фонд содействия инновациям логотип"
                    />
                </a>
            </li>
            <li class="partners__partner">
                <a href="https://astromodel.ru/" class="partners__link">
                    <img
                        :src="AstromodelImg"
                        alt="Фонд содействия инновациям логотип"
                    />
                </a>
            </li>
            <li class="partners__partner">
                <a href="https://kantiana.ru/" class="partners__link">
                    <img :src="BfuImg" alt="БФУ имени канта логотип" />
                </a>
            </li>
            <li class="partners__partner">
                <a href="https://maoulic23.ru/" class="partners__link">
                    <img :src="Lyceum23Img" alt="23 лицей логотип" />
                </a>
            </li>
            <li class="partners__partner">
                <a href="http://isakovo-shkola.ru/" class="partners__link">
                    <img :src="FutureSchoolImg" alt="Школа будущего логотип" />
                </a>
            </li>
        </ul>
    </section>
</template>

<style scoped lang="scss">
.partners {
    &__list {
        display: flex;
        flex-wrap: wrap;
        gap: 10%;
        justify-content: center;
        padding: 0;
        margin: 0;
        list-style: none;
    }

    &__partner {
        display: flex;
        align-items: center;
        width: 25%;
        margin-bottom: 60px;

        img {
            width: 100%;
            object-fit: contain;
        }
    }
}
</style>
