<script lang="ts" setup>
import { useRoute } from 'vue-router';

import { ref } from 'vue';

import TelescopeControlPanel from '@/telescope/ui/TelescopeControlPanel.vue';
import { useSignalling } from '@/telescope/model/useSignalling';

const telescopeId = useRoute().query.id;
if (!telescopeId) {
    throw new Error('Telescope ID not found');
}

const stream = ref<MediaStream | null>(null);
const { pc, dc } = useSignalling(Number(telescopeId));

pc.addEventListener('track', (event) => {
    console.log('Received track: ', event);
    stream.value = event.streams[0];
});

const whiteBalanse = ref(50);
</script>

<template>
    <div v-if="stream" class="control-page">
        <TelescopeControlPanel :pc="pc" :dc="dc" :stream="stream" />
        <div class="control-inputs">
            <div class="control-element">
                <span>Экспозиция</span>
                <ElInputNumber placeholder="ms" :min="1" />
            </div>
            <div class="control-element">
                <span>Баланс белого</span>
                <ElSlider
                    v-model="whiteBalanse"
                    type="range"
                    :min="0"
                    :max="100"
                />
            </div>
            <div class="control-element">
                <span>Синий</span>
                <ElInputNumber placeholder="pt" :min="1" />
            </div>
            <div class="control-element">
                <span>Красный</span>
                <ElInputNumber placeholder="pt" :min="1" />
            </div>
            <ElButton type="success">
                <a href="/img/current_image.jpg" download="current_image">
                    Cкачать изображение
                </a>
            </ElButton>
        </div>
    </div>
    <div
        v-else
        style="
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
        "
    >
        <p>Loading...</p>
    </div>
</template>

<style lang="scss" scoped>
.control-inputs {
    padding: 0 20px;
    display: grid;
    grid-template: auto / repeat(2, 1fr);
    gap: 20px;
    margin-top: 30px;
}

.control-element {
    display: flex;
    flex-direction: column;
    gap: 10px;
}
</style>
