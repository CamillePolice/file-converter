<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const navigateToConverter = (type) => {
    router.push({
        name: 'converter',
        params: { type }
    });
};
</script>

<template>
    <div class="min-h-screen bg-gray-50">
        <header class="bg-gradient-to-b from-gray-100 to-white border-b">
            <div class="max-w-7xl mx-auto py-4 px-6">
                <h1 class="text-2xl font-semibold text-gray-900">File Converter</h1>
            </div>
        </header>

        <main class="max-w-7xl mx-auto py-8 px-6">
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                <!-- PDF to Word -->
                <div class="bg-white rounded-lg shadow-sm hover:shadow-md transition-all duration-200 cursor-pointer p-6"
                    @click="navigateToConverter('pdf-to-word')">
                    <div class="w-8 h-8 bg-blue-100 rounded mb-4"></div>
                    <h3 class="text-lg font-medium mb-2">PDF to Word</h3>
                    <p class="text-gray-500 mb-4">Convert PDF files to editable Word documents</p>
                    <p class="text-sm text-gray-400">.pdf to .docx</p>
                </div>

                <!-- Word to PDF -->
                <div class="bg-white rounded-lg shadow hover:shadow-lg transition-shadow cursor-pointer"
                    @click="navigateToConverter('word-to-pdf')">
                    <div class="p-6">
                        <div class="w-12 h-12 bg-red-100 rounded-lg flex items-center justify-center mb-4">
                            <i class="pi pi-file-word text-red-600 text-2xl"></i>
                        </div>
                        <h3 class="text-lg font-semibold mb-2">Word to PDF</h3>
                        <p class="text-gray-600">Convert Word documents to PDF format</p>
                    </div>
                    <div class="px-6 py-4 bg-gray-50 rounded-b-lg">
                        <div class="flex items-center text-sm text-gray-500">
                            <i class="pi pi-arrow-right mr-2"></i>
                            <span>.docx to .pdf</span>
                        </div>
                    </div>
                </div>

                <!-- Image to PDF -->
                <div class="bg-white rounded-lg shadow hover:shadow-lg transition-shadow cursor-pointer"
                    @click="navigateToConverter('image-to-pdf')">
                    <div class="p-6">
                        <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center mb-4">
                            <i class="pi pi-image text-green-600 text-2xl"></i>
                        </div>
                        <h3 class="text-lg font-semibold mb-2">Image to PDF</h3>
                        <p class="text-gray-600">Convert images to PDF documents</p>
                    </div>
                    <div class="px-6 py-4 bg-gray-50 rounded-b-lg">
                        <div class="flex items-center text-sm text-gray-500">
                            <i class="pi pi-arrow-right mr-2"></i>
                            <span>.jpg/.png to .pdf</span>
                        </div>
                    </div>
                </div>

                <!-- PDF Merge -->
                <div class="bg-white rounded-lg shadow hover:shadow-lg transition-shadow cursor-pointer"
                    @click="navigateToConverter('pdf-merge')">
                    <div class="p-6">
                        <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center mb-4">
                            <i class="pi pi-copy text-purple-600 text-2xl"></i>
                        </div>
                        <h3 class="text-lg font-semibold mb-2">PDF Merge</h3>
                        <p class="text-gray-600">Combine multiple PDF files into one</p>
                    </div>
                    <div class="px-6 py-4 bg-gray-50 rounded-b-lg">
                        <div class="flex items-center text-sm text-gray-500">
                            <i class="pi pi-arrow-right mr-2"></i>
                            <span>Multiple .pdf to single .pdf</span>
                        </div>
                    </div>
                </div>

                <!-- PowerPoint to PDF -->
                <div class="bg-white rounded-lg shadow hover:shadow-lg transition-shadow cursor-pointer"
                    @click="navigateToConverter('ppt-to-pdf')">
                    <div class="p-6">
                        <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center mb-4">
                            <i class="pi pi-file-excel text-orange-600 text-2xl"></i>
                        </div>
                        <h3 class="text-lg font-semibold mb-2">PowerPoint to PDF</h3>
                        <p class="text-gray-600">Convert presentations to PDF format</p>
                    </div>
                    <div class="px-6 py-4 bg-gray-50 rounded-b-lg">
                        <div class="flex items-center text-sm text-gray-500">
                            <i class="pi pi-arrow-right mr-2"></i>
                            <span>.ppt/.pptx to .pdf</span>
                        </div>
                    </div>
                </div>

                <!-- Excel to PDF -->
                <div class="bg-white rounded-lg shadow hover:shadow-lg transition-shadow cursor-pointer"
                    @click="navigateToConverter('excel-to-pdf')">
                    <div class="p-6">
                        <div class="w-12 h-12 bg-yellow-100 rounded-lg flex items-center justify-center mb-4">
                            <i class="pi pi-file-excel text-yellow-600 text-2xl"></i>
                        </div>
                        <h3 class="text-lg font-semibold mb-2">Excel to PDF</h3>
                        <p class="text-gray-600">Convert spreadsheets to PDF format</p>
                    </div>
                    <div class="px-6 py-4 bg-gray-50 rounded-b-lg">
                        <div class="flex items-center text-sm text-gray-500">
                            <i class="pi pi-arrow-right mr-2"></i>
                            <span>.xls/.xlsx to .pdf</span>
                        </div>
                    </div>
                </div>
            </div>
        </main>
    </div>
</template>