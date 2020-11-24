<template>
  <div>
    <p> Location: {{ location }} </p>
    <FormulateForm name="config" v-model="config" @submit="handleSubmit">
      <div class="center">
        <b-tabs content-class="mt-3 center">
          <b-tab title="General" active>
            <div id="general-tab-container" class="tab-container">
              <FormulateInput
                  type="text"
                  :label="labels.name"
                  validation="required"
                  :validation-name="labels.name"
                  error-behavior="blur"
                  :placeholder="placeholders.name"
                  v-model="config.name"
              />
              <FormulateInput
                  type="text"
                  :label="labels.buildCommand"
                  validation="required"
                  :validation-name="labels.buildCommand"
                  error-behavior="blur"
                  :placeholder="placeholders.buildCommand"
                  v-model="config.buildCommand"
              />
              <FormulateInput
                  type="text"
                  :label="labels.runCommand"
                  validation="required"
                  :validation-name="labels.runCommand"
                  error-behavior="blur"
                  :placeholder="placeholders.runCommand"
                  v-model="config.runCommand"
              />
              <FormulateInput
                  type="text"
                  :label="labels.runtimeImage"
                  validation="required"
                  :validation-name="labels.runtimeImage"
                  error-behavior="blur"
                  :placeholder="placeholders.runtimeImage"
                  v-model="config.runtimeImage"
              />
              <FormulateInput
                  type="text"
                  :label="labels.outputBinary"
                  validation="required"
                  :validation-name="labels.outputBinary"
                  error-behavior="blur"
                  :placeholder="placeholders.outputBinary"
                  v-model="config.outputBinary"
              />
            </div>
          </b-tab>
          <b-tab title="Network">
            <div id="network-tab-container" class="tab-container">
              <FormulateInput
                  type="group"
                  name="exposePorts"
                  :repeatable="true"
                  :label="labels.exposePorts"
                  :add-label="addLabels.exposePorts"
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
                  name="networks"
                  :repeatable="true"
                  :label="labels.networks"
                  :add-label="addLabels.networks"
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
            </div>
          </b-tab>
          <b-tab title="Environment">
            <div id="environemnt-tab-container" class="tab-container">
              <FormulateInput
                  type="group"
                  name="environment"
                  :repeatable="true"
                  :label="labels.environment"
                  :add-label="addLabels.environment"
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
            </div>
          </b-tab>
          <b-tab title="Debugger">
            <div id="debugger-tab-container" class="tab-container">
              <FormulateInput
                  type="checkbox"
                  :label="labels.debuggerEnabled"
                  :placeholder="placeholders.debuggerEnabled"
                  v-model="config.debuggerEnabled"
              />
              <FormulateInput
                  type="number"
                  :label="labels.debuggerPort"
                  :placeholder="placeholders.debuggerPort"
                  v-model="config.debuggerPort"
                  v-if="config.debuggerEnabled"
                  :help="help.debuggerPort"
              />
            </div>
          </b-tab>
        </b-tabs>
      </div>
      <div class="actions">
        <div class="padder">
          <FormulateInput type="submit" label="Save"/>
          <FormulateInput type="button" label="Reset" data-ghost @click="reset"/>
        </div>
      </div>
    </FormulateForm>
  </div>
</template>

<script>
import ConfigService from "../services/ConfigService";
import lang from "../utils/lang";

export default {
  data() {
    return {
      config: {},
      location: '',
    };
  },
  computed: {
    placeholders() {
      return lang.placeholders;
    },
    labels() {
      return lang.labels;
    },
    addLabels() {
      return lang.addLabels;
    },
    help() {
      return lang.help;
    },
  },
  async mounted() {
    const res = await ConfigService.get(this.location);
    this.config = res.config;
    this.location = res.location;
  },

  methods: {
    reset() {
      this.$formulate.reset("config");
    },
    handleSubmit(data) {
      this.$bvToast.toast('Gebug configuration was updated successfully', {
        title: 'Update Configuration',
        autoHideDelay: 5000,
      })
      ConfigService.save(data);
    },
  },
};
</script>

<style scoped>
@import "../assets/snow.min.css";

.padder {
  padding-top: 50px;
  display: flex;
}
.padder > .formulate-input:first-child {
  margin-right: 20px;
}

.actions {
  display: flex;
  justify-content: center;
}

.environment {
  display: flex;
  margin-right: 1em;
  justify-content: center;
}

.center {
  margin: auto;
  width: 50%;
  padding: 10px;
}

#general-tab-container {
  text-align: left;
}

.tab-container {
  text-align: left;
}
</style>