import { ref } from 'vue';
import type { LogIn, Session } from '../models/session';


export const sessionStore = ref<Session | null>(null);

interface Response<T> {
  data: T[];
  meta: any;
}

export async function logIn(input: LogIn): Promise<Session> {
  const response = await fetch(`${import.meta.env.PUBLIC_SERVER_URL}/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(input)
  });

  if (!response.ok) {
    throw new Error('Network response was not ok');
  }

  const res = await (response.json() as Promise<Response<Session>>);

  console.log(res.data[0])

  localStorage.setItem('session', JSON.stringify(res.data[0]))
  sessionStore.value = res.data[0];

  return res.data[0];
}



// login
// signup
// logout