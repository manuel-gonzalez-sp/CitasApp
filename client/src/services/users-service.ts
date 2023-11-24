import type { User } from "../models/user";
import type { ApiResponse } from "../utils/api-response";
import { sessionStore } from "./session-service";


export async function getUsers(): Promise<User[] | null> {
    const response = await fetch(`${import.meta.env.PUBLIC_SERVER_URL}/user`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${sessionStore.value?.jwtToken}`
        },
    });
    if (response.ok) {
        const res = await (response.json() as Promise<ApiResponse<User>>);
        return res.data;
    }
    return null;
}

export async function getUserById(id: string): Promise<User | null> {
    const response = await fetch(`${import.meta.env.PUBLIC_SERVER_URL}/user/${id}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${sessionStore.value?.jwtToken}`
        },
    });
    if (response.ok) {
        const res = await (response.json() as Promise<ApiResponse<User>>);
        return res.data[0];
    }
    return null;
}