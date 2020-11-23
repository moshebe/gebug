<template>
  <div id="app">
    <img alt="Gebug" src="./assets/logo.png" width="500" height="300" />
    <div style="text-align: center;">
      <ConfigForm :location="settings.location" />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import settings from "@/utils/settings";
import ConfigForm from "./components/ConfigForm.vue";

@Component({
  components: {
    ConfigForm,
  },
  computed: {
    settings() {
      return settings;
    }
  }
})
export default class App extends Vue {}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
