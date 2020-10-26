<template>
  <div id="app">
    <AutoForm />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import AutoForm from "@/components/AutoForm.vue";

@Component({
  components: {
    AutoForm
  }
})
export default class App extends Vue {}
</script>

<style scoped>
#app {
  font-family: "sofia-pro", sans-serif;
  text-align: center;
  padding: 10vh 15vw;
}
</style>