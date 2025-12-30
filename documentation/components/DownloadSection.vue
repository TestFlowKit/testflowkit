<template>
    <div class="space-y-6">
        <div class="grid md:grid-cols-3 gap-4">
            <button v-for="platform in platforms" :key="platform.id" @click="selectPlatform(platform.id)" :class="[
                'p-4 rounded-lg border-2 transition-all duration-200 hover:shadow-md',
                selectedPlatform === platform.id
                    ? 'border-blue-500 bg-blue-50'
                    : 'border-gray-200 hover:border-gray-300'
            ]">
                <div class="text-center">
                    <div class="text-3xl mb-2">{{ platform.icon }}</div>
                    <div class="font-semibold text-gray-800">{{ platform.name }}</div>
                    <div class="text-sm text-gray-600">{{ platform.description }}</div>
                </div>
            </button>
        </div>

        <div v-if="selectedPlatform" class="space-y-4">
            <h3 class="text-lg font-semibold text-gray-800">Choose Architecture:</h3>
            <div class="grid md:grid-cols-2 gap-4">
                <button v-for="arch in getArchitecturesForPlatform(selectedPlatform)" :key="arch.id"
                    @click="selectArchitecture(arch.id)" :class="[
                        'p-4 rounded-lg border-2 transition-all duration-200 hover:shadow-md',
                        selectedArchitecture === arch.id
                            ? 'border-green-500 bg-green-50'
                            : 'border-gray-200 hover:border-gray-300'
                    ]">
                    <div class="text-center">
                        <div class="font-semibold text-gray-800">{{ arch.name }}</div>
                        <div class="text-sm text-gray-600">{{ arch.description }}</div>
                    </div>
                </button>
            </div>
        </div>

        <div v-if="selectedPlatform && selectedArchitecture" class="text-center">
            <div class="bg-green-50 border border-green-200 rounded-lg p-4 mb-4">
                <div class="flex items-center justify-center space-x-2 mb-2">
                    <div class="text-2xl">{{ getPlatformIcon(selectedPlatform) }}</div>
                    <div class="font-semibold text-green-800">
                        {{ getPlatformName(selectedPlatform) }} ({{ getArchitectureName(selectedArchitecture) }})
                    </div>
                </div>
                <p class="text-sm text-green-700 mb-2">
                    Latest version: {{ latestVersion || 'Loading...' }}
                </p>
                <div class="text-xs text-green-600">
                    <div class="font-mono">SHA256: {{ getSelectedBinaryInfo().sha256 }}</div>
                </div>
            </div>

            <button @click="downloadBinary" :disabled="isDownloading"
                class="bg-blue-600 text-white px-8 py-3 rounded-lg font-semibold hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors duration-200">
                <div class="flex items-center space-x-2">
                    <svg v-if="isDownloading" class="animate-spin h-5 w-5" fill="none" viewBox="0 0 24 24">
                        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4">
                        </circle>
                        <path class="opacity-75" fill="currentColor"
                            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z">
                        </path>
                    </svg>
                    <svg v-else class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z">
                        </path>
                    </svg>
                    <span>{{ isDownloading ? 'Downloading...' : 'Download TestFlowKit' }}</span>
                </div>
            </button>
        </div>

        <div class="text-center">
            <p class="text-sm text-gray-600 mb-2">Or download from:</p>
            <a :href="GITHUB_RELEASES_URL" target="_blank"
                class="text-blue-600 hover:underline font-medium">
                GitHub Releases
            </a>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { GITHUB_RELEASES_URL, GITHUB_RELEASES_LATEST_API_URL, GITHUB_RELEASES_LATEST_DOWNLOAD_BASE } from '~/constants/links';

const selectedPlatform = ref('');
const selectedArchitecture = ref('');
const latestVersion = ref('');
const isDownloading = ref(false);

