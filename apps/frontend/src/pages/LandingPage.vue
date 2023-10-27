<script lang="ts" setup>
import IntroBlock from '@/landing/ui/IntroBlock.vue';
import DescriptionBlock from '@/landing/ui/DescriptionSection.vue';
import AdvantagesBlock from '@/landing/ui/AdvantagesSection.vue';
import TeamList from '@/landing/ui/TeamList.vue';
import PartnersSection from '@/landing/ui/PartnersSection.vue';
import SocialSection from '@/landing/ui/SocialSection.vue';
</script>
<template>
    <div class="page-wrapper">
        <IntroBlock />
        <DescriptionBlock />
        <AdvantagesBlock
            class="page-wrapper__advantages--left"
            :advantages="[
                'Доступ к телескопам в любое время и из любого места',
                'Выбор телескопа по параметрам и расположению',
                'Управление телескопом через удобный веб-интерфейс',
                'Получение высококачественных изображений в разных спектрах',
                'Сохранение и обработка полученных астрофотографий',
                'Обмен опытом и впечатлениями с другими участниками проекта',
            ]"
            advantages-header="Преимущества для наблюдателей 👀"
        />
        <AdvantagesBlock
            class="page-wrapper__advantages--right"
            aligned="right"
            :advantages="[
                'Подключение телескопа к проекту без сложной настройки',
                'Получение дохода от предоставления доступа к телескопу',
                'Участие в научных и образовательных проектах',
                'Возможность наблюдать через другие телескопы в сети',
            ]"
            advantages-header="Преимущества для владельцев телескопов 🔭"
        />
        <TeamList />
        <p class="investor">
            Проект реализуется при финансовой поддержки Фонда содействия
            иновациям, договор №451ГССС15-L/78722 от 08.09.2022
        </p>
        <PartnersSection />
        <SocialSection />
    </div>
</template>

<style lang="scss" scoped>
.page-wrapper {
    display: flex;
    flex-direction: column;
    gap: 80px;

    &__advantages {
        &--right {
            margin-left: auto;
        }
    }
}
</style>
