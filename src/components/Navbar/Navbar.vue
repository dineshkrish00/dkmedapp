<template>
    <div>
      <v-toolbar app outlined class="green white--text">
        <v-img src="../../assets/med1.png" class="shrink mr-0" cover width="50"></v-img>
        <v-toolbar-title class="text-h5 font-weight-medium text-uppercase pl-10">{{ this.$store.state.useractive.role }}</v-toolbar-title>
        <v-spacer></v-spacer>
        <!-- <navbar @adichutan="beatCount" :numValue="childNum" content="Dashboard" iconname="view-dashboard-edit-outline" navigate="/dashboard"></navbar> -->
        <navbar content="Dashboard" iconname="view-dashboard-edit-outline" navigate="/dashboard"></navbar>
        <navbar  v-if="role=='Admin'" content="AddUser" iconname="account-plus-outline" navigate="/adduser"></navbar>
        <navbar  v-if="role=='Admin'"  content="Login history" iconname="history" navigate="/history"></navbar>   
        <navbar  v-if="role=='Inventry' ||role=='Manager'"  content="Stock entry" iconname="stocking" navigate="/stockentry"></navbar>
        <navbar  v-if="role=='Inventry' || role =='Manager' ||role=='Biller'"  content="Stock view" iconname="eye-arrow-left-outline" navigate="/stockview"></navbar>
        <navbar  v-if="role=='Biller'"  content="Bill entry" iconname="cash" navigate="/billentry"></navbar>
        <navbar  v-if="role=='Manager'" content="Sales report" iconname="chart-line" navigate="/salesreport"></navbar>
        <v-btn small class="ml-4 white green--text" large @click="logout"><v-icon>mdi-logout</v-icon>LOGOUT</v-btn>
        <!-- parentValue  {{ parentValue }} 
        <v-btn @click="childNum++" >parent btn</v-btn> -->
    </v-toolbar>
</div>
</template>
<script>
import logoutservice from "../../service/login/userloginservice";
import navbar from "../Navbar/Navtemp.vue";

export default{
  data(){
    return{
      role : this.$store.state.useractive.role,
      userid: this.$store.state.useractive.userid,
      date:'',
      sign:false,
      overlay:false,
    }
  },
  components : {
    navbar
  },
  methods:{
   
  logout(){
    this.sign=true
    if (this.sign==true){
    console.log("logout....")
    ///write api ogic for logout data
    logoutservice.LogoutData(this.userid)
    .then(response=>{
          if(response.data.status==="S"){
            this.$store.state.useractive.role = null;
            this.snackbar=true;        
            this.colour='green';
            this.text="logout Successfully"; 
            this.$router.push('/'); 
          }else{
            this.Msg=response.errMsg;
            this.text="Logout was failed";
            this.colour='red';
            this.snackbar = true;
            console.log("invalid crendentials")
          }
        }).catch((error)=>{
          console.log("logout submit btn error",error);
        });
      }else{
        console.log("couldn't logout")
      }
  },
  }
}
</script>
