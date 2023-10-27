<script setup lang="ts">
defineProps<{
    imageUrl: string;
    initials: string;
    role: string;
}>();
</script>

<template>
    <div class="member">
        <img :src="imageUrl" class="member__image" alt="Team member image" />
        <div class="member__text">
            <span class="member__initials">
                {{ initials }}
            </span>
            <span class="member__role">
                {{ role }}
            </span>
        </div>
    </div>
</template>

<style scoped lang="scss">
.member {
    display: flex;
    gap: 20px;
    align-items: center;

    &__image {
        width: 125px;
        aspect-ratio: 1/1;
        border-radius: 50%;
    }

    &__text {
        display: flex;
        flex-direction: column;
        gap: 5px;
        font-size: 17px;
    }
}
</style>
