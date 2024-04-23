import { createSharedComposable } from '@vueuse/core';

const _useDashboard = () => {
    const route = useRoute();
    const router = useRouter();
    const isHelpSlideoverOpen = ref(false);
    const isNotificationsSlideoverOpen = ref(false);

    defineShortcuts({
        'g-h': () => router.push('/'),
        'g-c': () => router.push('/citizens'),
        'g-v': () => router.push('/vehicles'),
        'g-d': () => router.push('/documents'),
        'g-j': () => router.push('/jobs'),
        'g-q': () => router.push('/qualifications'),
        'g-m': () => router.push('/livemap'),
        'g-w': () => router.push('/centrum'),
        'g-p': () => router.push('/rector'),
        '?': () => (isHelpSlideoverOpen.value = true),
        b: () => (isNotificationsSlideoverOpen.value = true),
    });

    watch(
        () => route.fullPath,
        () => {
            isHelpSlideoverOpen.value = false;
            isNotificationsSlideoverOpen.value = false;
        },
    );

    return {
        isHelpSlideoverOpen,
        isNotificationsSlideoverOpen,
    };
};

export const useDashboard = createSharedComposable(_useDashboard);
