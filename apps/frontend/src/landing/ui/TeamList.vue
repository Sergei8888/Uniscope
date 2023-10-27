<script setup lang="ts">
import TeamMember from '@/landing/ui/TeamMember.vue';

import AlexandrImg from '@/landing/assets/team/alexander.png';
import PavelImg from '@/landing/assets/team/pavel.jpg';
import SergeiImg from '@/landing/assets/team/sergei.png';
import PavelLilImg from '@/landing/assets/team/pavel_lil.png';
import MaximImg from '@/landing/assets/team/maxim.jpg';
import AlexImg from '@/landing/assets/team/alex.png';
</script>

<template>
    <section class="team">
        <h2 class="team__header">Команда стартап проекта</h2>
        <ul class="team__list">
            <li class="team__item">
                <TeamMember
                    :image-url="AlexandrImg"
                    initials="Тепляков Александр"
                    role="Генеральный директор"
                />
            </li>
            <li class="team__item">
                <TeamMember
                    :image-url="PavelImg"
                    initials="Васильев Павел"
                    role="Финансовый директор"
                />
            </li>
            <li class="team__item">
                <TeamMember
                    :image-url="SergeiImg"
                    initials="Кузьменков Сергей"
                    role="Лидер продуктовой команды"
                />
            </li>
            <li class="team__item">
                <TeamMember
                    :image-url="AlexImg"
                    initials="Байгашов Алексей"
                    role="Директор по апробации и внедрению"
                />
            </li>
            <li class="team__item">
                <TeamMember
                    :image-url="MaximImg"
                    initials="Царьков Максим"
                    role="Технический директор"
                />
            </li>
            <li class="team__item">
                <TeamMember
                    :image-url="PavelLilImg"
                    initials="Волков Павел"
                    role="Hardened разработчик"
                />
            </li>
        </ul>
    </section>
</template>

<style scoped lang="scss">
.team {
    &__list {
        display: grid;
        margin: 0;
        padding: 0;
        grid-template: auto / repeat(2, 375px);
        row-gap: 60px;
        column-gap: 120px;
        justify-content: left;
        list-style: none;
    }
}
</style>
