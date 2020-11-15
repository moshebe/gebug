<template>
<div>
    Global:
    <input placeholder="Standard Input" v-model="globalCounter">
    <ChildInput v-model="globalCounter"/>
</div>
    
    <!-- <ChildInput/> -->
</template>
<script>
import ChildInput from "@/components/ChildInput";
export default {
    name: "parent",
    components: {ChildInput},
    data() {
        return {
            globalCounter: 0,
        }
    },    
}
setTimeout(() => {  
  console.log("timeout called"); 
  // this.data.config.name = "MDASDASDSADAS";

  }, 5000);
</script>