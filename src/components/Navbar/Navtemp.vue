<template>
  <router-link
    :to="navigate"
    class="ml-5 mr-1 text-decoration-none text-uppercase white--text text-body-1 font-weight-medium"
  >
    <v-container class="d-flex justify-center align-center">
      <v-icon class="white--text">mdi-{{ iconname }}</v-icon>
    </v-container>
    <!-- <v-btn @click="emitValue" small >child btn</v-btn> childValue {{ numValue }} -->
    <v-divider></v-divider>
    <span>{{ content }}</span>
  </router-link>
</template>

<script>
export default {
  props: {
    iconname: String,
    content: String,
    navigate: String,
    // numValue : Number,
  },
  // emit funciton--
  // data () {
  //   return {
  //     childValue :0,
  //   }
  // },
  // methods : {
  //   emitValue() {
  //     let obj = {}
  //     obj.beat = 2
  //     obj.kick = 1
  //     obj.attack = ["punch", "bat", true]
  //     this.$emit("adichutan", obj)
  //   }
  // }
};
</script>
