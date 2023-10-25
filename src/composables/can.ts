import { useAuthStore } from '~/store/auth';
import slug from '~/utils/slugify';

export function can(perm: string | string[]): boolean {
    const permissions = useAuthStore().permissions;
    if (permissions.includes('superuser')) {
        return true;
    } else {
        const input: String[] = [];
        if (typeof perm === 'string') {
            input.push(perm);
        } else {
            const vals = perm as String[];
            input.push(...vals);
        }

        // Iterate over permissions and check in "OR" condition manner
        for (let index = 0; index < input.length; index++) {
            const val = slug(input[index] as string);
            if (permissions && (permissions.includes(val) || val === '')) {
                return true;
            }
        }
    }

    return false;
}

export function attr(perm: string, name: string, val: string): boolean {
    return can(perm + '.' + name + (val !== undefined ? '.' + val : ''));
}

export function attrList(perm: string, name: string): string[] {
    const key = slug(perm + '.' + name + '.');
    const permissions = useAuthStore().permissions;
    return permissions.filter((p) => p.startsWith(key)).map((p) => p.substring(key.length + 1));
}
