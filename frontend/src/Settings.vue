<script setup>
import Navbar from './components/Navbar.vue'

onclick = "document.documentElement.setAttribute('data-theme', '$1'); localStorage.setItem('theme', '$1');"

</script>

<template>
     <Navbar />
     <center>
          <br>

          <div style="width: 80%; margin-left: 120px"
               class="rounded-box grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
             
               <div onclick="document.documentElement.setAttribute('data-theme', 'nordic'); localStorage.setItem('theme', 'nordic'); ch()"
                    data-set-theme="nordic"
                    class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
                    data-act-class="!outline-base-content">
                    <div data-theme="nord" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                         <div class="grid grid-cols-5 grid-rows-3">
                              <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                              <div class="bg-base-300 col-start-1 row-start-3"></div>
                              <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                                   <div class="font-bold">Nordic</div>
                                   <div class="flex flex-wrap gap-1">
                                        <div
                                             class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                             <div class="text-primary-content text-sm font-bold">A</div>
                                        </div>
                                        <div
                                             class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                             <div class="text-secondary-content text-sm font-bold">A</div>
                                        </div>
                                        <div
                                             class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                             <div class="text-accent-content text-sm font-bold">A</div>
                                        </div>
                                        <div
                                             class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                             <div class="text-neutral-content text-sm font-bold">A</div>
                                        </div>
                              </div>
                         </div>
                    </div>
               </div>
          </div>
          <div onclick="document.documentElement.setAttribute('data-theme', 'business'); localStorage.setItem('theme', 'business'); ch()"
               data-set-theme="business"
               class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
               data-act-class="!outline-base-content">
               <div data-theme="business" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                    <div class="grid grid-cols-5 grid-rows-3">
                         <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                         <div class="bg-base-300 col-start-1 row-start-3"></div>
                         <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                              <div class="font-bold">Business</div>
                              <div class="flex flex-wrap gap-1">
                                   <div
                                        class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-primary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-secondary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-accent-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-neutral-content text-sm font-bold">A</div>
                                   </div>
                              </div>
                         </div>
                    </div>
               </div>
          </div>
          <div onclick="document.documentElement.setAttribute('data-theme', 'forest'); localStorage.setItem('theme', 'forest'); ch()"
               data-set-theme="forest"
               class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
               data-act-class="!outline-base-content">
               <div data-theme="forest" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                    <div class="grid grid-cols-5 grid-rows-3">
                         <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                         <div class="bg-base-300 col-start-1 row-start-3"></div>
                         <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                              <div class="font-bold">Forest</div>
                              <div class="flex flex-wrap gap-1">
                                   <div
                                        class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-primary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-secondary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-accent-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-neutral-content text-sm font-bold">A</div>
                                   </div>
                              </div>
                         </div>
                    </div>
               </div>
          </div>
          <div onclick="document.documentElement.setAttribute('data-theme', 'dracula'); localStorage.setItem('theme', 'dracula'); ch()"
               data-set-theme="dracula"
               class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
               data-act-class="!outline-base-content">
               <div data-theme="dracula" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                    <div class="grid grid-cols-5 grid-rows-3">
                         <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                         <div class="bg-base-300 col-start-1 row-start-3"></div>
                         <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                              <div class="font-bold">Dracula</div>
                              <div class="flex flex-wrap gap-1">
                                   <div
                                        class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-primary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-secondary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-accent-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-neutral-content text-sm font-bold">A</div>
                                   </div>
                              </div>
                         </div>
                    </div>
               </div>
          </div>
          <div onclick="document.documentElement.setAttribute('data-theme', 'luxury'); localStorage.setItem('theme', 'luxury'); ch()"
               data-set-theme="dracula"
               class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
               data-act-class="!outline-base-content">
               <div data-theme="luxury" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                    <div class="grid grid-cols-5 grid-rows-3">
                         <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                         <div class="bg-base-300 col-start-1 row-start-3"></div>
                         <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                              <div class="font-bold">Luxury</div>
                              <div class="flex flex-wrap gap-1">
                                   <div
                                        class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-primary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-secondary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-accent-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-neutral-content text-sm font-bold">A</div>
                                   </div>
                              </div>
                         </div>
                    </div>
               </div>
          </div>
          <div onclick="document.documentElement.setAttribute('data-theme', 'night'); localStorage.setItem('theme', 'night'); ch()"
               data-set-theme="dracula"
               class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
               data-act-class="!outline-base-content">
               <div data-theme="night" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                    <div class="grid grid-cols-5 grid-rows-3">
                         <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                         <div class="bg-base-300 col-start-1 row-start-3"></div>
                         <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                              <div class="font-bold">Night</div>
                              <div class="flex flex-wrap gap-1">
                                   <div
                                        class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-primary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-secondary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-accent-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-neutral-content text-sm font-bold">A</div>
                                   </div>
                              </div>
                         </div>
                    </div>
               </div>
          </div>
          <div onclick="document.documentElement.setAttribute('data-theme', 'coffee'); localStorage.setItem('theme', 'coffee'); ch()"
               data-set-theme="coffee"
               class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
               data-act-class="!outline-base-content">
               <div data-theme="coffee" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                    <div class="grid grid-cols-5 grid-rows-3">
                         <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                         <div class="bg-base-300 col-start-1 row-start-3"></div>
                         <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                              <div class="font-bold">Coffee</div>
                              <div class="flex flex-wrap gap-1">
                                   <div
                                        class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-primary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-secondary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-accent-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-neutral-content text-sm font-bold">A</div>
                                   </div>
                              </div>
                         </div>
                    </div>
               </div>
          </div>
          <div onclick="document.documentElement.setAttribute('data-theme', 'aqua'); localStorage.setItem('theme', 'aqua'); ch()"
               data-set-theme="dracula"
               class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
               data-act-class="!outline-base-content">
               <div data-theme="aqua" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                    <div class="grid grid-cols-5 grid-rows-3">
                         <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                         <div class="bg-base-300 col-start-1 row-start-3"></div>
                         <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                              <div class="font-bold">Aqua</div>
                              <div class="flex flex-wrap gap-1">
                                   <div
                                        class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-primary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-secondary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-accent-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-neutral-content text-sm font-bold">A</div>
                                   </div>
                              </div>
                         </div>
                    </div>
               </div>
          </div>
          <div onclick="document.documentElement.setAttribute('data-theme', 'light'); localStorage.setItem('theme', 'light'); ch()"
               data-set-theme="dracula"
               class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
               data-act-class="!outline-base-content">
               <div data-theme="light" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                    <div class="grid grid-cols-5 grid-rows-3">
                         <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                         <div class="bg-base-300 col-start-1 row-start-3"></div>
                         <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                              <div class="font-bold">Light</div>
                              <div class="flex flex-wrap gap-1">
                                   <div
                                        class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-primary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-secondary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-accent-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-neutral-content text-sm font-bold">A</div>
                                   </div>
                              </div>
                         </div>
                    </div>
               </div>
          </div>
          <div onclick="document.documentElement.setAttribute('data-theme', 'retro'); localStorage.setItem('theme', 'retro'); ch()"
               data-set-theme="dracula"
               class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
               data-act-class="!outline-base-content">
               <div data-theme="retro" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                    <div class="grid grid-cols-5 grid-rows-3">
                         <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                         <div class="bg-base-300 col-start-1 row-start-3"></div>
                         <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                              <div class="font-bold">Retro</div>
                              <div class="flex flex-wrap gap-1">
                                   <div
                                        class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-primary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-secondary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-accent-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-neutral-content text-sm font-bold">A</div>
                                   </div>
                              </div>
                         </div>
                    </div>
               </div>
          </div>
          <div onclick="document.documentElement.setAttribute('data-theme', 'cyberpunk'); localStorage.setItem('theme', 'cyberpunk'); ch()"
               data-set-theme="dracula"
               class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent"
               data-act-class="!outline-base-content">
               <div data-theme="cyberpunk" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                    <div class="grid grid-cols-5 grid-rows-3">
                         <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                         <div class="bg-base-300 col-start-1 row-start-3"></div>
                         <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                              <div class="font-bold">Cyberpunk</div>
                              <div class="flex flex-wrap gap-1">
                                   <div
                                        class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-primary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-secondary-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-accent-content text-sm font-bold">A</div>
                                   </div>
                                   <div
                                        class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                                        <div class="text-neutral-content text-sm font-bold">A</div>
                                   </div>
                              </div>
                         </div>
                    </div>
               </div>
          </div>

     </div>
</center></template>