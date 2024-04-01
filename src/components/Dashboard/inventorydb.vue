<template>
    <v-container>
      <v-card class="outlined elevation-9" color="green lighten-1 mt-15 white--text"  max-width="600">
  <v-card-title>
      <span v-show="temp">{{ show }}</span>
        <p class="text-h4 font-weight-bold white--text text-uppercase">Current Inventry values </p>
        <v-spacer></v-spacer>
        <v-icon color="yellow" size="45">mdi-palette-swatch-variant</v-icon>
  </v-card-title>
    <v-card-text>
        <span class="text-h4 white--text font-weight-bold ml-2">
            <v-icon color="white" size="45">mdi-palette-swatch-variant</v-icon>
    {{ values }}</span>
      <v-spacer></v-spacer>
    </v-card-text>
      </v-card>
    </v-container>
</template>

<script>
import inventservice from "../../service/dashboard/dashboard"
    export default{
        data(){
            return{
            temp:false,
            values:'',
            arr:[],
            checking:'',
            oldtemp:'',
            }
        },
        methods:{
            invenvalue(){ 
            inventservice.DashboardReport()
            .then((response) =>{
            if (response.data.status == "S") {
                this.values=response.data.curvalue
                }else{
                console.log("Fetching data failed");
                }}
            ).catch((error)=>{
                console.log("Error in_ Billentry vue :", error) 
            });
            } 
        },
        computed:{
            show(){
            this.invenvalue();
            return 0;
            },
        }
    }
</script>