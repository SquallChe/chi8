<template>
  <div>
    <v-card class="d-flex flex-wrap justify-space-between px-4" flat>
      <template v-for="(item, idx) in items">
        <router-link :to="{name:'ItemList',params:{mode: 2, categoryId: $route.params.id, kindId: item.kindId}}" :key="idx" style="text-decoration:none;">
          <v-card class="mt-5" flat>
            <v-avatar size="140" tile>
              <v-img :src="'https://chi8.store/img/category/' + $route.params.id + '/' + item.kindId + '.png'"></v-img>
            </v-avatar>
            <div class="d-flex">
              <span class="mx-auto">{{ item.kindName }}</span>
            </div>
          </v-card>
        </router-link>
      </template>
    </v-card>
  </div>
</template>

<script>
export default {
  data: () => ({
    items: []
  }),
  created() {
    this.getItems()
  },
  methods: {
    getItems() {

      this.$http.get('/kind/' + this.$route.params.id)
      .then(response => {
        window.console.log(response.data)
        this.items = response.data
      })
    },
  },
  watch: {
    '$route.path': function() {
      this.getItems()
    }
  }
}
</script>
