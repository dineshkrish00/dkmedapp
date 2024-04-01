<template>
<div>
<v-btn class="green white--text" @click="preview">PREVIEW</v-btn>
    <v-dialog v-model="demo" max-width="550px">    
    <v-card>
        <v-card-title class="green white--text text-h5 font-weight-medium">Bill Preview</v-card-title>
        <h2 class="green--text font-weight-bold text-center py-4">MED MART</h2>
    <v-simple-table dense>
    <template>
      <thead>
        <tr>
          <th class="text-left"> Medicine Name </th>
          <th class="text-left"> QTy </th>
          <th class="text-left">Amount </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in medinfo" :key="item">
          <td>{{ item.medicinename }}</td>
          <td>{{ item.quantity }}</td>
          <td>{{ item.unitprice }}</td>
        </tr>
        <tr> <td> </td> <td>Total</td> <td>{{ total }}</td> </tr>
        <tr> <td> </td> <td>GST</td> <td>18%</td> </tr>
        <tr> <td> </td> <td>Netprice</td> <td>{{ netpayble }}</td> </tr>
      </tbody>
    </template>
  </v-simple-table>
  <v-card-text class="text-right">
     <v-btn class="green white--text mr-2 mt-4 py-2" @click="demo=false">Close</v-btn>
      <v-btn class="green white--text ml-4 mt-4 py-2" @click="demo=false">Print</v-btn> 
    </v-card-text>
    </v-card>
    </v-dialog>
</div>
</template>
<script>

export default {
    props: {
    medinfo:Array,
    total:Number,
    netpayble:Number
    },
    data(){
    return{
        demo:false,
    }
  },

  methods:{
    preview(){
        this.demo=true;
    }
  }
}
  </script>