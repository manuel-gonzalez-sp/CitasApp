import { atom } from 'nanostores';
import type { LogIn, Session, SignUp } from '../models/session';
import type { ApiResponse } from '../utils/api-response';


const SESSION_KEY = 'session';

export const sessionStore = atom<Session | null>(null);

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
  sessionStore.set(null);
}

export function getSession(): Session | null {
  const data = localStorage.getItem(SESSION_KEY);
  if (data) {
    const session: Session = JSON.parse(data);
    sessionStore.set(session);
    return session;
  }
  return null;
}

async function setSession(response: Response): Promise<Session | null> {
  if (response.ok) {
    const res = await (response.json() as Promise<ApiResponse<Session>>);
    localStorage.setItem(SESSION_KEY, JSON.stringify(res.data[0]));
    sessionStore.set(res.data[0]);
    return res.data[0];
  }
  return null;
}