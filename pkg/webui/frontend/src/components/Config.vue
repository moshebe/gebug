<template>
  <html>
  <head>
    <meta charset="utf-8">
    <meta content="width=device-width,initial-scale=1,minimal-ui" name="viewport">
  </head>
  <div class="jumbotron">
    <div class="container">
      <div class="row">
        <div class="col-sm-8 offset-sm-2">
          <div>
            <h2>Gebug Configuration</h2>
            <form @submit.prevent="handleSubmit">
              <ConfigInput fieldName="Name" type="text" :placeholder="placeholders.name" v-model="config.name" />
              <ConfigInput fieldName="Build Command" type="text" :placeholder="placeholders.buildCommand" v-model="config.buildCommand" />
              <ConfigInput fieldName="Run Command" type="text" :placeholder="placeholders.runCommand" v-model="config.runCommand" />
              <ConfigInput fieldName="Runtime Image" type="text" :placeholder="placeholders.runtimeImage" v-model="config.runtimeImage" />


              <ConfigInput fieldName="Debugger Port" type="number" :placeholder="placeholders.debuggerPort.toString()" v-model="config.debuggerPort" />
              <ConfigInput fieldName="Debugger Enabled" type="checkbox" :placeholder="placeholders.debuggerEnabled.toString()" v-model="config.debuggerEnabled" />
<!-- TODO: Debugger enabled (T/F) -->

              <div class="input-group mb-3">
                <div class="input-group-prepend">
                  <span class="input-group-text">Name</span>
                </div>
                <input type="text" class="form-control" placeholder="awesome-app"
                       aria-label="Project Name" aria-describedby="name"
                       :class="{ 'is-invalid': submitted && $v.config.name.$error }"
                       v-model="config.name"/>
                <div v-if="submitted && !$v.config.name.required" class="invalid-feedback">Name is required</div>
              </div>

              <!-- <div class="input-group mb-3">
                <div class="input-group-prepend">
                  <span class="input-group-text">Output Artifact</span>
                </div>
                <input type="text" class="form-control" placeholder="/app" aria-describedby="Output Artifact"
                       :class="{ 'is-invalid': submitted && $v.config.outputBinPath.$error }"
                       v-model="config.outputBinPath"/>
                <div v-if="submitted && !$v.config.outputBinPath.required" class="invalid-feedback">Output artifact path is required</div>
              </div> -->

              <!-- <div class="form-group">
                <label for="outputBinPath">Output Binary Path</label>
                <input type="text" class="form-control"
                       :class="{ 'is-invalid': submitted && $v.config.outputBinPath.$error }"
                       v-model="config.outputBinPath" id="outputBinPath" name="outputBinPath"/>
                <div v-if="submitted && !$v.config.outputBinPath.required" class="invalid-feedback">Output Binary Path
                  is required
                </div>
              </div> -->
              <!-- <div class="form-group">
                <label for="buildCommand">Build Command</label>
                <input type="text" class="form-control"
                       :class="{ 'is-invalid': submitted && $v.config.buildCommand.$error }"
                       v-model="config.buildCommand" id="buildCommand" name="buildCommand"/>
                <div v-if="submitted && !$v.config.buildCommand.required" class="invalid-feedback">Build Command is
                  required
                </div>
              </div>
              <div class="form-group">
                <label for="runCommand">Run Command</label>
                <input type="text" class="form-control"
                       :class="{ 'is-invalid': submitted && $v.config.runCommand.$error }"
                       v-model="config.runCommand" id="runCommand" name="runCommand" value="config.runCommand"/>
                <div v-if="submitted && !$v.config.runCommand.required" class="invalid-feedback">Run Command is
                  required
                </div>
              </div>
              <div class="form-group">
                <label for="runtimeImage">Runtime Image</label>
                <input type="text" class="form-control"
                       :class="{ 'is-invalid': submitted && $v.config.runtimeImage.$error }"
                       v-model="config.runtimeImage" id="runtimeImage" name="runtimeImage"/>
                <div v-if="submitted && !$v.config.runtimeImage.required" class="invalid-feedback">Runtime Image is
                  required
                </div>
              </div> -->


              <div class="form-group">
                <button class="btn btn-primary">Reload</button>
                <button class="btn btn-primary">Save</button>                
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
  </html>
</template>

<script>
import {required} from "vuelidate/lib/validators";
import axios from 'axios';
import ConfigInput from "@/components/ConfigInput";

export default {
  name: "app",
  components: {ConfigInput},
  mounted() {
    console.log("mounted");
    // TODO: refactor the base url & path
    axios.get('http://localhost:3030/config?path=/Users/moshe/Dev/cpp-gebug').then(
      response => {
        // TODO: check why it's under 'data.data' and not a single nesting level
        this.config = response.data.data.config;
        this.config.name='';
        }
      );
  },
  data() {
    return {
      config: {
        name: "",
        outputBinPath: "",
        buildCommand: "",
        runCommand: "",
        runtimeImage: "",
        debuggerPort: 0,
        debuggerEnabled: false,
        exposePorts: [],
        networks: [],
        environment: [],
      },
      placeholders: {
        name: "awesome-app",
        outputBinPath: "",
        buildCommand: "go build -o {{.output_binary}}",
        runCommand: "{{.output_binary}}",
        runtimeImage: "golang:latest",
        debuggerPort: 0,
        debuggerEnabled: false,
        exposePorts: [],
        networks: [],
        environment: [],
      },
      submitted: false
    };
  },
  validations: {
    config: {
      name: {required},
      outputBinPath: {required},
      buildCommand: {required},
      runCommand: {required},
      runtimeImage: {required},
    }
  },
  methods: {
    handleSubmit() {
      console.log('handle submit called');
      this.submitted = true;

      // TODO: validation
      // stop here if form is invalid
      this.$v.$touch();
      if (this.$v.$invalid) {
        return;
      }

      console.log(JSON.stringify(this.config));
      // alert("SUCCESS!! :-)\n\n" + e + JSON.stringify(this.config));
    }
  },
};
// setTimeout(() => {  
//   console.log("timeout called in parent");
//   self.config = {
//     name: 'my-app',
//     outputBinPath: '/app',
//     buildCommand: 'go build -o {{.output_binary}}',
//     runtimeImage: 'golang:1.15.2'
//   }
//   // console.log("after set self.config");
//   // console.log("this.config: ", this.config);
//   // console.log("self.config: ", self.config);
//   // console.log("config: ", config);
  
//   }, 5000);
</script>

<style>
.invalid-feedback {
  color: red;
}
</style>