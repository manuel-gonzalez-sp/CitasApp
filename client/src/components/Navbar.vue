<template>
  <div class="navbar bg-base-300 px-8">
    <div class="navbar-start">
      <a class="text-xl">CitasApp</a>
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
