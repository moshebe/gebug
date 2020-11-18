<template>
  <FormulateForm name="config" v-model="config" @submit="handleSubmit">
    <div class="inputs">
      <FormulateInput
        type="text"
        :label="labels.name"
        validation="required"
        :validation-name="labels.name"
        error-behavior="live"
        :placeholder="placeholders.name"
        v-model="config.name"
      />
      <FormulateInput
        type="text"
        :label="labels.buildCommand"
        validation="required"
        :validation-name="labels.buildCommand"
        error-behavior="live"
        :placeholder="placeholders.buildCommand"
        v-model="config.buildCommand"
      />
      <FormulateInput
        type="text"
        :label="labels.runCommand"
        validation="required"
        :validation-name="labels.runCommand"
        error-behavior="live"
        :placeholder="placeholders.runCommand"
        v-model="config.runCommand"
      />
      <FormulateInput
        type="text"
        :label="labels.runtimeImage"
        validation="required"
        :validation-name="labels.runtimeImage"
        error-behavior="live"
        :placeholder="placeholders.runtimeImage"
        v-model="config.runtimeImage"
      />
      <FormulateInput
        type="text"
        :label="labels.outputBinary"
        validation="required"
        :validation-name="labels.outputBinary"
        error-behavior="live"
        :placeholder="placeholders.outputBinary"
        v-model="config.outputBinary"
      />
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
        v-if="config.debuggerEnabled == true"
      />
    </div>
    <div class="actions">
      <FormulateInput type="submit" label="Save" />
      <FormulateInput type="button" label="Reset" data-ghost @click="reset" />
    </div>
  </FormulateForm>
</template>

<script>
import ConfigService from "../services/ConfigService";
import lang from "../lang";

export default {
  props: {
    location: String,
  },
  data() {
    return {
      config: {},
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
  },
  async mounted() {
    const remoteConfig = await ConfigService.get(this.location);
    this.config = remoteConfig;
  },

  methods: {
    reset() {
      this.$formulate.reset("config");
    },
    handleSubmit(data) {
      ConfigService.save(this.location, data);
    },
  },
};
</script>

<style scoped>
@import "../assets/snow.min.css";
.inputs {
  display: grid;
  padding: 150px;
  justify-content: center;
}
.actions {
  display: flex;
  margin-bottom: 1em;
  justify-content: center;
}
.actions .formulate-input {
  margin-right: 1em;
  margin-bottom: 0;
}

.environment {
  display: flex;
  margin-right: 1em;
  justify-content: center;
}
</style>