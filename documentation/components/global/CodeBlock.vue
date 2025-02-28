<template>
    <pre v-highlight><code :class="className" class="py-0">
{{ processedCode }}
    </code></pre>
</template>

<script setup lang="ts">

const processedCode = computed(() => {
    return code.replace(/\\n/g, '\n').trim();
})

const { language, code } = defineProps<{ code: string, language: string }>()

const className = computed(() => `language-${language}`)
</script>

<style>
code.hljs {
    padding: 0 1em !important;
}
</style>