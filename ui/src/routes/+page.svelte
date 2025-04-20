<script>
  import { onMount } from 'svelte';

  import TidligereUploads from "$lib/TidligereUploads.svelte"

  import {formatDate, formatFileSize} from "$lib/index.js"
  import Upload from '$lib/Upload.svelte';

  let fileInput;
  let file = null;
  let isDragging = false;
  let isUploading = false;
  let uploadProgress = 0;
  let response = null;
  let error = null;

  // Track uploaded files
  let uploadedFiles = [];
  // Track which files have visible deletion codes
  let visibleCodes = {};

  // Load previously uploaded files from localStorage on mount
  onMount(() => {
    const storedFiles = localStorage.getItem('uploadedFiles');
    console.log(storedFiles)
    if (storedFiles) {
      uploadedFiles = JSON.parse(storedFiles);
    }
  });

  // Save uploaded files to localStorage
  function saveUploadedFiles() {
    localStorage.setItem('uploadedFiles', JSON.stringify(uploadedFiles));
  }


  // Add a file to the uploaded files list
  function addUploadedFile(fileData) {
   console.log("addUploadedFile")
    const newFile = {
      name: file.name,
      size: file.size,
      type: file.type,
      url: fileData.url,
      code: fileData.code,
      date: new Date().toISOString()
    };

    uploadedFiles = [newFile, ...uploadedFiles];
    saveUploadedFiles();
  }

  // Remove a file from the uploaded files list
  function removeUploadedFile(index) {
    uploadedFiles = uploadedFiles.filter((_, i) => i !== index);
    saveUploadedFiles();
  }

  // Toggle visibility of deletion code
  function toggleCodeVisibility(index) {
    visibleCodes = {
      ...visibleCodes,
      [index]: !visibleCodes[index]
    };
  }

  // Handle file selection
  function handleFileSelect(event) {
    const selectedFile = event.target.files[0];
    if (selectedFile) {
      file = selectedFile;
    }
  }

  // Handle drag events
  function handleDragEnter(event) {
    event.preventDefault();
    isDragging = true;
  }

  function handleDragLeave(event) {
    event.preventDefault();
    isDragging = false;
  }

  function handleDragOver(event) {
    event.preventDefault();
  }

  function handleDrop(event) {
    event.preventDefault();
    isDragging = false;

    const droppedFile = event.dataTransfer.files[0];
    if (droppedFile) {
      file = droppedFile;
    }
  }

  // Handle file upload using fetch with raw binary data
  async function uploadFile() {
    if (!file) return;

    try {
      isUploading = true;
      error = null;
      response = null;
      uploadProgress = 0;

      // Set up progress simulation
      const progressInterval = setInterval(() => {
        if (uploadProgress < 90) {
          uploadProgress += 5;
        }
      }, 500);

      try {
        // Read the file as binary data instead of using FormData
        const fileContent = await file.arrayBuffer();

        const res = await fetch('/api/file', {
          method: 'POST',
          headers: {
            'Content-Type': file.type,
            'Content-Disposition': `attachment; filename="${file.name}"`
          },
          body: fileContent
        });

        // Clear the progress interval
        clearInterval(progressInterval);

        if (!res.ok) {
          throw new Error(`Server svarede med ${res.status}: ${res.statusText}`);
        }

        // Set progress to 100% when complete
        uploadProgress = 100;

        // Parse the response
        response = await res.json();

        // Store the uploaded file information
        if (response && response.url && response.code) {
          addUploadedFile(response);
        }
      } catch (fetchError) {
        clearInterval(progressInterval);
        throw fetchError;
      }
    } catch (err) {
      error = `Upload fejl: ${err.message}`;
    } finally {
      isUploading = false;
    }
  }

  // Reset the upload
  function resetUpload() {
    file = null;
    response = null;
    error = null;
    uploadProgress = 0;
  }


</script>

