<script setup lang="ts">
import SignupLink from '@/auth/ui/SignupLink.vue';
</script>

<template>
    <section class="description">
        <div class="description__video">
            <iframe
                src="https://www.youtube.com/embed/Sf7Xg5AHcV0"
                title="Uniscope presentation video"
                frameborder="0"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                allowfullscreen
            ></iframe>
        </div>
        <div class="description__text">
            <p>
                Uniscope – это интернет-платформа для дистанционных
                астрономических наблюдений и астрофотографии. Uniscope
                объединяет владельцев телескопов и любителей астрономии в единую
                сеть, которая даёт возможность управлять астрономической
                техникой из любой точки мира. Присоединяйтесь к Uniscope и
                откройте для себя красоту звездного неба!
            </p>
            <SignupLink size="big" />
        </div>
    </section>
</template>

<style scoped lang="scss">
.description {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 30px;

    &__text {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 20px;
        max-width: 50%;
        font-size: 18px;
        text-align: center;

        p {
            margin: 0;
        }
    }

    &__video {
        width: 70%;
        aspect-ratio: 16/9;

        iframe {
            width: 100%;
            height: 100%;
        }
    }
}
</style>