const platforms = [
    {
        id: 'windows',
        name: 'Windows',
        icon: '🪟',
        description: 'Windows 10/11, Server'
    },
    {
        id: 'darwin',
        name: 'macOS',
        icon: '🍎',
        description: 'macOS 10.15+'
    },
    {
        id: 'linux',
        name: 'Linux',
        icon: '🐧',
        description: 'Ubuntu, CentOS, etc.'
    }
];

const binaryInfo = {
    'darwin-amd64': {
        name: 'Intel Mac',
        description: 'Intel processors',
    },
    'darwin-arm64': {
        name: 'Apple Silicon',
        description: 'M1/M2/M3 chips',
    },
    'linux-amd64': {
        name: '64-bit (x64)',
        description: 'Most common',
    },
    'linux-arm64': {
        name: 'ARM 64-bit',
        description: 'ARM processors',
    },
    'windows-amd64': {
        name: '64-bit (x64)',
        description: 'Most common',
    }
};

const architectures = {
    windows: [
        { id: 'amd64', name: '64-bit (x64)', description: 'Most common' }
    ],
    darwin: [
        { id: 'amd64', name: 'Intel Mac', description: 'Intel processors' },
        { id: 'arm64', name: 'Apple Silicon', description: 'M1/M2/M3 chips' }
    ],
    linux: [
        { id: 'amd64', name: '64-bit (x64)', description: 'Most common' },
        { id: 'arm64', name: 'ARM 64-bit', description: 'ARM processors' }
    ]
};

const getArchitecturesForPlatform = (platform: string) => {
    return architectures[platform as keyof typeof architectures] || [];
};

const getPlatformIcon = (platform: string) => {
    const p = platforms.find(p => p.id === platform);
    return p?.icon || '💻';
};

const getPlatformName = (platform: string) => {
    const p = platforms.find(p => p.id === platform);
    return p?.name || platform;
};

const getArchitectureName = (arch: string) => {
    const archMap: { [key: string]: string } = {
        'amd64': '64-bit',
        'arm64': 'ARM 64-bit'
    };
    return archMap[arch] || arch;
};

const getSelectedBinaryInfo = () => {
    const key = `${selectedPlatform.value}-${selectedArchitecture.value}`;
    return binaryInfo[key as keyof typeof binaryInfo] || { sha256: 'Unknown' };
};

const selectPlatform = (platform: string) => {
    selectedPlatform.value = platform;
    selectedArchitecture.value = ''; // Reset architecture selection
};

const selectArchitecture = (arch: string) => {
    selectedArchitecture.value = arch;
};

const fetchLatestVersion = async () => {
    try {
        const response = await fetch(GITHUB_RELEASES_LATEST_API_URL);
        const data = await response.json();
        latestVersion.value = data.tag_name || 'v1.0.0';
    } catch (error) {
        console.error('Failed to fetch latest version:', error);
        latestVersion.value = 'v1.0.0';
    }
};

const downloadBinary = async () => {
    if (!selectedPlatform.value || !selectedArchitecture.value) return;

    isDownloading.value = true;

    try {
        const filename = `tkit-${selectedPlatform.value}-${selectedArchitecture.value}.zip`;
        const downloadUrl = `${GITHUB_RELEASES_LATEST_DOWNLOAD_BASE}/${filename}`;

        // Create a temporary link element to trigger download
        const link = document.createElement('a');
        link.href = downloadUrl;
        link.download = filename;
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);

        // Add a small delay to show the downloading state
        setTimeout(() => {
            isDownloading.value = false;
        }, 1000);

    } catch (error) {
        console.error('Download failed:', error);
        isDownloading.value = false;

        const filename = `tkit-${selectedPlatform.value}-${selectedArchitecture.value}`;
        const downloadUrl = `${GITHUB_RELEASES_LATEST_DOWNLOAD_BASE}/${filename}`;
        window.open(downloadUrl, '_blank');
    }
};

onMounted(() => {
    fetchLatestVersion();
});
</script>
