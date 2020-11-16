<template>
  <div>
    <FormulateForm
    name="config"
    v-model="config"
    @submit="handleSubmit"
    >
      <FormulateInput
        type="text"
        label="Name"
        validation="required"
        validation-name="Name"
        error-behavior="live"
        :placeholder="placeholders.name"
        v-model="config.name"
      />
      <FormulateInput
        type="text"
        label="Build Command"
        validation="required"
        validation-name="Build Command"
        error-behavior="live"
        :placeholder="placeholders.buildCommand"
        v-model="config.buildCommand"
      />
      <FormulateInput
        type="text"
        label="Run Command"
        validation="required"
        validation-name="Run Command"
        error-behavior="live"
        :placeholder="placeholders.runCommand"
        v-model="config.runCommand"
      />
      <FormulateInput
        type="text"
        label="Runtime Image"
        validation="required"
        validation-name="Runtime Image"
        error-behavior="live"
        :placeholder="placeholders.runtimeImage"
        v-model="config.runtimeImage"
      />
      <FormulateInput
        type="checkbox"
        label="Debugger Enabled"
        :placeholder="placeholders.debuggerEnabled"
        v-model="config.debuggerEnabled"
      />
      <FormulateInput
        type="number"
        label="Debugger Port" 
        :placeholder="placeholders.debuggerPort"
        v-model="config.debuggerPort"      
        v-if="config.debuggerEnabled == true"
      />

      <FormulateInput
        type="group"
        name="exposePorts"
        :repeatable="true"
        label="Expose Ports"
        add-label="+ Add Port"        
        :values="config.exposePorts"
      >
       <div class="ports">
          <FormulateInput            
            name="port"            
            validation="required|number|between:1,65535"
            :placeholder="placeholders.exposePorts"    
          />
        </div>
     </FormulateInput>
     <FormulateInput
        type="group"
        name="environment"
        :repeatable="true"
        label="Environment"
        add-label="+ Add Environment Variable"
        :values="config.environment"
      >
       <div class="environment">
          <FormulateInput            
            name="envName"        
            :placeholder="placeholders.envName"    
            validation="required"
          />
          <FormulateInput
            name="envValue"                        
            :placeholder="placeholders.envValue"    
          />
        </div>
     </FormulateInput>        
      <FormulateInput
        type="group"
        name="networks"
        :repeatable="true"
        label="Networks"
        add-label="+ Add Network"        
        :values="config.networks"
      >
       <div class="networks">
          <FormulateInput
            name="network"                        
            validation="required"
            :placeholder="placeholders.networks"    
          />
        </div>
     </FormulateInput>
      <div class="actions">          
        <FormulateInput type="submit" label="Save" />        
      <FormulateInput
        type="button"
        label="Reset"
        data-ghost
        @click="reset"
      />
      </div>
      
      
      
    </FormulateForm>
  </div>
</template>

<script>
import ConfigService from '../services/ConfigService';

export default {
   data () {
    return {
      config: {},
      placeholders: {
        name: "awesome-app",
        outputBinPath: "",
        buildCommand: "go build -o {{.output_binary}}",
        runCommand: "{{.output_binary}}",
        runtimeImage: "golang:latest",
        debuggerPort: 4321,
        debuggerEnabled: false,
        exposePorts: 'PORT[:PORT]',
        networks: "private-network",
        envName: "FOO",
        envValue: "BAR",
      },
    }
  },
  mounted() {
    console.log("mounted");
    ConfigService.get('/Users/moshe/Dev/cpp-gebug').then(
      response => {          
        this.config = ConfigService.decodeModel(response.data.data.config);
        console.log('got response from sever, set config to: ', this.config, response.data.data.config);
        }
      );
  },

  methods: {
    reset () {
      console.log('reset called');    
      console.log(JSON.stringify(this.config));
      this.$formulate.reset('config')
    },
    handleSubmit(data) {
      console.log("config: ", JSON.stringify(this.config));
      console.log('handle submit called: ', data);     
      ConfigService.save(ConfigService.encodeModel(data));
    }
  },
};

</script>

<style scoped>
@import '../assets/snow.min.css';
.actions {
  display: flex;
  margin-bottom: 1em;
}
.actions .formulate-input {
  margin-right: 1em;
  margin-bottom: 0;
}

 .environment {
    display: flex;
    margin-right: 1em;
  }
</style>