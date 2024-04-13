import { defineStore, type StoreDefinition } from 'pinia';

export interface SettingsState {
    version: string;
    updateAvailable: false | string;
    locale: string | null;

    nuiEnabled: boolean;
    nuiResourceName: string | undefined;
    livemap: {
        markerSize: number;
        centerSelectedMarker: boolean;
        showUnitNames: boolean;
        showUnitStatus: boolean;
        showAllDispatches: boolean;
        activeLayers: string[];
    };
    startpage: string;
    design: {
        documents: {
            editorTheme: 'default' | 'dark';
            listStyle: 'single' | 'double';
        };
        ui: {
            primary: string;
            gray: string;
        };
    };
    audio: {
        notificationsVolume: number;
    };
    streamerMode: boolean;
    featureGates: {};
}

export const useSettingsStore = defineStore('settings', {
    state: () =>
        ({
            version: __APP_VERSION__,
            updateAvailable: false,
            locale: null,

            nuiEnabled: false,
            nuiResourceName: undefined,

            livemap: {
                markerSize: 22,
                centerSelectedMarker: false,
                showUnitNames: true,
                showUnitStatus: true,
                showAllDispatches: false,
                activeLayers: [],
            },
            startpage: '/overview',
            design: {
                documents: {
                    editorTheme: 'default',
                    listStyle: 'double',
                },
                ui: {
                    primary: 'sky',
                    gray: 'neutral',
                },
            },
            audio: {
                notificationsVolume: 0.15,
            },
            streamerMode: false,
            featureGates: {},
        }) as SettingsState,
    persist: {
        paths: [
            'version',
            'locale',
            'nuiEnabled',
            'nuiResourceName',
            'livemap',
            'documents',
            'startpage',
            'theme',
            'audio',
            'streamerMode',
            'featureGates',
        ],
    },
    actions: {
        setVersion(version: string): void {
            this.version = version;
        },
        async setUpdateAvailable(version: string): Promise<void> {
            this.updateAvailable = version;
        },
        setNuiDetails(enabled: boolean, resourceName: string | undefined): void {
            this.nuiEnabled = enabled;
            this.nuiResourceName = resourceName;
        },
        setLocale(locale: string): void {
            this.locale = locale;
        },
    },
    getters: {
        isNUIAvailable(state): boolean {
            return state.nuiEnabled ?? false;
        },
    },
});

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useSettingsStore as unknown as StoreDefinition, import.meta.hot));
}
