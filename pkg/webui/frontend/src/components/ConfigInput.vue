<template>
  <div class="input-group mb-3">
    <div class="input-group-prepend">
      <span class="input-group-text"> {{ fieldName }}</span>
    </div>
  
  <input v-bind="$attrs" :value="value" :placeholder="placeholder"
  @input="$emit('input', $event.target.value)" :type="type || text" class="form-control" />

 <!-- :class="{ 'is-invalid': submitted && $v.model.$error }"  -->
  <!-- <div v-if="submitted && required" class="invalid-feedback">Name is required</div> -->
  </div>
</template>

<script>

export default {
  name: "ConfigInput",
  props: {    
    value: [Number, String, Array, Boolean], // TODO: check this one
    type: String,
    fieldName: String,
    placeholder: String,
    validator: Object,
  },
  data() {
   return {
    //  submitted: false,
    //  hello: 'world',
      // fieldname: String,
      // placeholder: String,
    } 
  },
  // methods: {
  //   handleSubmit(e) {
  //     this.submitted = true;
  //     console.log('1', e);
  //     this.$emit("inputData", this.hello);
  //     console.log('2');
  //     this.submitted = true;
  //     console.log('3');
  //   },
  // }
}
// setTimeout(() => {  
//   console.log("child");
//   }, 5000);
</script>

<style scoped>

</style>