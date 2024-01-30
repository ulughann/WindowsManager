<script setup>
import Navbar from "./components/Navbar.vue";
import { onMounted, ref } from "vue";

import {
  CpuInfo,
  GpuInfo,
  GetCPUFrequencies,
  GetMemUsage,
  CpuFreq,
  GetStorage
} from "../wailsjs/go/main/App";

onMounted(async () => {
  let cpuFreq = await CpuFreq();

  const labels =
    cpuFreq.map((_, index) => `Core ${index + 1}`);

  const cpuchart =
    new Chart("myChart", {
      type: "bar",
      scales: {
        y: {
          beginAtZero: true,
        },
      },
      data: {
        labels: labels,
        datasets: [{
          label: 'CPU Frequency (%)',
          data: cpuFreq.map((_, index) => cpuFreq[index].toFixed(2)),
          backgroundColor: [
            'rgba(255, 99, 132, 0.2)',
            'rgba(255, 159, 64, 0.2)',
            'rgba(255, 205, 86, 0.2)',
            'rgba(75, 192, 192, 0.2)',
            'rgba(54, 162, 235, 0.2)',
            'rgba(153, 102, 255, 0.2)',
            'rgba(201, 203, 207, 0.2)',
            'rgba(255, 99, 132, 0.2)',
            'rgba(255, 159, 64, 0.2)',
            'rgba(255, 205, 86, 0.2)',
            'rgba(75, 192, 192, 0.2)',
            'rgba(54, 162, 235, 0.2)',
            'rgba(153, 102, 255, 0.2)',
            'rgba(201, 203, 207, 0.2)',
            'rgba(255, 99, 132, 0.2)',
            'rgba(255, 159, 64, 0.2)',
            'rgba(255, 205, 86, 0.2)',
            'rgba(75, 192, 192, 0.2)',
            'rgba(54, 162, 235, 0.2)',
            'rgba(153, 102, 255, 0.2)',
            'rgba(201, 203, 207, 0.2)',
          ],
          borderColor: [
            'rgb(255, 99, 132)',
            'rgb(255, 159, 64)',
            'rgb(255, 205, 86)',
            'rgb(75, 192, 192)',
            'rgb(54, 162, 235)',
            'rgb(153, 102, 255)',
            'rgb(201, 203, 207)',
            'rgb(255, 99, 132)',
            'rgb(255, 159, 64)',
            'rgb(255, 205, 86)',
            'rgb(75, 192, 192)',
            'rgb(54, 162, 235)',
            'rgb(153, 102, 255)',
            'rgb(201, 203, 207)',
            'rgb(255, 99, 132)',
            'rgb(255, 159, 64)',
            'rgb(255, 205, 86)',
            'rgb(75, 192, 192)',
            'rgb(54, 162, 235)',
            'rgb(153, 102, 255)',
            'rgb(201, 203, 207)',
          ],
          borderWidth: 1
        }]
      },
    });
  const memchart = new Chart("memchart", {
    type: "doughnut",
    label: "memory usage",
    data: {
      labels: ["Used Memory", "Free Memory"],
      datasets: [
        {
          label: "Memory Usage",
          data: [0, 0],
          backgroundColor: ['rgba(255, 99, 132, 0.2)', 'rgba(54, 162, 235, 0.2)'],
          borderColor: ['rgb(255, 99, 132)', 'rgb(54, 162, 235)'],
          borderWidth: 1,
          hoverBackgroundColor: ["#FF6384", "#36A2EB"],
        },
      ],
    },
  });
  const storechart = new Chart("storechart", {
    type: "doughnut",
    label: "storage usage",
    data: {
      labels: ["x disk", "y disk"],
      datasets: [
        {
          label: "Storage Usage",
          data: [0, 0],
          backgroundColor: ['rgba(255, 99, 132, 0.2)', 'rgba(54, 162, 235, 0.2)'],
          borderColor: ['rgb(255, 99, 132)', 'rgb(54, 162, 235)'],
          borderWidth: 1,
          hoverBackgroundColor: ["#FF6384", "#36A2EB"],
        },
      ],
    },
  });

  let c = global.m ? global.m : await CpuInfo();
  let g = global.g ? global.g : await GpuInfo();
  let m = global.m ? global.m : await GetMemUsage();
  let f = global.f ? global.f : await GetCPUFrequencies();
  setInterval(async () => {
    let storage = await GetStorage();
    m = await GetMemUsage();
    let total = m[0].total;
    total = total / 1000000000;
    total = total.toFixed(2);
    let labels = storage.map((_, index) => storage[index].label);
    let data = storage.map((_, index) => ((storage[index].size_bytes) / 1000000000).toFixed(2));

    storechart.data.labels = labels;
    storechart.data.datasets[0].data = data;
    storechart.update();
    


    memchart.data.datasets[0].data = [(m[0].used / 1000_000).toFixed(2), (m[0].free / 1000_000).toFixed(2)];
    memchart.update();


    let free = m[0].free;
    free = free / 1000000000;
    free = free.toFixed(2);

    let fre = await CpuFreq();
    fre = fre.reduce((a, b) => a + b, 0) / fre.length;
    fre = fre.toFixed(2);
    document.getElementById("total").innerHTML = `${total}gb`;
    document.getElementById("free").innerHTML = `${free}gb`;
    document.getElementById("percent").innerHTML = `${m[0].usedPercent}%`;

    let cpuload = await CpuFreq();
    cpuload = cpuload.map((_, index) => cpuload[index].toFixed(2));
    cpuchart.data.datasets[0].data = cpuload;
    cpuchart.update();

    // await cpuchart.update();
    localStorage.setItem("cpu", JSON.stringify(cpuload));
  }, 1000);

  m = await GetMemUsage();
  let total = m[0].total;
  total = total / 1000000000;
  total = total.toFixed(2);

  let free = m[0].free;
  free = free / 1000000000;
  free = free.toFixed(2);

  document.getElementById("cpu").innerHTML = `${c[0]}`;
  document.getElementById("core").innerHTML = `${c[1]}`;
  document.getElementById("thread").innerHTML = `${c[3]}`;
  document.getElementById("gpu").innerHTML = `${g}`;
  document.getElementById("base").innerHTML = `${(f[0] / (1000)).toFixed(2)}ghz`;
  document.getElementById("total").innerHTML = `${total}gb`;
  document.getElementById("free").innerHTML = `${free}gb`;
  document.getElementById("percent").innerHTML = `${m[0].usedPercent}%`;
});
document.addEventListener("DOMContentLoaded", async (event) => { });
</script>

