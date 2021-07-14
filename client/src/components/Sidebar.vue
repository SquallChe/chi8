<template>
<v-navigation-drawer
      v-model="drawer"
      absolute
      temporary
      color="yellow lighten-5"
    >
      <v-list nav dark>
        <!-- <v-list-item>
          <v-list-item-avatar>
            <v-icon large>mdi-account-circle</v-icon>
          </v-list-item-avatar>
  
          <v-list-item-content>
            <v-list-item-title class="subtitle-1">{{userName}}</v-list-item-title>
            <v-list-item-subtitle class="yellow--text">
              {{ roleName }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
  
        <v-divider></v-divider> -->

        <v-list-item v-for="item in items" :key="item.MenuId"
          :class="$route.path.length != 1 ? ($route.path == '/' + item.link.toLowerCase() + '/' + item.id ? 'purple' : '') : ($route.name == item.link ? 'purple' : '')">
          <v-list-item-icon>
            <v-icon class="text-h3" color="light-green">
              mdi-{{ item.icon }}
            </v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <router-link style="text-decoration:none;"
              :to="{name: item.link, params: {id: item.id}}"
              class="light-green--text font-weight-black text-h3"
            >
              <v-list-item-title>{{ item.title }}</v-list-item-title>
            </router-link>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <!-- <template v-slot:append>
        <div class="pa-2">
          <v-btn block dark v-if="isLogin" @click="logout">ログアウト</v-btn>
        </div>
      </template> -->
    </v-navigation-drawer>
</template>

<script>
export default {
  name: "Sidebar",

  data: () => ({
    drawer: false,
    items: [
    //   { title: '热门推荐', icon: 'fire', link: 'RecommandList'},
    //   { title: '男装', icon: 'face', link: 'Category', id: '001'},
    //   { title: '女装', icon: 'face-woman', link: 'Category', id: '002'},
    //   { title: '童装', icon: 'baby-face', link: 'Category', id: '003'}
      ]
  }),
  created() {
    this.getMenu()
    window.console.log('ok')
  },
  methods: {
    getMenu() {
      this.$http.get('/menu')
      .then(response => {
        if(response.data.result == '0') {
            this.items = response.data.menu
        }
        else {
            window.console.log(response.data.errMsg)
        }
        
      }).catch(err => {
        window.console.log(err)
      })
    },
  },
  computed: {
    // isLogin() {
    //   return this.$store.getters.getLoginStatus 
    // },
    // userName() {
    //   return this.$store.getters.getUserName
    // },
    // bg() {
    //   return this.$vuetify.theme.dark ? '' : "primary"
    // }
  },
};
</script>