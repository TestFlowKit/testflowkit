<template>
  <div class="code-block-container">
    <!-- Code Block Header -->
    <div v-if="title || language || showCopy" class="code-block-header">
      <div class="flex items-center space-x-3">
        <!-- Language Badge -->
        <span v-if="language" class="language-badge">
          <svg class="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/>
          </svg>
          {{ language }}
        </span>
        
        <!-- Title -->
        <span v-if="title" class="code-title">{{ title }}</span>
      </div>
      
      <!-- Copy Button -->
      <button 
        v-if="showCopy"
        @click="copyCode"
        class="copy-button"
        :class="{ 'copied': copied }"
      >
        <svg v-if="!copied" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
        </svg>
        <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
        </svg>
        <span>{{ copied ? 'Copied!' : 'Copy' }}</span>
      </button>
    </div>
    
    <!-- Code Content -->
    <div class="code-block-content">
      <pre 
        ref="codeElement"
        class="code-content"
        :class="{ 'rounded-t-none': title || language || showCopy }"
      ><code 
        ref="codeBlock"
        :class="language ? `language-${language}` : ''"
        v-text="codeText"
      ></code></pre>
      
      <!-- Line Numbers (if enabled) -->
      <div v-if="showLineNumbers" class="line-numbers">
        <div v-for="line in lineCount" :key="line">
          {{ line }}
        </div>
      </div>
    </div>
    
    <!-- Bottom Actions -->
    <div v-if="showActions" class="code-block-footer">
      <div class="flex items-center space-x-4">
        <span>{{ lineCount }} lines</span>
        <span v-if="charCount">{{ charCount }} characters</span>
      </div>
      <div class="flex items-center space-x-2">
        <button 
          v-if="allowWrap"
          @click="toggleWrap"
          class="wrap-button"
          :class="{ 'active': isWrapped }"
        >
          {{ isWrapped ? 'No Wrap' : 'Wrap' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, nextTick, onMounted, useSlots, watch } from 'vue'
import hljs from 'highlight.js'

interface Props {
  code?: string
  language?: string
  title?: string
  showCopy?: boolean
  showLineNumbers?: boolean
  showActions?: boolean
  allowWrap?: boolean
  highlightLines?: number[]
  maxHeight?: string
}

const props = withDefaults(defineProps<Props>(), {
  code: '',
  language: '',
  title: '',
  showCopy: true,
  showLineNumbers: false,
  showActions: false,
  allowWrap: true,
  highlightLines: () => [],
  maxHeight: 'none'
})

const slots = useSlots()

// Refs
const codeElement = ref<HTMLElement>()
const codeBlock = ref<HTMLElement>()
const copied = ref(false)
const isWrapped = ref(false)

// Computed properties
const rawCode = computed(() => {
  // Use prop code if provided, otherwise use slot content
  if (props.code) {
    return props.code
  }
  
  // For slot content, we need to get the text content
  if (slots.default) {
    const slotContent = slots.default()
    return slotContent.map(node => {
      if (typeof node.children === 'string') {
        return node.children
      }
      return ''
    }).join('')
  }
  
  return ''
})

const codeText = computed(() => {
  return rawCode.value
})

const lineCount = computed(() => {
  return codeText.value.split('\n').length
})

const charCount = computed(() => {
  return codeText.value.length
})

// Methods
const copyCode = async () => {
  try {
    await navigator.clipboard.writeText(codeText.value)
    copied.value = true
    setTimeout(() => {
      copied.value = false
    }, 2000)
  } catch (err) {
    console.error('Failed to copy code:', err)
    // Fallback for older browsers
    const textArea = document.createElement('textarea')
    textArea.value = codeText.value
    document.body.appendChild(textArea)
    textArea.select()
    document.execCommand('copy')
    document.body.removeChild(textArea)
    copied.value = true
    setTimeout(() => {
      copied.value = false
    }, 2000)
  }
}

const toggleWrap = () => {
  isWrapped.value = !isWrapped.value
  if (codeElement.value) {
    codeElement.value.style.whiteSpace = isWrapped.value ? 'pre-wrap' : 'pre'
  }
}

const applyHighlighting = () => {
  if (codeBlock.value) {
    try {
      hljs.highlightElement(codeBlock.value)
    } catch (error) {
      console.warn('Syntax highlighting failed:', error)
    }
  }
}

// Lifecycle
onMounted(async () => {
  await nextTick()
  
  // Apply syntax highlighting
  applyHighlighting()
  
  // Apply max height if specified
  if (props.maxHeight !== 'none' && codeElement.value) {
    codeElement.value.style.maxHeight = props.maxHeight
    codeElement.value.style.overflowY = 'auto'
  }
  
  // Apply line number padding if enabled
  if (props.showLineNumbers && codeElement.value) {
    const lineNumberWidth = String(lineCount.value).length * 8 + 32
    codeElement.value.style.paddingLeft = `${lineNumberWidth}px`
  }
})

// Watch for code changes to re-apply highlighting
watch(() => props.code, () => {
  nextTick(() => {
    applyHighlighting()
  })
})
</script>

<style scoped>
.code-block-container {
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  border: 1px solid #e5e7eb;
  background: #1f2937;
  position: relative;
}

.code-block-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
  background: #374151;
  border-bottom: 1px solid #4b5563;
}

