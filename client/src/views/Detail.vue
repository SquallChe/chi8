<template>
  <div>
    <v-carousel v-model="activeSlide" :show-arrows="showArrow" :height="height" hide-delimiters>
      <v-carousel-item
        v-for="(item,i) in imageItems"
        :key="i"
        :src="item.src"
      ></v-carousel-item>
      <!-- <v-carousel-item src='./img/items/1/1-1.png'></v-carousel-item>
      <v-carousel-item src='./img/items/1/1-2.png'></v-carousel-item>
      <v-carousel-item src='./img/items/1/1-3.png'></v-carousel-item>
      <v-carousel-item src='./img/items/1/1-4.png'></v-carousel-item> -->
    </v-carousel>
    <!-- <v-chip class="page black white--text" :style="{top: (height - 50) + 'px'}" small label>
      <span class="mr-2">{{ activeSlide + 1 + ' / ' + imageItems.length }}</span>
    </v-chip> -->
    <v-sheet class="page black white--text" :style="{top: (height - 30) + 'px', fontSize:'10px'}">
      <span class="pa-2">{{ activeSlide + 1 + ' / ' + imageItems.length }}</span>
    </v-sheet>
    <v-card flat class="pa-4">
      <v-row class="error--text display-1">
        <v-col>¥ {{ formatAmount(priceJP) }}</v-col>
      </v-row>
      <v-row class="title">
        <v-col>{{ itemName }}</v-col>
      </v-row>
    </v-card>
    <v-btn v-if="$route.params.mode" class="back grey darken-1" fab x-small>
      <v-icon class="white--text" @click="goBack">mdi-chevron-left</v-icon>
    </v-btn>
    <v-menu dark nudge-bottom="50">
      <template v-slot:activator="{on}">
        <v-btn v-on="on" class="more grey darken-1" fab x-small>
          <v-icon class="white--text">mdi-dots-horizontal</v-icon>
        </v-btn>
      </template>
      <v-list width="170">
        <!-- <template v-for="(item,index) in menus">
          <v-list-item :key="index">
            <v-list-item-avatar>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-avatar>
            <v-list-item-content>
              {{ item.title }}
            </v-list-item-content>
          </v-list-item>
          <v-divider inset v-if="menus.length -1 != index" :key="index"></v-divider>
        </template> -->
        <router-link :to="{name:'RecommandList'}" style="text-decoration:none;">
          <v-list-item>
            <v-list-item-avatar>
              <v-icon>mdi-home</v-icon>
            </v-list-item-avatar>
            <v-list-item-content>
              首页
            </v-list-item-content>
          </v-list-item>
        </router-link>

        <v-divider inset></v-divider>

        <v-list-item @click="copyUrl">
          <v-list-item-avatar>
            <v-icon>mdi-link-variant</v-icon>
          </v-list-item-avatar>
          <v-list-item-content>
            复制链接
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-menu>
    <v-snackbar v-model="snackbar" multi-line="true" :timeout="timeout" color="blue lighten-2" centered>
      <v-row class="title" justify="center">
        <v-icon class="mr-5" color="white" large left>mdi-check-circle-outline</v-icon>
        <span>复制成功，快去粘贴吧。</span>    
      </v-row>
    </v-snackbar>
  </div>
</template>

<script>
export default {
  data: () => ({
    imageItems: [],
    itemName: '',
    priceJP: '',
    showArrow: false,
    height: 0,
    snackbar: false,
    timeout: 3000,
    activeSlide: 0,
    menus: [
      {icon: 'mdi-home',title: '首页'},
      {icon: 'mdi-link-variant',title: '复制链接'}
    ]
  }),
  created() {
    this.getDetailImage()
  },
  methods: {
    getDetailImage() {
      this.$http.get('/detail',{
        params: {
          itemId: this.$route.params.id,
          mode: this.$route.params.mode || 9,
          keyWord: this.$route.params.keyWord || ''
        }
      }).then(response => {
        window.console.log(response.data)
        let data = response.data;
        this.imageItems = data.imageItems
        this.itemName = data.itemName
        this.priceJP = data.priceJP
      })
    },
    formatAmount(num) {
      return num.toString().replace(/(\d)(?=(?:\d{3})+$)/g, '$1,')
    },
    goBack() {
      let page = this.$route.params.page
      let pageCount = this.$route.params.pageCount
      let position = this.$route.params.position
      let mode = this.$route.params.mode
      let categoryId = this.$route.params.categoryId
      let kindId = this.$route.params.kindId
      let keyWord = this.$route.params.keyWord
      
      this.$router.push({name:'ItemList', 
        params:{
          page: page, 
          pageCount: pageCount,
          position: position,
          mode: mode,
          categoryId: categoryId,
          kindId: kindId,
          keyWord: keyWord
          }
      })
    },
    copyUrl() {
        let el = document.createElement('textarea')
        el.value = document.URL
        document.body.appendChild(el)
        el.select()
        el.setSelectionRange(0, 99999) /*For mobile devices*/
        document.execCommand('copy')
        document.body.removeChild(el)
        this.snackbar = true
    }
  },
  mounted() {
    this.height = document.documentElement.clientWidth
  }
}
</script>

<style lang="scss" scoped>
  .back {
    position: fixed;
    top: 5px;
    left: 5px;
    /*opacity: 0.5;*/
    z-index: 99999999;
  }

  .more {
    position: fixed;
    top: 5px;
    right: 5px;
    /*opacity: 0.5;*/
    z-index: 99999999;
  }

  .page {
    position: absolute;
    top: 0px;
    right: 0px;
    opacity: 0.3;
    border-radius: 10px 0 0 10px;
  }
  
</style>
