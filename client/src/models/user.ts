export interface User {
    id: string;
    username: string;
    fullname: string;
    birthDate?: Date;
    gender?: string;
    introduction?: string;
    lookingFor?: string;
    city?: string;
    country?: string;
    createdAt: Date;
    photos?: Photo[];
    // interests
}

export interface Photo {
    id: string;
    url: string;
    isMain: boolean;
}

export interface CreateUser {
    username: string;
    fullname: string;
    birthDate?: Date;
    gender?: string;
    introduction?: string;
    lookingFor?: string;
    city?: string;
    country?: string;
}


export interface UpdateUser {
    fullname?: string;
    birthDate?: Date;
    gender?: string;
    introduction?: string;
    lookingFor?: string;
    city?: string;
    country?: string;
}