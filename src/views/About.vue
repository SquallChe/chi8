<template>
  <div class="about">
    <h1>This is an about page</h1>
    <v-snackbar v-model="snackTest" timeout="3000" centered>This is a test.</v-snackbar>
    <v-btn @click="test">Test</v-btn>
  </div>
</template>

<script>
export default {
  data: () => ({
    snackTest: false
  }),
  methods: {
    test() {
      this.snackTest = true
    }
  }
}
</script>
