<template>
    <div>    
    <!-- <button v-on:click="count++">You clicked me {{ count }} times global: {{ globalCounter }}</button> -->
    <input v-bind="$attrs" :value="value" @input="$emit('input', $event.target.value)">
    </div>
</template>

<script>
export default {
    props: ['value'],
    inheritAttrs: false,
    data() {
        return {
            count: 0
        }
    },
}
</script>