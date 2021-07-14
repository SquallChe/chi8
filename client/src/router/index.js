import Vue from "vue";
import VueRouter from "vue-router";
import ItemList from "@/views/ItemList.vue";
import Category from "@/views/Category.vue";
import Detail from "@/views/Detail.vue";
import Regist from "@/views/Regist.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "RecommandList",
    component: ItemList
  },
  {
    path: "/category/:id",
    name: "Category",
    component: Category
  },
  {
    path: "/item-list",
    name: "ItemList",
    component: ItemList
  },
  {
    path: "/search-list/:id",
    name: "SearchList",
    component: ItemList
  },
  {
    path: "/detail/:id",
    name: "Detail",
    component: Detail
  },
  {
    path: "/regist",
    name: "Regist",
    component: Regist
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
