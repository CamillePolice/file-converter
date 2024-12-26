<script setup>
import { ref } from 'vue';
import { useToast } from 'primevue/usetoast';

const toast = useToast();

const file1 = ref(null);
const file2 = ref(null);
const isDraggingFirst = ref(false);
const isDraggingSecond = ref(false);

const handleDragOver = (dropzone) => {
    if (dropzone === 'first') {
        isDraggingFirst.value = true;
    } else {
        isDraggingSecond.value = true;
    }
};

const handleDragLeave = (dropzone) => {
    if (dropzone === 'first') {
        isDraggingFirst.value = false;
    } else {
        isDraggingSecond.value = false;
    }
};

const handleDrop = (event, dropzone) => {
    const droppedFile = event.dataTransfer.files[0];

    if (!droppedFile) return;

    if (dropzone === 'first') {
        isDraggingFirst.value = false;
        file1.value = droppedFile;
    } else {
        isDraggingSecond.value = false;
        file2.value = droppedFile;
    }

    // Optional: Show success message
    toast.add({
        severity: 'success',
        summary: 'File Added',
        detail: `${droppedFile.name} has been added`,
        life: 3000
    });
};

const removeFile = (dropzone) => {
    if (dropzone === 'first') {
        file1.value = null;
    } else {
        file2.value = null;
    }
};

const formatFileSize = (bytes) => {
    if (bytes === 0) return '0 Bytes';

    const k = 1024;
    const sizes = ['Bytes', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));

    return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`;
};
</script>

<template>
    <div class="w-full max-w-4xl mx-auto p-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- First Drop Zone -->
            <div class="flex flex-col space-y-4">
                <h3 class="text-lg font-medium">File 1</h3>
                <div :class="[
                    'border-2 border-dashed rounded-lg p-8',
                    'transition-colors duration-200 ease-in-out',
                    'flex flex-col items-center justify-center',
                    'min-h-48 space-y-4',
                    isDraggingFirst ? 'border-primary bg-primary-50' : 'border-gray-300',
                    file1 ? 'bg-gray-50' : 'bg-white'
                ]" @dragover.prevent="handleDragOver('first')" @dragleave.prevent="handleDragLeave('first')"
                    @drop.prevent="handleDrop($event, 'first')">
                    <i class="pi pi-upload text-4xl" :class="isDraggingFirst ? 'text-primary' : 'text-gray-400'" />

                    <template v-if="!file1">
                        <p class="text-center text-gray-600">
                            Drag and drop first file here
                            <br>
                            <span class="text-sm text-gray-400">or click to browse</span>
                        </p>
                    </template>

                    <template v-else>
                        <div class="w-full">
                            <div class="flex items-center justify-between p-3 bg-white rounded-lg shadow-sm">
                                <div class="flex items-center space-x-3 truncate">
                                    <i class="pi pi-file text-xl text-gray-500" />
                                    <div class="flex-1 min-w-0">
                                        <p class="text-sm font-medium text-gray-900 truncate">
                                            {{ file1.name }}
                                        </p>
                                        <p class="text-sm text-gray-500">
                                            {{ formatFileSize(file1.size) }}
                                        </p>
                                    </div>
                                </div>
                                <Button icon="pi pi-times" severity="danger" text @click="removeFile('first')"
                                    aria-label="Remove file" />
                            </div>
                        </div>
                    </template>
                </div>
            </div>

            <!-- Second Drop Zone -->
            <div class="flex flex-col space-y-4">
                <h3 class="text-lg font-medium">File 2</h3>
                <div :class="[
                    'border-2 border-dashed rounded-lg p-8',
                    'transition-colors duration-200 ease-in-out',
                    'flex flex-col items-center justify-center',
                    'min-h-48 space-y-4',
                    isDraggingSecond ? 'border-primary bg-primary-50' : 'border-gray-300',
                    file2 ? 'bg-gray-50' : 'bg-white'
                ]" @dragover.prevent="handleDragOver('second')" @dragleave.prevent="handleDragLeave('second')"
                    @drop.prevent="handleDrop($event, 'second')">
                    <i class="pi pi-upload text-4xl" :class="isDraggingSecond ? 'text-primary' : 'text-gray-400'" />

                    <template v-if="!file2">
                        <p class="text-center text-gray-600">
                            Drag and drop second file here
                            <br>
                            <span class="text-sm text-gray-400">or click to browse</span>
                        </p>
                    </template>

                    <template v-else>
                        <div class="w-full">
                            <div class="flex items-center justify-between p-3 bg-white rounded-lg shadow-sm">
                                <div class="flex items-center space-x-3 truncate">
                                    <i class="pi pi-file text-xl text-gray-500" />
                                    <div class="flex-1 min-w-0">
                                        <p class="text-sm font-medium text-gray-900 truncate">
                                            {{ file2.name }}
                                        </p>
                                        <p class="text-sm text-gray-500">
                                            {{ formatFileSize(file2.size) }}
                                        </p>
                                    </div>
                                </div>
                                <Button icon="pi pi-times" severity="danger" text @click="removeFile('second')"
                                    aria-label="Remove file" />
                            </div>
                        </div>
                    </template>
                </div>
            </div>
        </div>

        <!-- Optional Toast for errors -->
        <Toast />
    </div>
</template>
