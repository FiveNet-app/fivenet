import { StoreDefinition, defineStore } from 'pinia';

export interface SettingsState {
    version: string;
    locale: string;
    livemap: {
        markerSize: number;
        centerSelectedMarker: boolean;
        showUnitNames: boolean;
    };
    startpage: string;
}

export const useSettingsStore = defineStore('settings', {
    state: () =>
        ({
            version: __APP_VERSION__ as string,
            locale: 'de',
            livemap: {
                markerSize: 22,
                centerSelectedMarker: false,
                showUnitNames: false,
            },
            startpage: '/overview',
        }) as SettingsState,
    persist: true,
    actions: {
        setVersion(version: string): void {
            this.version = version;
        },
        setLocale(locale: string): void {
            this.locale = locale;
        },
    },
});

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useSettingsStore as unknown as StoreDefinition, import.meta.hot));
}