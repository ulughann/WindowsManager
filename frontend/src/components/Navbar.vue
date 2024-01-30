<script setup>
import { onMounted } from 'vue';
import { RouterLink, RouterView } from 'vue-router'
const activeClass = "flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-1.5 bg-base-100 text-base-content bg-primary-focus"
const nonActiveClass = "flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-1.5 text-base-content hover:bg-primary-focusflex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-1.5 text-base-content hover:bg-primary-focus"


// get current route
const currentRoute = window.location.href.split("#")[1];
let activeTab = "home";
if (currentRoute === "/") {
  activeTab = "home";
} else {
  activeTab = currentRoute.split("/")[1];
} 

onMounted(() => {
  const active = document.getElementById(activeTab);
  active.className = activeClass;

  const boxIcons = document.getElementsByTagName("box-icon");
  for (let i = 0; i < boxIcons.length; i++) {
    const theme = localStorage.getItem("theme")
    if (["light", "corporate", "retro"].includes(theme)) {
      boxIcons[i].setAttribute("color", "#000000");
    } else {
      boxIcons[i].setAttribute("color", "#ffffff");
    }
  }
})
</script>

<template>
    <div class="relative bg-gray-50 dark:bg-slate-900">
      <nav

        class="z-20 flex shrink-0 grow-0 justify-around gap-4 border-t border-base-200 bg-base-100/50 p-2.5 shadow-lg backdrop-blur-lg fixed top-2/4 -translate-y-2/4 left-6 min-h-[auto] min-w-[64px] flex-col rounded-lg border">
        <a id="home" href="/"
          class="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-1.5 text-base-content hover:bg-primary-focusflex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-1.5 text-base-content hover:bg-primary-focus">
          <box-icon name='home' type='solid'></box-icon>
          <small className="text-xs font-medium">Home</small>
        </a>
        <a href="/#/scripts" id="scripts"
          class="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-1.5 text-base-content hover:bg-primary-focus">
          <box-icon name='code-curly'></box-icon>
          <small class="text-center text-xs font-medium"> Scripts </small>
        </a>
        <a href="/#/store" id="store"
          class="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-1.5 text-base-content hover:bg-primary-focus">
          <box-icon name='package' type='solid'></box-icon>
          <small class="text-center text-xs font-medium"> Apps </small>
        </a>
        <a href="/#/monitor" id="monitor"
          class="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-1.5 text-base-content hover:bg-primary-focus">
          <box-icon name='analyse' type='solid'></box-icon>
          <small class="text-center text-xs font-medium"> System Monitor </small>
        </a>
        <hr class="dark:border-gray-700/60" />
        <a href="/#/settings" id="settings"
          class="flex aspect-square min-h-[32px] w-16 flex-col items-center justify-center gap-1 rounded-md p-1.5 text-base-content hover:bg-primary-focus">
          <box-icon name='cog' type='solid' color="#ffffff"></box-icon>
          <small class="text-center text-xs font-medium"> Settings </small>
        </a>
      </nav>
    </div>
</template>
