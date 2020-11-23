<template>
  <div id="app">
    <img alt="Gebug" src="./assets/logo.png" width="500" height="300" />
    <center>
      <ConfigForm :location="'/Users/moshe/Dev/cpp-gebug'" />
    </center>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import ConfigForm from "./components/ConfigForm.vue";

@Component({
  components: {
    ConfigForm,
  },
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
