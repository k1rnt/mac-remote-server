<template>
  <div>
    <navi />
    <h1>Settings!</h1>
    <p>Hello! {{ user.name }}</p>
    <p>You age is {{ user.age }}</p>
  </div>
</template>

<script　lang="ts">
import navi from "../components/navi.vue";

const user = {name: 'test', age: 20};

export default {
  data(){
    return {
      user: user
    }
  },
  components: { navi },
};
</script>

<style>
</style>