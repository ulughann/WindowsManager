<script setup>
import Navbar from './components/Navbar.vue'
import { CheckChoco, Install, InstallApp, RenderAppsList, Search } from '../wailsjs/go/main/App'

import { onMounted } from 'vue';

onMounted(async () => {
  let v = await CheckChoco()
  if (v == "Error") {
    document.querySelector(".stat").innerHTML = `
    <div class="stat-title">Chocolatey Not Found</div>
    <div class="stat-value">
      <button class="btn btn-primary" id="install">Install Chocolatey</button>
    </div>`
  } else {
    document.querySelector(".stat-value").innerHTML = v
  }
  if (document.querySelector("#install")) {
    document.querySelector("#install").addEventListener("click", async () => {
      document.querySelector(".stat").innerHTML = `
    <div class="stat-title">Installing Chocolatey</div>
    <div class="stat-value">
      <button class="btn btn-primary" id="install">Installing...</button>
    </div>
    `
      let v = await Install()
      if (v == "Error") {
        document.querySelector(".stat").innerHTML = `
    <div class="stat-title">Chocolatey Not Found</div>
    <div class="stat-value">
      <button class="btn-primary" id="install">Install Failed</button>
    `
      } else {
        document.querySelector(".stat-title").innerHTML = "Chocolatey Version"
        document.querySelector(".stat-value").innerHTML = v
      }
    })
  }

  let str = await RenderAppsList()
  let pkgs = document.querySelector("#pkgs")
  pkgs.innerHTML += str


  let install = document.querySelectorAll("[id^='install_']")
  install.forEach((e) => {
    e.addEventListener("click", async () => {
      if (document.querySelector(".stat-title").innerHTML == "Chocolatey Not Found") {
        alert("Please install Chocolatey first")
        return
      }
      e.innerHTML = "Installing..."
      let app = e.id.split("_")[1]
      let v = await InstallApp(app)
      if (v == "Error") {
        e.innerHTML = "Install Failed"
      } else {
        e.innerHTML = "Installed"
        e.classList.remove("btn-neutral")
        e.classList.add("btn-success")
      }
    })
  })

  let search = document.querySelector("#confirms")
  search.onclick = async () => {
    console.log(document.querySelector("#boxs").value)
    let str = await Search(document.querySelector("#boxs").value)

    let pkgs = document.querySelector("#packages")
    pkgs.innerHTML = str.split(" ").join("<br>")
    my_modal_1.showModal()
    document.querySelector("#boxs").value = ""
  }

  document.querySelector("#confirm").addEventListener("click", async () => {
    if (document.querySelector(".stat-title").innerHTML == "Chocolatey Not Found") {
      alert("Please install Chocolatey first")
      return
    }
    let apps = document.querySelector(".input").value
    let arr = apps.split(",")
    arr.forEach(async (e) => {
      let v = await InstallApp(e)
      if (v == "Error") {
        document.querySelector("#confirm").innerHTML = "Install Failed"
        document.querySelector("#confirm").classList.remove("btn-primary")
        document.querySelector("#confirm").classList.remove("btn-success")
        document.querySelector("#confirm").classList.add("btn-error")
        document.querySelector("#box").value = ""
      } else {
        document.querySelector("#confirm").innerHTML = "Installed"
        document.querySelector("#confirm").classList.remove("btn-primary")
        document.querySelector("#confirm").classList.add("btn-success")
        document.querySelector("#box").value = ""
        setTimeout(() => {
          document.querySelector("#confirm").innerHTML = "Download"
          document.querySelector("#confirm").classList.remove("btn-success")
          document.querySelector("#confirm").classList.add("btn-primary")
        }, 3000)
      }
    })
  })
})
</script>

<template>
  <Navbar />
  <dialog id="my_modal_1" class="modal">
    <div class="modal-box">
      <div class="overflow-x-auto">
        <p class="
        whitespace-pre-wrap
        " id="packages">

        </p>
      </div>
      <div id="insallah" class="input-group mt-4">
        <input id="box" type="text" class="input input-bordered w-[80%]" placeholder="App, App1, App2" />
        <button id="confirm" class="btn btn-primary input-group-btn">Download</button>
      </div>
      <div class="modal-action">
        <form method="dialog">
          <button class="btn">Close</button>
        </form>
      </div>
    </div>
  </dialog>
  <div style="width: 80%; margin-left: 120px;">
    <div class="hero mt-2 flex items-start text-left  left bg-gradient-to-r from-primary to-red rounded-box
    ">
      <div class="stats bg-transparent">
        <div id="stat" class="stat">
          <div class="stat-title">Chocolatey Version</div>
          <div class="stat-value">fetching...</div>
        </div>
      </div>
    </div>
    <div id="search" class="input-group mt-4">
      <input id="boxs" type="text" class="input input-bordered w-[80dvw]" placeholder="Search for an app" />
      <button id="confirms" class="btn btn-primary input-group-btn" onclick="my_modal_1.showModal()">Search</button>
    </div>
    <div id="pkgs" class="grid w-[100%] mt-4 lg:grid-cols-5 xl:grid-cols-6 2xl:grid-cols-7 gap-x-8 gap-y-2">
      <div class="item bg-base-200 w-[200px] h-[200px] rounded-box">
        <center>
          <box-icon name='discord' size="128px" type='logo' color='#7289da'></box-icon>
          <br><a id="install_discord" class="btn btn-neutral">Install Discord</a>
        </center>
      </div>
      <div class="item bg-base-200 w-[200px] h-[200px] rounded-box">
        <center>
          <div class="h-[112px]"><img sizes="128px" class="pt-8" src="https://www.7-zip.org/7ziplogo.png"></div>
          <br><a id="install_7zip" class="btn btn-neutral">Install 7Zip</a>
        </center>
      </div>
    </div>
  </div>
</template>