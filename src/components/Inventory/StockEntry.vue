<template>
    <v-app>
      <v-container>
        <h1 class="green--text text-center">STOCK ENTRY</h1>
        <v-divider></v-divider>
        <v-container fluid>
        <v-card >
      <v-card-title class="ml-0 green white--text text-h5 font-weight-medium">
        Stock Entry Page
        <v-spacer></v-spacer>
        <v-btn class="white green--text" @click="Addbtn"> + ADD</v-btn>
        <v-snackbar v-model="snackbar" class="" >{{ text }}
      <template v-slot:action="{ attrs }">
        <v-btn :color="colour" text v-bind="attrs" @click="snackbar = false">Close</v-btn></template>
        </v-snackbar>
        <v-form ref="form" v-model="valid" lazy-validation>
        <v-dialog v-model="demo" max-width="600px">
        <v-card >
            <v-card-title class="green white--text">Add Stock</v-card-title>
            <v-text-field outlined class="pa-2 mx-8" :rules="rules" v-model="newName" label="Medicine Name"></v-text-field>
            <v-text-field outlined class="pa-2 mx-8" :rules="rules" v-model="newBrand" label="Brand"></v-text-field>
            <v-btn block @click="newAdd" :disabled="!valid" class="green white--text">ADD</v-btn>
        </v-card>
        </v-dialog>
      </v-form>
      </v-card-title>
      <!-- <v-form ref="form" v-model="valid" lazy-validation> -->
        <v-form ref="form" lazy-validation>
        <v-row class="my-2 mx-2">
        <v-col cols="10" sm="4" md="2" lg="3">
          <v-select :items="list" v-model="SelectedMedicine"   @change="Medbrand" :rules="rules" outlined label="SelectedMedicine"></v-select>
        </v-col>
        <!-- <span v-show="val">{{ tempvariable }}</span> -->
        <v-col cols="10" sm="4" md="2" lg="3">
          <v-text-field v-model="brandlist" :rules="rules" readonly outlined label="Brand"></v-text-field>
        </v-col>
        <v-col cols="10" sm="4" md="2" lg="3">
          <v-text-field v-model="Qty" outlined type="number" :rules="numberRules1" required label="Qty"></v-text-field>
        </v-col>
        <v-col cols="10" sm="4" md="2" lg="3">
          <v-text-field v-model="price" outlined type="number" :rules="numberRules2" required label="price"></v-text-field>
        </v-col>
      </v-row>
      <!-- {{ total_compute }} -->
      <v-row  class="mx-1">
        <v-btn block class="green font-weight-bold white--text"  @click="change">UPDATE MEDICINES TO STOCK</v-btn>
        <!-- <v-btn block class="green font-weight-bold white--text" :disabled="!valid" @click="change">UPDATE MEDICINES TO STOCK</v-btn> -->
        <v-snackbar v-model="snackbar" class="" >{{ text }}
      <template v-slot:action="{ attrs }">
        <v-btn :color="colour" text v-bind="attrs" @click="snackbar = false">Close</v-btn></template>
        </v-snackbar>
      </v-row>
     </v-form> 
    </v-card>
    </v-container>
  </v-container>
  </v-app>
</template>

<script>
import fetchservice from "../../service/stockentry/fetchstock"
import updateservice from "../../service/stockentry/updatestock"
import stockservice from "../../service/stockentry/addstock"

export default {
  data(){
    return{
      total :0,
      SelectedMedicine: null,
      val:false,
      text:'',
      snackbar:false,
      colour:'green',
      newName:'',
      newBrand:'',
      temp:'',
      currentUser:this.$store.state.useractive.userid,
      valid: true,
    rules: [
      v => (!!v && v.length > 0 && v.length <= 16 && !v.includes(' ')) || 'Invalid characters',
      v => !/[!@#$%^&*(),.?":{}|<>]/g.test(v) || 'Invalid characters',],
        list:[],
        brandlist:'',
        Qty: null,
            numberRules1: [ value => !!value || 'Qty is required', value => (value >= 0) || 
            'Qty must be positive',
            v => /^-?\d+$/.test(v) || 'Enter a valid integer number'],
      price:null,
        numberRules2: [ value => !!value || 'Price is required', value => (value >= 0) || 
            'Price must be positive',
            v => /^-?\d+$/.test(v) || 'Enter a valid integer number'],
      demo: false,
    }
  },
  mounted() {
    this.Medname();
  },
  
//AddStockData
  methods:{
    Addbtn(){
      this.demo = true
     },
///checking for medname auto update
  Medname(){
    fetchservice.FetchStockData()
    .then((response) =>{
            if (response.data.status == "S") {
              this.temp=response.data.productsArr
              this.temp.forEach((product) => {
             this.medname = product.medicinename;
          this.list.push(this.medname);
          });
            }}
          ).catch((error)=>{
            console.log("Error in AddMEdStock_SERVICES :", error) 
        });
      },
      //extracting brands to the medicine name
  Medbrand(){
    this.temp.forEach((product) => {
        if (this.SelectedMedicine===product.medicinename){
          this.brandlist=product.brand
        }});     
      },
    //new stock adding method
    newAdd() {
      if(this.newName!=null && this.newBrand!=null){
        stockservice.AddStockData(this.newName,this.newBrand, this.currentUser)
        .then((response) =>{
            if (response.data.status == "S") {
                this.snackbar=true; 
                this.text="Stock Added succesfully";
                this.demo = false
            }else if(response.data.status == "E"){
                this.snackbar=true; 
                this.colour='red';
                this.text="Fail to Add stock";
            }else{
              this.snackbar=true;
                this.text="Entered medicine is already exist";
            }
          }
          ).catch((error)=>{
            console.log("Error in AddMEdStock_SERVICES :", error) 
        });
            }
          else{
                this.text="Please fill all the fields correctly";
                this.snackbar=true;
          }
          this.newName=null;
          this.newBrand=null;
    },
    //update the stock
    change(){
      if(this.SelectedMedicine!='' && this.brandlist!='' && this.Qty!=null && this.price!=null){
        this.valid=true
        updateservice.UpdateStockData(this.SelectedMedicine, this.brandlist, this.Qty, this.price, this.currentUser)
        .then((response) =>{
            if (response.data.status == "S") {
                this.snackbar=true; 
                this.text="Medicine Update succesfully";
            }else if(response.data.status == "E"){
                this.snackbar=true; 
                this.colour='red';
                this.text="Fail to Update";
            }else{
              this.snackbar=true;
                this.text="Invalid data";
            }
          }
          ).catch((error)=>{
            console.log("Error in UpdateMEdStock_SERVICES :", error) 
        }); 
  }
  else{
    this.text="Please choose the details";
    this.snackbar=true;
  }
    },
  },
  };
</script>
  





