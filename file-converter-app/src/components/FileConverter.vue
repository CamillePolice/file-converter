<script setup>
import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';
import FileDropZone from '@/components/FileDropZone.vue';

const route = useRoute();
const router = useRouter();
const toast = useToast();

const type = computed(() => route.params.type);

const pageTitle = computed(() => {
    const titles = {
        'pdf-to-word': 'Convert PDF to Word',
        'word-to-pdf': 'Convert Word to PDF',
        'image-to-pdf': 'Convert Image to PDF',
        'pdf-merge': 'Merge PDF Files',
        'ppt-to-pdf': 'Convert PowerPoint to PDF',
        'excel-to-pdf': 'Convert Excel to PDF'
    };
    return titles[type.value] || 'File Converter';
});

const acceptedFormats = computed(() => {
    const formats = {
        'pdf-to-word': '.pdf',
        'word-to-pdf': '.doc,.docx',
        'image-to-pdf': '.jpg,.jpeg,.png',
        'pdf-merge': '.pdf',
        'ppt-to-pdf': '.ppt,.pptx',
        'excel-to-pdf': '.xls,.xlsx'
    };
    return formats[type.value] || '';
});

const supportedFormats = computed(() => {
    const formats = {
        'pdf-to-word': 'Supports: PDF files',
        'word-to-pdf': 'Supports: DOC, DOCX files',
        'image-to-pdf': 'Supports: JPG, PNG files',
        'pdf-merge': 'Supports: PDF files',
        'ppt-to-pdf': 'Supports: PPT, PPTX files',
        'excel-to-pdf': 'Supports: XLS, XLSX files'
    };
    return formats[type.value] || '';
});

const handleConversionStart = () => {
    toast.add({
        severity: 'info',
        summary: 'Converting',
        detail: 'Your file conversion has started...',
        life: 3000
    });
};

const handleConversionComplete = (result) => {
    if (result.success) {
        toast.add({
            severity: 'success',
            summary: 'Success',
            detail: 'Your file has been converted successfully!',
            life: 3000
        });
    } else {
        toast.add({
            severity: 'error',
            summary: 'Error',
            detail: result.error || 'An error occurred during conversion',
            life: 5000
        });
    }
};
</script>

<template>
    <div class="min-h-screen bg-gray-50">
        <header class="bg-white shadow">
            <div class="max-w-7xl mx-auto px-4 py-6">
                <div class="flex items-center justify-between">
                    <div class="flex items-center space-x-4">
                        <Button icon="pi pi-arrow-left" text @click="router.push('/')" class="!p-2" />
                        <h1 class="text-2xl font-bold text-gray-900">{{ pageTitle }}</h1>
                    </div>
                    <div class="text-sm text-gray-500">
                        {{ supportedFormats }}
                    </div>
                </div>
            </div>
        </header>

        <main class="max-w-4xl mx-auto py-12 px-4">
            <FileDropZone :accepted-formats="acceptedFormats" :conversion-type="type"
                @conversion-start="handleConversionStart" @conversion-complete="handleConversionComplete" />
        </main>
    </div>
</template>