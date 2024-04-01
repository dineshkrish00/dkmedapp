<template>
    <v-container>
      <span class="green--text text-h3 font-weight-bold">DASHBOARD</span>
      <v-card class="outlined elevation-9" color="green lighten-1 mt-15 white--text"  max-width="600">
  <v-card-title>
      <span v-show="temp">{{ show }}</span>
        <p class="text-h4 font-weight-bold white--text text-uppercase">Your today sales </p>
        <v-spacer></v-spacer>
        <v-icon color="yellow" size="45">mdi-currency-usd</v-icon>
  </v-card-title>
    <v-card-text>
      <span class="text-h4 font-weight-bold black--text red--text ml-2">
        <v-icon color="white" size="45">mdi-currency-usd</v-icon>
        {{ todaysales }}</span>
        <span class="text-h5 white--text font-weight-bold ml-2">
           <v-icon :color="checking ? 'purple' : 'red'" class="ml-1" size="50">mdi-{{arrow}}</v-icon>
       ({{percent}}%)</span>
      <v-spacer></v-spacer>
    </v-card-text>
      </v-card>
    </v-container>
</template>


<script>
import dashservice from "../../service/dashboard/dashboard"
  export default{
    data(){
      return{
        temp:false,
        checking:'',
        arrow: '',
        percent:0,
        percolor:'',
        date:'',
        todaysales:'',
        yesterdaysales:'',
        difference:'',
      }
    },
    methods:{
      updatevalues(){
        dashservice.DashboardReport()
      .then((response) =>{
            if (response.data.status == "S") {
              this.medtemp=response.data
              this.todaysales=Number(this.medtemp.tsales)
              this.yesterdaysales=Number(this.medtemp.ysales)
            }else{
              console.log("Fetching data failed");
            }}
          ).catch((error)=>{
            console.log("Error in_ Billentry vue :", error) 
        });

          if(this.yesterdaysales!=0){
            this.difference=Number(this.todaysales)-Number(this.yesterdaysales);
            let val=((Number(this.difference)) * 10) / Number(this.yesterdaysales);
              if(this.todaysales!=''){
                this.percent=Number(val.toFixed(1));
              }
              else{
                this.percent=0;
              }
                if(this.todaysales>this.yesterdaysales){
                    this.arrow="arrow-up";
                    this.checking=true;
                    this.percolor="green";
                  }
                else if(this.todaysales<this.yesterdaysales){
                    this.arrow="arrow-down";
                    this.checking=false;
                }
          }   
      }
    },
    computed:{
    show(){
      this.updatevalues();
      return 0;
    },
  }
  }
</script>