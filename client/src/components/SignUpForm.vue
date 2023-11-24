<template>
    <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
        Crea una cuenta
    </h1>
    <form class="space-y-4 md:space-y-6" @submit.prevent="handleSignUp">

        <div class="form-control w-full max-w-xs">
            <label class="label">
                <span class="label-text">Nombre de usuario</span>
            </label>
            <input type="text" placeholder="yomero_kwamero" class="input input-bordered w-full max-w-xs"
                v-model="username" />
        </div>
        <div class="form-control w-full max-w-xs">
            <label class="label">
                <span class="label-text">Nombre completo</span>
            </label>
            <input type="text" placeholder="Yomero Kwamero" class="input input-bordered w-full max-w-xs"
                v-model="fullname" />
        </div>
        <div class="form-control w-full max-w-xs">
            <label class="label">
                <span class="label-text">Contraseña</span>
            </label>
            <input type="password" placeholder="••••••••" class="input input-bordered w-full max-w-xs" v-model="password" />
        </div>

        <div class="divider"></div>

        <button type="submit" class="btn btn-block btn-primary">Registrarse</button>

        <p class="text-sm font-light text-gray-500 dark:text-gray-400">
            Ya tienes una cuenta?
            <a class="font-medium text-primary hover:underline dark:text-primary-500" @click="$emit('toggleLogIn')">Ingresa
                aquí</a>
        </p>

    </form>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { signUp } from '../services/session-service';
defineEmits(['toggleLogIn'])

const username = ref('');
const fullname = ref('');
const password = ref('');

async function handleSignUp() {
    const data = await signUp({
        username: username.value,
        fullname: fullname.value,
        password: password.value,
    })
    if (data) {
        alert('Te registraste correctamente!')
    }
};
</script>
  