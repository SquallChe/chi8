<template>
  <v-app>
    <sidebar ref="sidebar"></sidebar>
    <v-app-bar v-if="this.$route.name != 'Detail'" app dark color="red" hide-on-scroll scroll-threshold="200">
      <v-app-bar-nav-icon @click="showDrawer"></v-app-bar-nav-icon>
      <div class="d-flex align-center purple--text">
        <!-- <div class="title">蒲公英ファッションショップ</div> -->
        <!-- <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png"
          transition="scale-transition"
          width="40"
        /> -->

        <!-- <v-img
          alt="Vuetify Name"
          class="shrink mt-1 hidden-sm-and-down"
          contain
          min-width="100"
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-name-dark.png"
          width="100"
        /> -->
      </div>

      <v-spacer></v-spacer>
      <v-text-field ref="searchText" placeholder="儿童 外套 黑色" v-if="this.$route.name != 'About' && this.$route.name != 'Regist'" v-model="keyWord" dense append-icon="mdi-magnify" rounded outlined style="height:40px;" @keypress.enter="search" @click:append="search"></v-text-field>
      <!-- <v-btn
        href="https://github.com/vuetifyjs/vuetify/releases/latest"
        target="_blank"
        text
      >
        <span class="mr-2">Latest Release</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn> -->
    </v-app-bar>

    <v-main>
      <router-view />
      <v-snackbar v-model="snackbar" color="pink lighten-2" :timeout="timeout">请输入关键词，如需要使用多个关键词请使用空格分割。</v-snackbar>
    </v-main>
    <!-- <v-footer color="primary" fixed>
      <v-row class="justify-space-around">
        <v-icon>mdi-home</v-icon>
        <v-icon>mdi-cart</v-icon>
        <router-link to="/regist"><v-icon>mdi-account</v-icon></router-link>
      </v-row>
    </v-footer> -->
    <v-bottom-navigation v-model="value" dark shift fixed grow :background-color="color">
      <v-btn @click="$router.push({'name':'RecommandList'})">
        <span>Home</span>
        <v-icon>mdi-home</v-icon>
      </v-btn>
      
      <v-btn @click="$router.push({'name':'About'})">
        <span>Cart</span>
        <v-icon>mdi-cart</v-icon>
      </v-btn>
      <v-btn @click="$router.push({'name':'Regist'})">
        <span>Account</span>
        <v-icon>mdi-account</v-icon>
      </v-btn>
    </v-bottom-navigation>
  </v-app>
</template>

<script>
import Sidebar from "./components/Sidebar.vue";

export default {
  name: "App",

  components: {
    Sidebar
  },
  methods: {
    showDrawer() {
      this.$vuetify.goTo(0,{duration:0})
      this.$refs.sidebar.drawer = true
    },
    search() {
      this.$refs.searchText.blur()
      if(this.keyWord == '') {
        this.snackbar = true
        return
      }
      this.$router.push({name: 'SearchList', params:{id: Math.ceil(Math.random() * 100), mode: 3,keyWord: this.keyWord}})
    }
  },
  computed: {
    // key() {
    //   return this.$router.name ? this.$router.name + new Date() : this.$router + new Date()
    // }
    color() {
      switch (this.value) {
        case 0: return 'blue-grey'
        case 1: return 'teal'
        case 2: return 'brown'
        default: return 'blue-grey'
      }
    }
  },

  data: () => ({
    keyWord: '',
    timeout: 2000,
    snackbar: false,
    value: 0
  })
};
</script>
