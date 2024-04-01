<template>
    <v-app>
      <v-container>
        <h1 class="green--text text-center">LOGIN HISTORY</h1>
        <v-divider></v-divider>
        <p class="green--text font-weight-bold text-h5">Login Details</p>
        <v-container fluid>
        <v-card >
      <v-card-title class="ml-8 green--text text-h5 font-weight-medium">
       History
        <v-spacer></v-spacer>
        <v-text-field class="red--text"
          v-model="search" 
          append-icon="mdi-magnify"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
      </v-card-title>
      <v-data-table  class="text-caption font-weight-medium px-8"
        :headers="headers"  :items="list" :search="search" 
        :sort-by="['Date', 'Time']" :sort-asc="[true, true]">
    
        <template v-slot:item.Type="{ item }">
      <v-chip :color="getColor(item.Type)" dark>
        {{ item.Type }}
      </v-chip>
    </template>
    </v-data-table>
    </v-card>
    </v-container>
  </v-container>
  </v-app>
  </template>

<script>

import history from "../../service/loginhistory/loginhistory"

export default {
  data(){
    return{
        search: "",
        headers: [
          { text: 'User info', align: 'start',sortable: false, value: 'UserId',},
          { text: 'Status', value: 'Type', align: 'center' },
          { text: 'Date', value: 'Date' , align: 'center'},
          { text: 'Time', value: 'Time' , align: 'center'},],
        id:'',
        pass:'',
        role:'',
        list:[],
    }
  },
  methods:{
    FetchingHistory(){
      history.LogedHistory()
        .then((response) =>{
        let data = response.data
            if (data.errMsg === "") {
              let loginHistoryArr = data.userDetailArr;
                loginHistoryArr.forEach(
                    val => this.list.push(val)
                )
            }      
          }).catch((error)=>{
            console.log("Error in getUserloginHistoryDetails_SERVICES :", error) 
        });
    },
  
    getColor(Type) {
        if (Type == 'Login'){return 'green'}
        else if (Type == 'Logout'){ return 'red'}
      },
  },

  mounted(){
    this.FetchingHistory()
  },
  };
</script>