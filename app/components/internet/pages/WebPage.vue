<script lang="ts" setup>
import type { Tab } from '~/store/internet';
import type { Page } from '~~/gen/ts/resources/internet/internet';

const props = defineProps<{
    modelValue?: Tab;
    page: Page;
}>();

const emit = defineEmits<{
    (e: 'update:modelValue', value: Tab): void;
}>();

const tab = useVModel(props, 'modelValue', emit);

function updateTabInfo(): void {
    if (!tab.value) {
        return;
    }

    tab.value.label = props.page.title ?? '';
    tab.value.icon = undefined;
}

watch(
    () => props.page,
    () => updateTabInfo(),
);
updateTabInfo();

onBeforeUnmount(() => console.log('unmount base page', props.page));
</script>

<template>
    <div>Page: {{ page }}</div>
</template>
