<template>
    <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
        Ingresa a tu cuenta
    </h1>
    <form class="space-y-4 md:space-y-6" @submit.prevent="handleLogIn">

        <div class="form-control w-full max-w-xs">
            <label class="label">
                <span class="label-text">Nombre de usuario</span>
            </label>
            <input type="text" placeholder="YomeroKwamero" class="input input-bordered w-full max-w-xs"
                v-model="username" />
        </div>
        <div class="form-control w-full max-w-xs">
            <label class="label">
                <span class="label-text">Contraseña</span>
            </label>
            <input type="password" placeholder="••••••••" class="input input-bordered w-full max-w-xs" v-model="password" />
        </div>

        <div class="divider"></div>

        <button type="submit" class="btn btn-block btn-primary">Ingresar</button>

        <p class="text-sm font-light text-gray-500 dark:text-gray-400">
            Todavía no tienes una cuenta?
            <a class="font-medium text-primary hover:underline dark:text-primary-500" @click="$emit('toggleLogIn')">Crea
                una</a>
        </p>

    </form>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { logIn } from '../services/session-service';
defineEmits(['toggleLogIn'])

const username = ref('');
const password = ref('');

async function handleLogIn() {
    const data = await logIn({
        username: username.value,
        password: password.value,
    })
    if (data) {
        alert('Ingresaste correctamente!')
    }
};
</script>
  