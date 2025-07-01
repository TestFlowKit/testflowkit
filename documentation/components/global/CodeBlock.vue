<template>
    <div class="relative group">
        <pre v-highlight class="relative">
            <code :class="className" class="py-0">
{{ processedCode }}
            </code>
        </pre>

        <button @click="copyToClipboard"
            class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200 bg-gray-800 hover:bg-gray-700 text-white px-3 py-1 rounded text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            :class="{ 'bg-green-600 hover:bg-green-700': copied }">
            <span v-if="!copied" class="flex items-center gap-1">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z">
                    </path>
                </svg>
                Copy
            </span>
            <span v-else class="flex items-center gap-1">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
                Copied!
            </span>
        </button>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const copied = ref(false)

const processedCode = computed(() => {
    return code.replace(/\\n/g, '\n').trim();
})

const { language, code } = defineProps<{ code: string, language: string }>()

const className = computed(() => `language-${language}`)

const copyToClipboard = async () => {
    try {
        await navigator.clipboard.writeText(processedCode.value)
        copied.value = true
        setTimeout(() => {
            copied.value = false
        }, 2000)
    } catch (err) {
        console.error('Failed to copy code:', err)
    }
}
</script>

<style>
code.hljs {
    padding: 0 1em !important;
}
</style>