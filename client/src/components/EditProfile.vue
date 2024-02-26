<template>
    <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
        Editar perfil
    </h1>
    <form class="space-y-4 md:space-y-6" @submit.prevent="handleSaveChanges">

        <div class="form-control w-full">
            <label class="label">
                <span class="label-text">Nombre completo</span>
            </label>
            <input type="text" placeholder="Yomero Kwamero" class="input input-bordered w-full" v-model="fullname" />
        </div>
        <div class="form-control w-full">
            <label class="label">
                <span class="label-text">Introducción</span>
            </label>
            <textarea class="textarea textarea-bordered" placeholder="Introducción..." v-model="introduction"></textarea>
        </div>
        <div class="form-control w-full">
            <label class="label">
                <span class="label-text">Ciudad</span>
            </label>
            <input type="text" placeholder="Aguascalientes..." class="input input-bordered w-full" v-model="city" />
        </div>
        <div class="form-control w-full ">
            <label class="label">
                <span class="label-text">País</span>
            </label>
            <input type="text" placeholder="México..." class="input input-bordered w-full" v-model="country" />
        </div>

        <div class="divider"></div>

        <button type="submit" class="btn btn-block btn-primary">Guardar cambios</button>

    </form>
</template>

<script setup lang="ts">
import { useStore } from '@nanostores/vue';
import { ref } from 'vue';
import type { UpdateUser } from '../models/user';
import { sessionStore } from '../services/session-service';
import { updateUserById } from '../services/users-service';

const $session = useStore(sessionStore);

const fullname = ref($session.value?.fullname);
const introduction = ref($session.value?.introduction);
const lookingFor = ref($session.value?.lookingFor)
const city = ref($session.value?.lookingFor)
const country = ref($session.value?.country)

async function handleSaveChanges() {
    const id = $session.value!.id
    const data = await updateUserById(id, {
        fullname: fullname.value,
        introduction: introduction.value,
        lookingFor: lookingFor.value,
        city: city.value,
        country: country.value,
    } as UpdateUser);
    if (data) {
        alert('Usuario actualizado correctamente!')
    } else {
        alert('Hubo un error al actualizar!')
    }
};
</script>
  
