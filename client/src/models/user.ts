export interface User {
    id: number;
    username: string;
    fullname: string;
    birthDate: Date;
    gender: string;
    introduction: string;
    lookingFor: string;
    city: string;
    country: string;
    createdAt: Date;
    // photos
    // interests
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