import type { CreateUser, User } from "./user";

export interface LogIn {
    username: string;
    password: string;
}

export interface SignUp extends CreateUser {
    password: string;
}

export interface Session extends User {
    jwtToken: string;
}