.language-badge {
  display: inline-flex;
  align-items: center;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
  background: rgba(249, 115, 22, 0.2);
  color: #fb923c;
  border: 1px solid rgba(249, 115, 22, 0.3);
}

.code-title {
  font-size: 14px;
  font-weight: 500;
  color: #d1d5db;
}

.copy-button {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 12px;
  font-size: 12px;
  background: #4b5563;
  color: #d1d5db;
  border: none;
  border-radius: 6px;
  transition: all 0.2s ease;
  cursor: pointer;
}

.copy-button:hover {
  background: #6b7280;
  color: #ffffff;
  transform: translateY(-1px);
}

.copy-button.copied {
  background: #059669;
  color: #ffffff;
}

.code-block-content {
  position: relative;
}

.code-content {
  padding: 0;
  margin: 0;
  overflow-x: auto;
  background: #1f2937;
  white-space: pre;
}

.code-content code {
  display: block;
  padding: 24px;
  font-family: 'Fira Code', 'Monaco', 'Cascadia Code', 'Roboto Mono', monospace;
  font-size: 14px;
  line-height: 1.6;
  color: #f3f4f6;
  background: transparent;
  tab-size: 2;
  white-space: pre;
}

.line-numbers {
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  padding: 24px 16px;
  background: rgba(55, 65, 81, 0.5);
  border-right: 1px solid #4b5563;
  color: #9ca3af;
  font-size: 14px;
  font-family: 'Fira Code', 'Monaco', 'Cascadia Code', 'Roboto Mono', monospace;
  user-select: none;
  pointer-events: none;
  white-space: pre;
}

.line-numbers div {
  line-height: 1.6;
}

.code-block-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 16px;
  background: #374151;
  border-top: 1px solid #4b5563;
  font-size: 12px;
  color: #9ca3af;
}

.wrap-button {
  padding: 4px 8px;
  background: transparent;
  color: #9ca3af;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: colors 0.2s ease;
}

.wrap-button:hover {
  background: #4b5563;
  color: #d1d5db;
}

.wrap-button.active {
  color: #fb923c;
}

/* Custom scrollbar */
.code-content::-webkit-scrollbar {
  height: 8px;
  width: 8px;
}

.code-content::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.1);
}

.code-content::-webkit-scrollbar-thumb {
  background: rgba(249, 115, 22, 0.3);
  border-radius: 4px;
}

.code-content::-webkit-scrollbar-thumb:hover {
  background: rgba(249, 115, 22, 0.5);
}

/* Responsive design */
@media (max-width: 768px) {
  .code-content code {
    padding: 16px;
    font-size: 12px;
  }
  
  .code-block-header {
    padding: 8px 12px;
  }
  
  .code-block-footer {
    padding: 8px 12px;
  }
}
</style>

<style>
/* Global styles for highlight.js */
.hljs {
  background: transparent !important;
  color: #f3f4f6 !important;
}
</style>