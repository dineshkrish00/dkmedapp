<template>
  <v-app>
    <Navi v-if="this.$store.state.useractive.role != null"></Navi>

    <v-main>
      <router-view />
    </v-main>
    {{ this.$store.state.Login }}
  </v-app>
</template>

<script>
import Navi from "../src/components/Navbar/Navbar.vue";

export default {
  name: "navApp",
  components: {
    Navi,
  },
  mounted() {
      this.$router.push("/");
  },
};
</script>
