<template>
  <v-container>
    <h1 class="green--text text-center">Bill Entry Details</h1>
    <v-divider></v-divider>
    <p class="green--text font-weight-bold mx-3 text-h5">Bill Entry</p>
<v-row justify="center">
  <v-expansion-panels accordion>
    <v-expansion-panel class="py-0">
      <v-expansion-panel-header class="green py-2 text-h5 white--text">ITEMS</v-expansion-panel-header>
      <v-expansion-panel-content>
        <!-- 1st part ITEMs added section -->
       <v-form ref="form" v-model="valid" lazy-validation>
            <v-container>
              <v-row>
               <v-col cols="4">  
                    <v-select class="mx-15" label="MedicineName" :items="list"  v-model="medname" outlined></v-select>
               </v-col>                    
               <v-col cols="4">
                  <v-text-field class="mx-15" type="number" label="Quantity" :rules="numberRules" v-model="quantity" outlined></v-text-field>
               </v-col>
               <v-col cols="4">
                    <v-btn @click="change" :disabled="!valid" class="green white--text mx-15" x-large outlined>ADD</v-btn>
               </v-col>
              </v-row>
            </v-container>
          </v-form>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</v-row>
<v-container>
   <v-divider></v-divider>
   <v-container fluid>
    <!-- 2nd part bill details section -->
   <v-card >
       <v-card-title class="ml-8 mt-5 primary--text text-h5 font-weight-medium">
          <prepage :medinfo="this.medinfo" :total="this.total_compute" :netpayble="this.net_compute"></prepage>
          <v-spacer></v-spacer>
          <v-btn class="green white--text" @click="save">SAVE</v-btn> 
          <v-spacer></v-spacer>
          <p class="font-weight-bold text-caption black--text">Bill NO:
           <span class="grey--text">{{ billno }}</span>
          </p> <v-spacer></v-spacer>
          <p class="font-weight-bold text-caption black--text">DATE:
           <span class="grey--text">{{ date }}</span>
          </p> <v-spacer></v-spacer>
          <p class="font-weight-bold text-caption black--text">TOTAL:
           <span class="grey--text">{{ total_compute }}</span>
          </p> <v-spacer></v-spacer>
          <p class="font-weight-bold text-caption black--text">GST:
           <span class="grey--text">{{ 18 }}%</span>
          </p> <v-spacer></v-spacer>
          <p class="font-weight-bold text-caption black--text">Net Payable:
            <span class="grey--text">{{ net_compute }}</span>
          </p> <v-spacer></v-spacer>
       </v-card-title>
<v-card-title><v-spacer></v-spacer>
     <down :downval="this.downval"></down>
   </v-card-title>
 <v-card-title class="ml-8 primary--text text-h5 font-weight-medium">
   <v-text-field class="red--text"  append-icon="mdi-magnify" label="Search" single-line hide-details></v-text-field>
 </v-card-title>
 <!-- 3rd part bill data table section -->
 <v-data-table  class="text-caption font-weight-medium px-8" :headers="headers" :items="medinfo"></v-data-table>
</v-card>
</v-container>
  </v-container>
  <v-snackbar v-model="snackbar" class="" >{{ text }}
 <template v-slot:action="{ attrs }">
   <v-btn :color="colour" text v-bind="attrs" @click="snackbar = false">Close</v-btn></template>
   </v-snackbar>  
</v-container>   
</template>

<script>
import mednameservice from  "../../service/billentry/fetchmedname"
import addbillservice from "../../service/billentry/addbill"
import savebill from "../../service/billentry/savebill"

