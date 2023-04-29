import { NavigationGuard, RouteLocationNormalized } from 'vue-router';
import { useAuthStore } from '~/store/auth';
import { CheckTokenRequest } from '@fivenet/gen/services/auth/auth_pb';
import { RpcError } from 'grpc-web';

export default defineNuxtRouteMiddleware(
    (to: RouteLocationNormalized, from: RouteLocationNormalized): ReturnType<NavigationGuard> => {
        const authStore = useAuthStore();

        const expiration = authStore.getAccessTokenExpiration;
        // Check if we have an expiration time, make sure the token isn't expired (yet)
        if (expiration !== null) {
            const now = new Date();
            // Token expired, redirect to login
            if (expiration <= now) {
                console.log('TOKEN EXPIRED');

                authStore.setAccessToken(null, null);

                // Only update the redirect query param if it isn't set already
                const redirect = to.query.redirect ?? to.fullPath;
                return navigateTo({
                    name: 'auth-login',
                    // save the location we were at to come back later
                    query: { redirect: redirect },
                });
            }

            // If expiration is in range of less than 4 hours, check token against server
            const renewTime = 4 * 60 * 60 * 1000;
            if (expiration.valueOf() - now.valueOf() < renewTime && !tokenCheckInProgress) {
                tokenCheckInProgress = true;
                setTimeout(async () => {
                    await checkToken();
                    tokenCheckInProgress = false;
                }, 1);
            }

            return;
        }
    }
);

let tokenCheckInProgress = false;

async function checkToken(): Promise<void> {
    return new Promise(async (res, rej) => {
        const { $grpc } = useNuxtApp();
        const authStore = useAuthStore();

        const req = new CheckTokenRequest();
        req.setToken(authStore.getAccessToken!);
        try {
            const resp = await $grpc.getAuthClient().checkToken(req, null);

            if (resp.hasNewToken() && resp.hasExpires()) {
                console.log('Token successfully renewed');
                authStore.setAccessToken(resp.getNewToken(), toDate(resp.getExpires()) as null | Date);
            }

            return res();
        } catch (e) {
            $grpc.handleRPCError(e as RpcError);
            return rej(e as RpcError);
        }
    });
}
