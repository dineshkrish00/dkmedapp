<template>
    <v-app>
      <v-container>
        <h1 class="green--text text-center">STOCK VIEW</h1>
        <v-divider></v-divider>
        <v-container fluid>
        <v-card >
      <v-card-title class="ml-8 green--text text-h5 font-weight-medium">
        Stock View Page
        <v-spacer></v-spacer>
        <v-text-field class="red--text" v-model="search" append-icon="mdi-magnify" label="Search"
        single-line
          hide-details
        ></v-text-field>
      </v-card-title>
      <v-data-table  class="text-h5 font-weight-medium px-8"
        :headers="headers"
        :items="list"
        :search="search"
      ></v-data-table>
    </v-card>
    </v-container>
  </v-container>
  </v-app>
  </template>

<script>
import CurrentStock from "../../service/stockview/stockview"
export default {
  data(){
    return{
        search: "",
        headers: [
          { text: 'Medicine Name', align: 'start',sortable: false, value: 'medicinename',},
          { text: 'Brand', value: 'brand' },
          { text: 'Quantity', value: 'quantity' },
          { text: 'Unitprice', value: 'unitprice' },],
        list:[],
    }
  },
  methods:{
    StockView(){
      CurrentStock.CurrentStockData()
      .then((response) =>{
            if (response.data.status == "S") {
                this.list=response.data.stockViewArr;
            }
          }).catch((error)=>{
            console.log("Error in AddMEdStock_SERVICES :", error) 
        });
    }
  },
  mounted(){
    this.StockView()
  }
  };
</script>