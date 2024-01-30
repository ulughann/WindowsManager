<script setup>
import Navbar from './components/Navbar.vue'
import { OsInfo, Memory, CpuInfo, GpuInfo } from '../wailsjs/go/main/App'

document.addEventListener('DOMContentLoaded', async (event) => {
  let a = await OsInfo();
  document.getElementById('os').innerHTML = `${a[1]} ${a[4]}`;

  let m = await Memory();
  m = m.replaceAll("(","").replaceAll(")","").split(' ')
  m = m[1].split('GB')[0] + "gb total, " + m[3].split('GB')[0];
  document.getElementById('memory').innerHTML = `${m}gb available`;

  let c = await CpuInfo();
  document.getElementById('cpu').innerHTML = `${c[0]}`;

  let g = await GpuInfo();
  document.getElementById('gpu').innerHTML = `${g[0]}`;

})

</script>

<template>
  <Navbar />
  <div class="hero min-h-screen bg-base">
    <div class="hero-content text-center">
      <div class="max-w-lg">
        <h1 class="text-5xl font-bold">Windows Manager</h1>
        <p class="py-6">Get started with a new Windows install, faster than ever.</p>
        <ul class="menu menu-vertical lg:menu-horizontal bg-base-200 rounded-box">
          <li><a href="/#/scripts">Scripts</a></li>
          <li><a href="/#/store">Apps</a></li>
          <li><a href="/#/settings">Settings</a></li>
        </ul>
        <div class="w-full">
          <div role="alert" class="alert mt-2  rounded-t-box rounded-b-none">
            <box-icon name='windows' type='logo' color='#ffffff' class="mr-2"></box-icon>
            <span>Windows Version:</span>
            <div>
              <a id="os" class="bg-base-300 rounded-box p-3">Windows</a>
            </div>
          </div>
          <div role="alert" class="alert mt-1 rounded-t-none rounded-b-none">
            <box-icon type='solid' name='memory-card' color='#ffffff' class="mr-2"></box-icon>
            <span>Memory:</span>
            <div>
              <a id="memory" class="bg-base-300 rounded-box p-3">Fetfching...</a>
            </div>
          </div>
          <div role="alert" class="alert mt-1 rounded-t-none rounded-b-none">
            <box-icon type='solid' name='game' color='#ffffff' class="mr-2"></box-icon>
            <span>Gpu:</span>
            <div>
              <a id="gpu" class="bg-base-300 rounded-box p-3">Fetfching...</a>
            </div>
          </div>
          <div role="alert" class="alert mt-1 rounded-t-none rounded-b-box">
            <box-icon type='solid' name='analyse' color='#ffffff' class="mr-2"></box-icon>
            <span>Cpu:</span>
            <div>
              <a id="cpu" class="bg-base-300 rounded-box p-3">Fetfching...</a>
            </div>
          </div>
        </div>
      </div>
    </div>
</div></template>