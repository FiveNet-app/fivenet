import * as resources_timestamp_timestamp from '~~/gen/ts/resources/timestamp/timestamp';
import * as google_protobuf_timestamp from '~~/gen/ts/google/protobuf/timestamp';

const seconds_per_minute = 60;
const seconds_per_hour = seconds_per_minute * 60;
const seconds_per_day = seconds_per_hour * 24;
const seconds_per_week = seconds_per_day * 7;
const seconds_per_year = seconds_per_week * 52;

export function fromSecondsToFormattedDuration(seconds: number): string {
    const { t } = useI18n();

    const years = Math.floor(seconds / seconds_per_year);
    seconds -= years * seconds_per_year;
    const weeks = Math.floor(seconds / seconds_per_week);
    seconds -= weeks * seconds_per_week;
    const days = Math.floor(seconds / seconds_per_day);
    seconds -= days * seconds_per_day;
    const hours = Math.floor(seconds / seconds_per_hour);
    seconds -= hours * seconds_per_hour;
    const minutes = Math.floor(seconds / seconds_per_minute);
    seconds -= minutes * seconds_per_minute;

    const parts = new Array<string>();
    if (years > 0) {
        parts.push(`${years} ${t(`common.time_ago.year`, years)}`);
    }
    if (weeks > 0) {
        parts.push(`${weeks} ${t(`common.time_ago.week`, weeks)}`);
    }
    if (days > 0) {
        parts.push(`${days} ${t(`common.time_ago.day`, days)}`);
    }
    if (hours > 0) {
        parts.push(`${hours} ${t(`common.time_ago.hour`, hours)}`);
    }
    if (minutes > 0) {
        parts.push(`${minutes} ${t(`common.time_ago.minute`, minutes)}`);
    }
    if (seconds > 0) {
        parts.push(`${seconds} ${t(`common.time_ago.second`, seconds)}`);
    }
    return parts.join(', ');
}

export function toDate(ts: resources_timestamp_timestamp.Timestamp | undefined): undefined | Date {
    if (ts === undefined || ts?.timestamp === undefined) {
        return new Date();
    }
    return google_protobuf_timestamp.Timestamp.toDate(ts?.timestamp!);
}

export function toDateLocaleString(
    ts: resources_timestamp_timestamp.Timestamp | undefined,
    d?: Function
): undefined | string {
    if (typeof ts === undefined) {
        return '-';
    }

    if (d) {
        return d(google_protobuf_timestamp.Timestamp.toDate(ts?.timestamp!), 'short');
    }
    return google_protobuf_timestamp.Timestamp.toDate(ts?.timestamp!).toLocaleDateString();
}

export function fromString(time: string): undefined | Date {
    return new Date(Date.parse(time));
}
