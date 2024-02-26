<template>
  <div class="navbar bg-base-300 px-8">
    <div class="navbar-start">
      <div class="dropdown">
        <div tabindex="0" role="button" class="btn btn-ghost btn-circle">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7" />
          </svg>
        </div>
        <a class="text-xl">CitasApp</a>
        <ul tabindex="0" class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52">
          <li><a href="/">Landing</a></li>
          <li><a href="/profile">Perfil</a></li>
          <li><a href="/users">Usuarios</a></li>
        </ul>
      </div>
    </div>
    <div class="navbar-center">
      <a class="text-xl">{{ props.title }}</a>
    </div>
    <div class="navbar-end">
      <a v-if="$session" class="text-xl px-4">
        {{ $session.fullname }}
      </a>
      <button class="btn btn-outline" v-if="$session" @click="handleLogOut">
        Cerrar sesión
      </button>
      <a class="btn btn-outline" v-else href="/auth">
        Iniciar sesión
      </a>
    </div>


  </div>
</template>

<script setup lang="ts">
import { useStore } from '@nanostores/vue';
import { onMounted } from 'vue';
import { getSession, logOut, sessionStore } from '../services/session-service';

const props = defineProps(['title'])
const $session = useStore(sessionStore);

onMounted(() => {
  const session = getSession();
  if (!session) {
    window.location.href = '/auth'
  }
})

async function handleLogOut() {
  await logOut()
  window.location.href = '/'
}

</script>