<div class="bg-white px-6 py-24 sm:py-32 lg:px-8">
  <div class="mx-auto max-w-2xl text-center">
    <p class="text-base/7 font-semibold text-indigo-600">Sikker, privat fildeling</p>
    <h2 class="mt-2 text-5xl font-semibold tracking-tight text-gray-900 sm:text-7xl">Upload fil</h2>
  </div>

  <div class="mt-12 justify-center items-center flex">
    <div class="container max-w-xl">
      <input
        type="file"
        bind:this={fileInput}
        on:change={handleFileSelect}
        class="hidden"
      />

      {#if !file && !response}
        <!-- Upload area -->
        <button
          type="button"
          class="relative block w-full rounded-lg border-2 border-dashed {isDragging ? 'border-indigo-500 bg-indigo-50' : 'border-gray-300'} p-12 text-center hover:border-gray-400 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:outline-hidden min-h-64"
          on:click={() => fileInput.click()}
          on:dragenter={handleDragEnter}
          on:dragleave={handleDragLeave}
          on:dragover={handleDragOver}
          on:drop={handleDrop}
        >
          <svg class="mx-auto size-12 text-gray-400" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 13v8"/><path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"/><path d="m8 17 4-4 4 4"/></svg>
          <span class="mt-2 block text-sm font-semibold text-gray-900">Smid fil her eller klik</span>
        </button>
      {:else if file && !isUploading && !response}
        <!-- File selected, ready to upload -->
        <div class="p-6 border-2 border-dashed border-gray-300 rounded-lg">
          <div class="flex items-center justify-between">
            <div>
              <p class="font-medium text-gray-900">{file.name}</p>
              <p class="text-sm text-gray-500">{(file.size / 1024 / 1024).toFixed(2)} MB</p>
            </div>
            <div class="flex gap-2">
              <button
                type="button"
                class="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700"
                on:click={uploadFile}
              >
                Upload
              </button>
              <button
                type="button"
                class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
                on:click={resetUpload}
              >
                Annuller
              </button>
            </div>
          </div>
        </div>
      {:else if isUploading}
        <!-- Upload in progress -->
        <div class="p-6 border-2 border-dashed border-gray-300 rounded-lg">
          <p class="font-medium text-gray-900 mb-2">Uploader {file.name}</p>
          <div class="w-full bg-gray-200 rounded-full h-2.5">
            <div class="bg-indigo-600 h-2.5 rounded-full" style="width: {uploadProgress}%"></div>
          </div>
          <p class="text-sm text-gray-500 mt-2">{uploadProgress}% fuldført</p>
        </div>
      {:else if response}
        <!-- Upload complete -->
        <div class="p-6 border-2 border-solid border-green-300 bg-green-50 rounded-lg">
          <div class="text-center mb-4">
            <svg class="mx-auto size-12 text-green-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
              <polyline points="22 4 12 14.01 9 11.01"></polyline>
            </svg>
            <h3 class="mt-2 text-lg font-medium text-gray-900">Upload fuldført!</h3>
          </div>

          {#if response.url}
            <div class="bg-white p-4 rounded-md border border-gray-200 mb-4">
              <p class="text-sm text-gray-500 mb-1">Del dette link:</p>
              <p class="font-mono bg-gray-100 p-2 rounded text-sm break-all">{response.url}</p>
            </div>
          {/if}

          <button
            type="button"
            class="w-full px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700"
            on:click={resetUpload}
          >
            Upload en ny fil
          </button>
        </div>
      {/if}

      {#if error}
        <div class="mt-4 p-4 bg-red-50 border border-red-200 rounded-md text-red-700">
          <p>{error}</p>
        </div>
      {/if}

      <!-- Previously uploaded files section -->
      {uploadedFiles.length}
    </div>
  </div>
      {#if uploadedFiles.length > 0}
         <TidligereUploads>
            {#each uploadedFiles as uploadedFile, i}
               <Upload name={uploadedFile.name} size={uploadedFile.size} url={uploadedFile.url} />
            {/each}
         </TidligereUploads>
      {/if}
</div>
