import { defineStore, type StoreDefinition } from 'pinia';
import { Category } from '~~/gen/ts/resources/documents/category';
import { DocContentType, Document, DocumentShort } from '~~/gen/ts/resources/documents/documents';
import { ObjectSpecs, TemplateData } from '~~/gen/ts/resources/documents/templates';
import { User, UserShort } from '~~/gen/ts/resources/users/users';
import { Vehicle } from '~~/gen/ts/resources/vehicles/vehicles';

export interface ClipboardData {
    documents: ClipboardDocument[];
    users: ClipboardUser[];
    vehicles: ClipboardVehicle[];
}

export interface ClipboardState extends ClipboardData {
    activeStack: ClipboardData;
}

export type ListType = 'users' | 'documents' | 'vehicles';

export const useClipboardStore = defineStore('clipboard', {
    state: () =>
        ({
            users: [],
            documents: [],
            vehicles: [],
            activeStack: {
                users: [],
                documents: [],
                vehicles: [],
            } as ClipboardData,
        }) as ClipboardState,
    persist: {
        serializer: jsonSerializer,
    },
    actions: {
        getTemplateData(): TemplateData {
            const data: TemplateData = {
                documents: [],
                users: [],
                vehicles: [],
            };

            this.activeStack.documents.forEach((v: ClipboardDocument) => {
                if (v !== undefined) data.documents.push(getDocument(v));
            });
            this.activeStack.users.forEach((v: ClipboardUser) => {
                if (v !== undefined) data.users.push(getUser(v));
            });
            this.activeStack.vehicles.forEach((v: ClipboardVehicle) => {
                if (v !== undefined) data.vehicles.push(getVehicle(v));
            });

            return data;
        },

        promoteToActiveStack(listType: ListType): void {
            switch (listType) {
                case 'documents':
                    this.activeStack.documents = jsonParse(jsonStringify(this.documents)) as ClipboardDocument[];
                    break;
                case 'users':
                    this.activeStack.users = jsonParse(jsonStringify(this.users)) as ClipboardUser[];
                    break;
                case 'vehicles':
                    this.activeStack.vehicles = jsonParse(jsonStringify(this.vehicles)) as ClipboardVehicle[];
                    break;
            }
        },

        clearActiveStack(): void {
            this.activeStack.documents.length = 0;
            this.activeStack.users.length = 0;
            this.activeStack.vehicles.length = 0;
        },

        // Documents
        addDocument(document: Document): void {
            const idx = this.documents.findIndex((o: ClipboardDocument) => {
                return o.id === document.id;
            });
            if (idx === -1) {
                this.documents.unshift(new ClipboardDocument(document));
            }
        },
        removeDocument(id: bigint): void {
            const idx = this.documents.findIndex((o: ClipboardDocument) => {
                return o.id === id;
            });
            if (idx > -1) {
                this.documents.splice(idx, 1);
            }
        },
        clearDocuments(): void {
            this.documents.splice(0, this.documents.length);
        },

        // Users
        addUser(user: User): void {
            const idx = this.users.findIndex((o: ClipboardUser) => {
                return o.userId === user.userId;
            });
            if (idx === -1) {
                this.users.unshift(new ClipboardUser(user!));
            }
        },
        removeUser(id: number): void {
            const idx = this.users.findIndex((o: ClipboardUser) => {
                return o.userId === id;
            });
            if (idx > -1) {
                this.users.splice(idx, 1);
            }
        },
        clearUsers(): void {
            this.users.splice(0, this.users.length);
        },

        // Vehicles
        addVehicle(vehicle: Vehicle): void {
            const idx = this.vehicles.findIndex((o: ClipboardVehicle) => {
                return o.plate === vehicle.plate;
            });
            if (idx === -1) {
                this.vehicles.unshift(new ClipboardVehicle(vehicle));
            }
        },
        removeVehicle(plate: string): void {
            const idx = this.vehicles.findIndex((o: ClipboardVehicle) => {
                return o.plate === plate;
            });
            if (idx > -1) {
                this.vehicles.splice(idx, 1);
            }
        },
        clearVehicles(): void {
            this.vehicles.splice(0, this.vehicles.length);
        },

        clear(): void {
            this.clearDocuments();
            this.clearUsers();
            this.clearVehicles();
            this.clearActiveStack();
        },

        checkRequirements(reqs: ObjectSpecs, listType: ListType): boolean {
            const length = BigInt(this[listType].length);
            if (reqs.required && length <= 0n) {
                return false;
            } else if (reqs.min && length < reqs.min && reqs.max && length > reqs.max) {
                return false;
            }

            return true;
        },
    },
});

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useClipboardStore as unknown as StoreDefinition, import.meta.hot));
}

export class ClipboardUser {
    public userId: number | undefined;
    public identifier: string | undefined;
    public job: string | undefined;
    public jobLabel: string | undefined;
    public jobGrade: number | undefined;
    public jobGradeLabel: string | undefined;
    public firstname: string | undefined;
    public lastname: string | undefined;
    public dateofbirth: string | undefined;

    constructor(u: UserShort | User) {
        this.userId = u.userId;
        this.identifier = u.identifier;
        this.job = u.job;
        this.jobLabel = u.jobLabel;
        this.jobGrade = u.jobGrade;
        this.jobGradeLabel = u.jobGradeLabel;
        this.firstname = u.firstname;
        this.lastname = u.lastname;
        this.dateofbirth = u.dateofbirth;

        return this;
    }
}

export function getUser(obj: ClipboardUser): User {
    const u: User = {
        userId: obj.userId!,
        identifier: obj.identifier!,
        job: obj.job!,
        jobLabel: obj.jobLabel!,
        jobGrade: obj.jobGrade!,
        jobGradeLabel: obj.jobGradeLabel!,
        firstname: obj.firstname!,
        lastname: obj.lastname!,
        dateofbirth: obj.dateofbirth!,
        licenses: [],
    };

    return u;
}

export class ClipboardDocument {
    public id: bigint;
    public createdAt?: string;
    public title: string;
    public creator: ClipboardUser;
    public category: Category | undefined;
    public state: string;
    public closed: boolean;
    public public: boolean;

    constructor(d: Document) {
        this.id = d.id;
        this.createdAt = d.createdAt ? toDate(d.createdAt).toJSON() : undefined;
        this.category = d.category;
        this.title = d.title;
        this.state = d.state;
        this.creator = new ClipboardUser(d.creator!);
        this.closed = d.closed;
        this.public = d.public;
    }
}

export function getDocument(obj: ClipboardDocument): DocumentShort {
    const user = getUser(obj.creator);

    const doc: DocumentShort = {
        id: BigInt(obj.id),
        categoryId: obj.category && obj.category.id ? obj.category.id : 0n,
        category: obj.category,
        title: obj.title,
        contentType: DocContentType.HTML,
        content: '',
        creatorId: user.userId,
        creator: user,
        state: obj.state,
        closed: obj.closed,
        public: obj.public,
    };
    if (obj.createdAt !== undefined) {
        doc.createdAt = toTimestamp(fromString(obj.createdAt));
    }
    return doc;
}

export class ClipboardVehicle {
    public plate: string;
    public model: string;
    public type: string;
    public owner: ClipboardUser;

    constructor(v: Vehicle) {
        this.plate = v.plate;
        this.model = v.model;
        this.type = v.type;
        this.owner = new ClipboardUser(v.owner!);
    }
}

export function getVehicle(obj: ClipboardVehicle): Vehicle {
    return {
        plate: obj.plate,
        model: obj.model,
        type: obj.type,
        owner: getUser(obj.owner),
    };
}
