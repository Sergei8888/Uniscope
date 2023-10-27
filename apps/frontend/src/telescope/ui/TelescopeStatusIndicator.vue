<script setup lang="ts">
import { TelescopeStatus } from '@uniscope/shared/vite';

defineProps<{
    status: TelescopeStatus;
}>();
</script>

<template>
    <div
        class="telescope-status-indicator"
        :class="{
            'telescope-status-indicator--offline':
                status === TelescopeStatus.Offline,
            'telescope-status-indicator--online':
                status === TelescopeStatus.Online,
            'telescope-status-indicator--busy': status === TelescopeStatus.Busy,
        }"
    ></div>
</template>

<style scoped lang="scss">
.telescope-status-indicator {
    width: 20px;
    height: 20px;
    border-radius: 50%;
    transition: 0.2s background-color;

    &--offline {
        background-color: var(--baseRed);
    }

    &--online {
        background-color: var(--baseGreen);
    }

    &--busy {
        background-color: var(--baseYellow);
    }
}
</style>
