<script setup>
import Navbar from './components/Navbar.vue'
import { onMounted } from 'vue';


onMounted(() => {
  
})


</script>

<template>
  <Navbar />
  <center>
    <div style="width: 90%; margin-left: 120px" class="rounded-box">
      <br>
      <div class="collapse collapse-plus mb-2	bg-base-200">
        <input type="radio" name="my-accordion-3" checked="checked" />
        <div class="collapse-title text-xl font-medium">
          Windows Settings
        </div>
        <div class="collapse-content">
          <div class="text-1xl bg-base-300 p-2 rounded-box mb-1"
            style="white-space: nowrap;overflow: hidden;text-overflow: ellipsis;display: flex;align-items: center;justify-content: space-between;text-align: left;width: 100%;">
            <a class="m-2 mb-2">
              Add "open with cmd" to context menu</a>
            <input type="button" name="switch" class="btn rounded-box" value="activate"
              data="resources/AddOpenWithCMD.bat" />
          </div>
          <div class="text-1xl bg-base-300 p-2 rounded-box mb-1"
            style="white-space: nowrap;overflow: hidden;text-overflow: ellipsis;display: flex;align-items: center;justify-content: space-between;text-align: left;width: 100%;">
            <a class="m-2 mb-2">
              Disable office telemetry</a>
            <input type="button" name="switch" class="btn rounded-box" value="activate"
              data="resources/DisableOfficeTelemetry.reg" />
          </div>











        </div>
      </div>
      <div class="collapse collapse-plus mb-2 bg-base-200">
        <input type="radio" name="my-accordion-3" />
        <div class="collapse-title text-xl font-medium">
          Optimizations
        </div>
        <div class="collapse-content">
          <p>hello</p>
        </div>
      </div>
      <div class="collapse collapse-plus mb-2 bg-base-200">
        <input type="radio" name="my-accordion-3" />
        <div class="collapse-title text-xl font-medium">
          Debloaters
        </div>
        <div class="collapse-content">
          <p>hello</p>
        </div>
      </div>
    </div>
  </center>
</template>