import prepage from "../Biller/Preview.vue";
import down from "../Biller/download.vue";
export default {
components:{
 prepage,
 down,
},
data(){
  return{
      date:'yyyy-mm-dd',  billno:'1000', 
      valid:false,
     netpayble:0,  userid : this.$store.state.useractive.userid,  
      numberRules: [ value => !!value || 'Qty is required', value => (value >= 0) || 
       'Qty must be positive',
       v => /^-?\d+$/.test(v) || 'Enter a valid integer number'],
     text:'',  snackbar:false, colour: "green", 
     amountdetails :0, unitdetails:0,
     upquantity:0, MedicineName:'', message:'',
     log:false, downval:[],
      headers: [
     { text: 'Medicine Name', align: 'start',sortable: false, value: 'medicinename',},
     { text: 'Brand', value: 'brand' },
     { text: 'Quantity', value: 'quantity' },
     { text: 'Unit Price', value: 'unitprice' },],
    //  stocktemp:[],
    //  arr:[],
     list:[], 
     medname:'',   quantity:0,
     medinfo:[],
     cart:'',
     mnamecart:'',
     brandcart:'',
     medidcart:'',
     quancart:'',
     pricecart: 0
  }
},
computed:{
  total_compute(){
    let total = this.medinfo.length > 0 ? this.medinfo.map( obj => 
    parseInt(obj.quantity) * obj.unitprice).
    reduce( ( num1, num2 )=> num1 + num2, 0 ) : 0;
      return total
  },
  net_compute(){
    let netPay = this.total_compute + ( this.total_compute * 18 / 100 )
      return parseFloat(netPay.toFixed(2))
  },
},
methods:{  
  //3rd method 
  save(){ 
    savebill.SaveBillData(this.medname,parseInt(this.quantity), this.pricecart,parseInt(this.billno),this.date,
    this.total_compute,this.net_compute,this.userid)
    .then((response) =>{
            if (response.data.status == "S") {
              this.snackbar=true; 
              this.text="Bill Added succesfully";
              console.log("Table DB :", response.data);
              this.date=this.convertDate(new Date())
      }else{
        this.snackbar=true;
        this.text="Fail to Add Bill";
        }
      }
      ).catch((error)=>{
        console.log("Error in AddBill_SERVICES :", error) 
    });
  },

//for date
convertDate(date) {
var dateObj = new Date(date);
var month = (dateObj.getMonth() + 1).toString().padStart(2, '0'); 
var day = dateObj.getDate().toString().padStart(2, '0');
var year = dateObj.getFullYear();
return year +`-`+ month +`-`+day  ;
},
 //2nd method for  add & CHecking bill 
 change(){
  if(this.medname!='' && this.quantity!='' && this.quantity!=0){
  addbillservice.AddBillData(this.medname, parseInt(this.quantity),this.userid)
  .then((response) =>{
            if (response.data.status == "S") {
              // this.billno++
             this.cart=response.data.addBillArr
              this.cart.forEach(out=>{
                this.quancart=parseInt(out.quantity) 
                this.pricecart= parseInt(out.unitprice)
    const index = this.medinfo.findIndex(medicine => medicine.medicinename === this.medname);
      if (index !== -1) {
        this.medinfo[index].quantity= Number(this.medinfo[index].quantity)+Number(this.quantity);
      } else {
        // Medicine not found, insert new medicine detail
        this.medinfo.push({
          medicinename:out.medicinename, 
          brand:out.brand,
          quantity:this.quantity,
          unitprice:this.pricecart
        });
        console.log("medinfo array", this.medinfo);
      } 
        })
          this.snackbar=true; 
          this.text="Medicine Added to cart succesfully";
          console.log("Table DB :", response.data);
          this.date=this.convertDate(new Date())
      }else{
        this.snackbar=true;
        this.text="Quantity not Exceed";
        }
      }
      ).catch((error)=>{
        console.log("Error in AddBill_SERVICES :", error) 
    });
        
      }
      else{
     this.snackbar=true;
     this.text="please fill All details";
   }
},

  //method 1 for fetching med name 
  FetchMedname(){
    mednameservice.FetchMedData()
    .then((response) =>{
            if (response.data.status == "S") {
              this.billno=response.data.billNo
          this.medtemp=response.data.productsArr
          this.medtemp.forEach((product) => {
          this.medname = product.medicinename;
          this.list.push(this.medname);
          });
            }else{
              console.log("Fetching data failed");
            }}
          ).catch((error)=>{
            console.log("Error in_ Billentry vue :", error) 
        });
    }
},
mounted(){
  this.FetchMedname()
} 
}
</script>

