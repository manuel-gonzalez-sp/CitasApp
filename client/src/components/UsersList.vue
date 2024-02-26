<template>
    <div class="container mx-auto p-4">
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
            <div v-for="user in users" :key="user.id" class="border p-4 rounded-lg shadow hover:shadow-md transition">
                <div class="mb-2">
                    <img :src="user.photos?.find(photo => photo.isMain)?.url" alt="User Photo"
                        class="w-full h-48 object-cover rounded">
                </div>
                <h2 class="text-xl font-semibold">{{ user.fullname }}</h2>
                <p class="text-gray-600">{{ user.username }}</p>
                <p class="text-gray-500 text-sm">{{ user.introduction }}</p>
                <div class="mt-2">
                    <span
                        class="bg-blue-100 text-blue-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-blue-200 dark:text-blue-800">{{
                            user.city }}, {{ user.country }}</span>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script setup lang="ts">
import { onMounted, ref } from 'vue';
import type { User } from '../models/user';
import { getUsers } from '../services/users-service';


const users = ref<User[]>([]);

onMounted(async () => {
    const data = await getUsers();
    if (!data) {
        alert('No se pudo cargar la lista de usuarios')
    } else {
        users.value = data
    }
});

</script>