<template>
  <v-snackbar v-model="snackbar" :timeout="timeout" :color="color" centered>
    <v-spacer></v-spacer>
    <div>{{ message }}</div>
    <v-spacer></v-spacer>
  </v-snackbar>
</template>

<script>
export default {
  name: "Snackbar",

  data: () => ({
    snackbar: false,
    //timeout: 2000
  }),
  methods: {
    show() {
      this.snackbar = true
    }
  },
  props: {
    message: {
      required: true,
      type: String
    },
    timeout: {
      type: Number,
      default: 4000
    },
    color: {
      type: String,
      default: ""
    }

  }
};
</script>