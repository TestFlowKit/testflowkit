<template>
    <header class="bg-blue-500 py-4 px-2 md:px-0">
        <div class="container mx-auto flex flex-col md:flex-row items-center justify-between text-white">
            <div class="w-full flex items-center justify-between md:w-auto md:inline">
                <nuxt-link tag="h1" to="/" class="text-3xl font-bold">Testflowkit</nuxt-link>
                <button class="md:hidden" @click="showMenu = !showMenu">
                    <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                        xmlns="http://www.w3.org/2000/svg">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 6h16M4 12h16m-7 6h7">
                        </path>
                    </svg>
                </button>
            </div>
            <nav class="md:block md:ml-4 mt-4 md:mt-0" :class="{ hidden: !showMenu }">
                <ul class="flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4">
                    <li><nuxt-link to="/" class="hover:underline">Home</nuxt-link></li>
                    <li><nuxt-link :to="{ name: 'get-started' }" class="hover:underline">Get
                            started</nuxt-link></li>
                    <li><nuxt-link :to="{ name: 'quick-start' }" class="hover:underline">Quick
                            Start</nuxt-link></li>
                    <li><nuxt-link :to="{ name: 'configuration' }" class="hover:underline">Configuration</nuxt-link>
                    </li>
                    <li><nuxt-link :to="{ name: 'sentences' }" class="hover:underline">Sentences</nuxt-link>
                    </li>
                    <li><a target="_blank" :href="githubUrl" class="hover:underline">Github</a></li>
                </ul>
            </nav>
        </div>
    </header>
</template>

<script setup lang="ts">
import { ref } from "vue";

const githubUrl = "https://github.com/TestFlowKit/testflowkit";
const showMenu = ref(false);
</script>