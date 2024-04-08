import { StatusDispatch } from '~~/gen/ts/resources/centrum/dispatches';
import { StatusUnit, Unit } from '~~/gen/ts/resources/centrum/units';

export type GroupedUnits = { status: StatusUnit; key: string; units: Unit[] }[];

export function dispatchStatusToFillColor(status: StatusDispatch | undefined): string {
    switch (status) {
        case StatusDispatch.UNSPECIFIED:
        case StatusDispatch.NEW:
        case StatusDispatch.UNASSIGNED:
        case StatusDispatch.UNIT_DECLINED:
            return '!text-error-600';
        case StatusDispatch.EN_ROUTE:
            return '!text-info-500';
        case StatusDispatch.ON_SCENE:
            return '!text-info-700';
        case StatusDispatch.NEED_ASSISTANCE:
            return '!text-warn-600';
        case StatusDispatch.COMPLETED:
            return '!text-success-600';
        case StatusDispatch.CANCELLED:
            return '!text-success-800';
        case StatusDispatch.ARCHIVED:
            return '!text-base-600';
        case StatusDispatch.UNIT_ACCEPTED:
            return '!text-info-600';
        default:
            return '!text-info-500';
    }
}

export function dispatchStatusToBGColor(status: StatusDispatch | undefined): string {
    switch (status) {
        case StatusDispatch.UNSPECIFIED:
        case StatusDispatch.NEW:
        case StatusDispatch.UNASSIGNED:
        case StatusDispatch.UNIT_DECLINED:
            return '!bg-error-600';
        case StatusDispatch.EN_ROUTE:
            return '!bg-info-500';
        case StatusDispatch.ON_SCENE:
            return '!bg-info-700';
        case StatusDispatch.NEED_ASSISTANCE:
            return '!bg-warn-600';
        case StatusDispatch.COMPLETED:
            return '!bg-success-600';
        case StatusDispatch.CANCELLED:
            return '!bg-success-800';
        case StatusDispatch.ARCHIVED:
            return '!bg-background';
        case StatusDispatch.UNIT_ACCEPTED:
            return '!bg-info-600';
        default:
            return '!bg-info-500';
    }
}

export const animateStates = [
    StatusDispatch.NEW.valueOf(),
    StatusDispatch.UNASSIGNED.valueOf(),
    StatusDispatch.NEED_ASSISTANCE.valueOf(),
];

export function dispatchStatusAnimate(status: StatusDispatch | undefined): boolean {
    return animateStates.includes((status ?? StatusDispatch.NEW).valueOf());
}

export function unitStatusToBGColor(status: StatusUnit | undefined): string {
    switch (status) {
        case StatusUnit.ON_BREAK:
        case StatusUnit.USER_ADDED:
        case StatusUnit.USER_REMOVED:
            return '!bg-info-500';
        case StatusUnit.AVAILABLE:
            return '!bg-success-600';
        case StatusUnit.BUSY:
            return '!bg-warn-600';
        case StatusUnit.UNSPECIFIED:
        case StatusUnit.UNKNOWN:
        case StatusUnit.UNAVAILABLE:
        default:
            return '!bg-error-600';
    }
}

export const statusOrder = [
    StatusUnit.AVAILABLE,
    StatusUnit.ON_BREAK,
    StatusUnit.BUSY,
    StatusUnit.USER_ADDED,
    StatusUnit.USER_REMOVED,
    StatusUnit.UNAVAILABLE,
    StatusUnit.UNKNOWN,
    StatusUnit.UNSPECIFIED,
];

export const unitStatuses: {
    icon: string;
    name: string;
    action?: Function;
    status?: StatusUnit;
}[] = [
    { icon: 'i-mdi-cancel', name: 'Unavailable', status: StatusUnit.UNAVAILABLE },
    { icon: 'i-mdi-calendar-check', name: 'Available', status: StatusUnit.AVAILABLE },
    { icon: 'i-mdi-coffee', name: 'On Break', status: StatusUnit.ON_BREAK },
    { icon: 'i-mdi-calendar-remove', name: 'Busy', status: StatusUnit.BUSY },
];

export const dispatchStatuses: {
    icon: string;
    name: string;
    action?: Function;
    status?: StatusDispatch;
}[] = [
    { icon: 'i-mdi-car-back', name: 'En Route', status: StatusDispatch.EN_ROUTE },
    { icon: 'i-mdi-marker-check', name: 'On Scene', status: StatusDispatch.ON_SCENE },
    { icon: 'i-mdi-help-circle', name: 'Need Assistance', status: StatusDispatch.NEED_ASSISTANCE },
    { icon: 'i-mdi-check-bold', name: 'Completed', status: StatusDispatch.COMPLETED },
    { icon: 'i-mdi-cancel', name: 'Cancelled', status: StatusDispatch.CANCELLED },
];

export function isStatusDispatchCompleted(status: StatusDispatch): boolean {
    return status === StatusDispatch.ARCHIVED || status === StatusDispatch.CANCELLED || status === StatusDispatch.COMPLETED;
}
