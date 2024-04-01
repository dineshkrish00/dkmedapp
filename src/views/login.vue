<template>
  <v-container class="auto d-flex justify-center align-center"> 
    <v-card class="elevation-16 mt-15" width="350px">
    <v-card-text >
      <v-form v-model="valid" ref="form" lazy-validation class="mp-16 py-8  px-6" >
        <div class="d-flex justify-center">
    <v-img src="../assets/med1.png" class="shrink mr-0"  width="100"></v-img> </div>
     <h2 class="green--text text-h5 font-weight-bold py-8 text-center">LOG-IN</h2>
     <v-text-field v-model="email" :rules="emailRules" outlined  label="E-mail" required class="py-2" ></v-text-field>
     <v-text-field  v-model="password" :rules="passwordRules" outlined  label="Password"  required :type="'password'"></v-text-field>
     <div class="d-flex justify-center">

     <v-btn color="green" class="white--text" @click="submit">Login <v-icon right>
        mdi-open-in-new 
      </v-icon></v-btn>
    </div>
     <v-snackbar v-model="snackbar" class="" >{{ text }}
      <template v-slot:action="{ attrs }">
        <v-btn :color="colour" text v-bind="attrs" @click="snackbar = false">Close</v-btn></template>
        </v-snackbar>
    </v-form>
</v-card-text>
</v-card>
<v-overlay :value="overlay">
            <v-progress-circular indeterminate size="64"></v-progress-circular>
        </v-overlay>
  </v-container>
</template>

<style scoped>
 .d-flex-align-center {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
 }
</style>

<script>
import logservice from "../service/login/userloginservice"

export default {
  name: "loginname",
  data() {
      return{
          valid: true,
          colour: '',
     emailRules: [ v => !!v || 'User ID is required',   v => /.+\..+/.test(v) || 'userID must be valid'],
     passwordRules: [ v => !!v || 'Password is required'],
     snackbar:false, passing: false, 

    email: '',    
    password: '',
     text:'', 
     overlay: false,
     log:"logout", 
     date:'', 
     userId:'',
     role:'',
      }
  },
  methods:{
      submit(){
        logservice.GetData(this.email, this.password)
        .then(response=>{
          console.log("succesfully response",response.data.userRole)
          if(response.data.status==="S"){
            this.userId=response.data.userName;
            this.role=response.data.userRole;
            this.$store.state.useractive.role=this.role
            this.$store.state.useractive.userid=this.userId
            this.overlay = true
            setTimeout(() => {
              this.snackbar=true;  
              this.colour='green';
            this.text="login Successfully";
                     }, 300)
            setTimeout(() => {
            this.$router.push('/dashboard'); 
          }, 1000)
          }else{
            this.Msg=response.errMsg;
            this.text="Invalid user details";
            this.colour='red';
            this.snackbar = true;
            console.log("invalid crendentials")
          }
        }).catch((error)=>{
          console.log("login submit btn error",error);
        });
    },
      submitLog() {
        let entry = {userid : this.email , Type : 'login' , Date : this.date=this.convertDate(new Date())};
          this.$store.state.log_history.push(entry);
          //userid updating
          this.$store.commit('updateuser',this.role);
        console.log('pushed'+this.role)
      }, 
      convertDate(date) {
      var day = date.getDate().toString().padStart(2, '0');
      var month = (date.getMonth() + 1).toString().padStart(2, '0'); 
      var year = date.getFullYear();

      var hours = date.getHours();
      var minutes = date.getMinutes();
      var seconds = date.getSeconds();

    return day  + '/' + month + '/' + year + ' ' + hours + ':' + minutes + ':' + seconds;
  },
  },
}
</script>