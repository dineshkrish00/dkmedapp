<template>
  <div>
     <v-card fluid class="mt-16 mx-16 ">
       <v-card-title class="text-h5 green--text font-weight-bold">Sales Report</v-card-title>
       <v-card-text>
         <v-row justify="space-around">
           <v-col cols="2">
             <v-text-field v-model="start" outlined label="Start date" type="date"></v-text-field>
           </v-col>
           <v-col cols="2">
             <v-text-field v-model="end" outlined label="End date" type="date"></v-text-field>
           </v-col>
           <v-col cols="2">
             <v-btn @click="change" class="green white--text">SEARCH</v-btn>
           </v-col>
         </v-row>
       </v-card-text>
     </v-card>
     <v-main>
       <v-data-table :headers="headers" :items="filteredItems" :search="search" class="elevation-1 my-12 mx-15">
         <template v-slot:top>
           <v-toolbar flat>
             <v-text-field v-model="search" label="Search" class="mx-4"></v-text-field>
           </v-toolbar>
         </template>
       </v-data-table>
     </v-main>
     <v-snackbar v-model="snackbar" class="">{{ text }}
       <template v-slot:action="{ attrs }">
         <v-btn :color="colour" text v-bind="attrs" @click="snackbar = false">Close</v-btn></template>
       </v-snackbar>
  </div>
 </template>
 
 <script>
 import saleservice from "../../service/salesreport/salesreport"
 export default {
  data() {
     return {
       search: "",
       start: "",
       end: "",
       scheck: false,
       text: "",
       colour: "green",
       snackbar: false,
       startdate:'',
       enddate:'',
       headers: [
         { text: "Billno", value: "billNo" },
         { text: "Bill Date", value: "billDate" },
         { text: "Medicine Id", value: "medicineid" },
         { text: "Qty", value: "quantity" },
         { text: "Amount", value: "amount" },
       ],
       filteredItems: [],
     };
  },
  methods: {
    change(){
      if (this.start != "" && this.end != "") {
      saleservice.SalesReport(this.start,this.end)
      .then((response) =>{
            if (response.data.status == "S") {
              this.snackbar=true; 
              this.filteredItems=response.data.salesArr
              this.text="Sales Report Fetched succesfully";
      }else{
        this.snackbar=true;
        this.text="Fail to Add Report";
        }
      }
      ).catch((error)=>{
        console.log("Error in AddBill_SERVICES :", error) 
    });
  }else {
         this.snackbar = true;
         this.text = "First Select the date";
       }
       this.start=''
       this.end=''
     },
  },

 };
 </script>