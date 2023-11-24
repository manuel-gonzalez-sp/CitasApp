import { ref } from 'vue';
import type { LogIn, Session, SignUp } from '../models/session';
import type { ApiResponse } from '../utils/api-response';


const SESSION_KEY = 'session';

export const sessionStore = ref<Session | null>(null);

export async function logIn(input: LogIn): Promise<Session | null> {
  const response = await fetch(`${import.meta.env.PUBLIC_SERVER_URL}/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(input)
  });
  return await setSession(response);
}


export async function signUp(input: SignUp): Promise<Session | null> {
  const response = await fetch(`${import.meta.env.PUBLIC_SERVER_URL}/signup`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(input)
  });
  return await setSession(response);
}

export async function logOut() {
  localStorage.removeItem(SESSION_KEY)
  sessionStore.value = null;
}

async function setSession(response: Response): Promise<Session | null> {
  if (response.ok) {
    const res = await (response.json() as Promise<ApiResponse<Session>>);
    localStorage.setItem(SESSION_KEY, JSON.stringify(res.data[0]))
    sessionStore.value = res.data[0];
    return res.data[0];
  }
  return null;
}