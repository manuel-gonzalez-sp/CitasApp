<template>
    <div class="card w-96 bg-base-100 shadow-xl">
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
            <LogInForm v-if="isLogIn" @toggleLogIn="isLogIn = false"></LogInForm>
            <SignUpForm v-else @toggleLogIn="isLogIn = true"></SignUpForm>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { getSession } from '../services/session-service';
import LogInForm from './LogInForm.vue';
import SignUpForm from './SignUpForm.vue';

const isLogIn = ref(true);

onMounted(() => {
    const session = getSession();
    if (session) {
        window.location.href = '/'
    }
})

</script>