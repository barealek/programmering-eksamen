<script>
   export let name, code, size, url
   import {formatFileSize} from "$lib"
</script>

<li class="col-span-1 flex rounded-md shadow-xs">
<div class="flex w-16 shrink-0 items-center justify-center rounded-l-md bg-pink-600 text-sm font-medium text-white">GA</div>
<div class="flex flex-1 items-center justify-between truncate rounded-r-md border-t border-r border-b border-gray-200 bg-white">
   <div class="flex-1 truncate px-4 py-2 text-sm">
      <button on:click={() => {navigator.clipboard.writeText(url); alert("Link kopieret for "+name)}} class="font-medium cursor-pointer text-gray-900 hover:text-gray-600">{name}</button>
      <p class="text-gray-500">{formatFileSize(size)}</p>
   </div>
   <div class="shrink-0 pr-2">
      <button type="button" class="inline-flex size-8 items-center justify-center rounded-full bg-white text-gray-400 hover:text-gray-500 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:outline-hidden">
      <span class="sr-only">Open options</span>
      <svg class="size-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" data-slot="icon">
         <path d="M10 3a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM10 8.5a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM11.5 15.5a1.5 1.5 0 1 0-3 0 1.5 1.5 0 0 0 3 0Z" />
      </svg>
      </button>
   </div>
</div>
</li>
