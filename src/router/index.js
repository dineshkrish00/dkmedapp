import Vue from "vue";
import VueRouter from "vue-router";
import login from "../views/login.vue";
import history from "../views/LoginHistory.vue";
import adduser from "../views/AddUser.vue";
import billentry from "../views/BillEntry.vue";
import salesreport from "../views/Salesreport.vue";
import stockentry from "../views/StockEntry.vue";
import stockview from "../views/Stockview.vue";
import dashboard from "../views/Dashboar.vue";


Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "login",
    component: login,
  },
  {
    path: "/history",
    name: "history",
    component: history,
  },
 
  {
    path: "/adduser",
    name: "adduser",
    component: adduser,
  },
  {
    path: "/billentry",
    name: "billentry",
    component: billentry,
  },
  {
    path: "/salesreport",
    name: "salesreport",
    component: salesreport,
  },
  {
    path: "/stockentry",
    name: "stockentry",
    component: stockentry,
  },
  {
    path: "/stockview",
    name: "stockview",
    component: stockview,
  },
  {
    path: "/dashboard",
    name: "dashboard",
    component: dashboard,
  },

  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
];

const router = new VueRouter({
  routes,
});

export default router;
