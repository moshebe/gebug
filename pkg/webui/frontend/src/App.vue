<template>
  <div id="app">
    <img alt="Gebug logo" src="./assets/logo.png">
    <p> Location: {{ location }}</p>
    <!-- <Config v-if="hasProject"/> -->
    <Config/>
    <!-- <ParentForm /> -->



  </div>
</template>

<script>
import Config from './components/Config.vue'
// import axios from "axios";
// import ParentForm from './components/ParentForm.vue'

export default {
  name: 'App',
  components: {
    Config
    // ParentForm
  },
  data() {
    return {
      hasProject: false,
      location: '',
      config: null,
    };
  },
  async created() {
  //   const config = await axios.get(`http://127.0.0.1:3030/config`)
  //   if (config.status == 200){
  //     // console.log('config: ', config.data)
  //     // location = config.data.location
  //
  //   } else {
  //     console.log("response: ", config)
  //   }
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