<template>
  <Navbar />
  <center>
    <div class="content w-[90%] ml-[120px]">
      <!-- create grid of 2x1 -->
      <div class="grid mt-2 grid-cols-2 gap-4">
        <div class="bg-base-200 rounded-box">
          <h1 class="text-1xl font-bold">Cpu Overview</h1>
          <div class="text-left p-2 m-2">
            <div class="flex flex-row justify-between">
              <div class="flex flex-col">
                <span class="text-sm">Cpu Name</span>
                <span class="text-sm">Cpu Cores</span>
                <span class="text-sm">Cpu Threads</span>
                <span class="text-sm">Base Frequency</span>
              </div>
              <div class="flex flex-col text-right">
                <span class="text-sm" id="cpu">Cpu</span>
                <span class="text-sm" id="core">Cpu Cores</span>
                <span class="text-sm" id="thread">Cpu Threads</span>
                <span class="text-sm" id="base">Base Frequency</span>
              </div>
            </div>
          </div>
        </div>
        <div>
          <div class="bg-base-200 rounded-box">
            <h1 class="text-1xl font-bold">Gpu Overview</h1>
            <div class="text-left p-2 m-2">
              <div class="flex flex-row justify-between">
                <div class="flex flex-col">
                  <span class="text-sm">Gpu</span>
                </div>
                <div class="flex flex-col">
                  <span class="text-sm" id="gpu">Fetching...</span>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-base-200 rounded-box">
            <h1 class="text-1xl font-bold">Memory Overview</h1>
            <div class="text-left p-2 m-2">
              <div class="flex flex-row justify-between">
                <div class="flex flex-col">
                  <span class="text-sm">Total Memory</span>
                  <span class="text-sm">Free Memory</span>
                  <span class="text-sm">% used</span>
                </div>
                <div class="flex flex-col text-right">
                  <span class="text-sm" id="total">Fetching...</span>
                  <span class="text-sm" id="free">Fetching...</span>
                  <span class="text-sm" id="percent">Fetching...</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="bg-base-200 w-[100%] mt-2 rounded-box">
        <canvas id="myChart" class="text-base-300 " style="width: 100%; max-width: 700px"></canvas>
      </div>
      <div class="grid mt-2 grid-cols-2 gap-4">
        <div class="bg-base-200 w-[100%] mt-2 rounded-box">
          <canvas id="memchart" class="text-base-300 " style="width: 100%; max-width: 700px"></canvas>
        </div>
        <div class="bg-base-200 w-[100%] mt-2 rounded-box">
          <canvas id="storechart" class="text-base-300 " style="width: 100%; max-width: 700px"></canvas>
        </div>
      </div>
    </div>
  </center>
</template>
