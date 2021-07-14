<template>
  <div v-scroll="getData">
    <v-card class="pa-3" v-for="(item,idx) in items" :key="idx" flat>
      <router-link :to="{name:'Detail',
          params:{
            id: item.ItemId, 
            page: page,
            pageCount: pageCount, 
            position: position,
            mode: $route.params.mode || 1,
            categoryId: $route.params.categoryId || '',
            kindId: $route.params.kindId || '',
            keyWord: $route.params.keyWord || '',
          }
        }" style="text-decoration:none;">
        <div class="d-flex black--text">
          <v-avatar size="125" tile>
            <v-img :src="'https://chi8.store/img/items/' + item.ItemId + '/' + item.ItemId + '.png'"></v-img>
          </v-avatar>
          <div class="ml-4">
              <div class="headline">{{ item.ItemName }}</div>
              <div class="error--text">¥ {{ formatAmount(item.PriceJP) }}</div>
              <div>{{ item.Description }}</div>
          </div>
        </div>
      </router-link>
    </v-card>
    <!-- <v-dialog v-model="isLoading">
      <v-img src="../assets/loading.gif" width="50"></v-img>
    </v-dialog> -->
    <v-chip v-if="items.length != 0" id="vchip" color="black white--text page">{{ `${page + ' / ' + pageCount}` }}</v-chip>
    <v-snackbar v-model="snackbar" multi-line color="blue lighten-2" :timeout="timeout" centered>
      <v-row justify="center" class="title">
        <v-icon class="mr-1">mdi-emoticon-dead</v-icon>
        暂无相关商品，看看其他的商品吧
      </v-row>
      </v-snackbar>
    <!-- <v-img v-if="items.length == 0" src="../assets/not-ready.png" v-cloak></v-img> -->
    <!-- <v-icon @click="test">mdi-account</v-icon> -->
  </div>
</template>

<script>
export default {
  data: () => ({
    items: [],
    page: 0,
    pageCount: 0,
    isLoading: false,
    loadAll: false,
    timeoutId: '',
    position: 0,
    snackbar: false,
    timeout: 3000
  }),
  created() {
    this.getPageCount()
  },
  methods: {
    getPageCount() {
      
      if(this.$route.params.pageCount) {
        this.items = this.$store.getters.getItems
        this.page = parseInt(this.$route.params.page)
        this.pageCount = parseInt(this.$route.params.pageCount)
        this.$vuetify.goTo(parseInt(this.$route.params.position),{duration: 0})
        return
      }
      
      this.$http.get('/page-count',{
        params: {
          mode: this.$route.params.mode || 1,
          categoryId: this.$route.params.categoryId || '',
          kindId: this.$route.params.kindId || '',
          keyWord: this.$route.params.keyWord || ''
        }
      }).then(response => {
        window.console.log(response.data)
        this.pageCount = parseInt(response.data.count)
        this.getItems()

        if(this.pageCount == 0) {
          this.snackbar = true
        }
      })
    },
    getItems() {
      //this.page = this.page + 1
      
      // if(this.page > this.pageCount || this.loadAll) {
      //   this.page = this.pageCount
      //   return
      // }
      if(this.loadAll) {
        return
      }

      if(this.page + 1 == this.pageCount) {
        this.loadAll = true
      }
      
      this.$http.get('/item',{
        params: {
          mode: this.$route.params.mode || 1,
          categoryId: this.$route.params.categoryId || '',
          kindId: this.$route.params.kindId || '',
          keyWord: this.$route.params.keyWord || '',
          page: this.page + 1
        }
      }).then(response => {
        window.console.log(response.data)
        response.data.forEach(item => {
          this.items.push(item)
        })
        this.$store.commit('setItems', this.items)
        this.isLoading = false
      })
    },
    getData() {
      let scrollTop = document.documentElement.scrollTop || document.body.scrollTop
      let clientHeight = document.documentElement.clientHeight
      let scrollHeight = document.documentElement.scrollHeight
      this.position = scrollTop

      console.log(this.items.length)
      console.log(this.items.length/10)
      let pageHeight = scrollHeight / (this.items.length / 10)
      //window.console.log(scrollTop)
      //window.console.log(scrollHeight)
      //window.console.log(this.items.length)
      this.page = Math.floor((scrollTop + clientHeight) / pageHeight) + 1
      if(this.page > this.pageCount) {
        this.page = this.pageCount
      }
      
      document.getElementById('vchip').style.bottom = '75px'
      document.getElementById('vchip').style.display = ''
      clearTimeout(this.timeoutId)
      this.timeoutId = setTimeout(() => {
        this.hidePage()
      }, 1500);

      //if(scrollTop + clientHeight == scrollHeight) {
      if((scrollTop + clientHeight) / scrollHeight >= 0.96) {
        if(!this.isLoading) {
          this.isLoading = true
          this.getItems()
        }
      }
    },
    formatAmount(num) {
      return num.toString().replace(/(\d)(?=(?:\d{3})+$)/g, '$1,')
    },
    hidePage() {
      document.getElementById('vchip').style.display = 'none'
    },
    test() {
      let scrollTop = document.documentElement.scrollTop
      let clientHeight = document.documentElement.clientHeight
      let scrollHeight = document.documentElement.scrollHeight
      let scrollTop2 = document.body.scrollTop
      alert(scrollTop)
      alert(clientHeight)
      alert(scrollHeight)
      alert(scrollTop2)
    }
  },
  watch: {
    '$route.path': function() {
      window.console.log(this.$route.path)
      this.page = 0
      this.items = []
      this.loadAll = false
      this.getPageCount()
    }
  }
}
</script>

<style lang="scss" scoped>
  .page {
    position: fixed;
    bottom: -100px;
    left: calc(50% - 26px);
    opacity: 0.5;
    background-color: black;
  }

  [v-cloak] {
    display: none;
  }
</style>
