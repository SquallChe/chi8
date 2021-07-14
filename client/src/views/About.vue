<template>
  <div class="about">
    <v-expand-x-transition>
      <v-img src="../assets/coming-soon.jpg" v-show="show"/>
    </v-expand-x-transition>
  </div>
</template>

<script>
export default {
  data: () => ({
    show: false
  }),
  created() {
    this.test();
  },
  methods: {
    test() {
      setTimeout(() => {
        this.show = true
      }, 300)
    }
  }
  
  
}
</script>
