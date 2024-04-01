<!-- <template>
  <v-container>
    <h1 class="green--text text-center">Bill Entry Details</h1>
    <v-divider></v-divider>
    <p class="green--text font-weight-bold mx-3 text-h5">Bill Entry</p>
<v-row justify="center">
  <v-expansion-panels accordion>
    <v-expansion-panel class="py-0">
      <v-expansion-panel-header class="green py-2 text-h5 white--text">ITEMS</v-expansion-panel-header>
      <v-expansion-panel-content>
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
   <v-card >
       <v-card-title class="ml-8 mt-5 primary--text text-h5 font-weight-medium">
          <prepage :medinfo="this.medinfo" :total="this.total" :gst="this.gst" :netpayble="this.netpayble"></prepage>
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
           <span class="grey--text">{{ total }}</span>
          </p> <v-spacer></v-spacer>
          <p class="font-weight-bold text-caption black--text">GST:
           <span class="grey--text">{{ Math.floor(gst) }}</span>
          </p> <v-spacer></v-spacer>
          <p class="font-weight-bold text-caption black--text">Net Payable:
           <span class="grey--text">{{ netpayble }}</span>
          </p> <v-spacer></v-spacer>
       </v-card-title>
<v-card-title><v-spacer></v-spacer>
     <down :downval="this.downval"></down>
   </v-card-title>
 <v-card-title class="ml-8 primary--text text-h5 font-weight-medium">
   <v-text-field class="red--text"  append-icon="mdi-magnify" label="Search" single-line hide-details></v-text-field>
 </v-card-title>
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
import prepage from "../Biller/Preview.vue";
import down from "../Biller/download.vue";
export default {
components:{
 prepage,
 down,
},
data(){
  return{
     medname:'',   quantity:'',  date:'',  Quantity:'',  billno:'', total:0,  gst:0,
     arr : [],
     netpayble:0,  userid : this.$store.state.useractive.userid,  valid: true,
     Qty: null,
      numberRules: [ value => !!value || 'Qty is required', value => (value >= 0) || 
       'Qty must be positive',
       v => /^-?\d+$/.test(v) || 'Enter a valid integer number'],
     list:[],  temp:[],
     addcheck:false, text:'',  snackbar:false, colour: "green", amountdetails :0, unitdetails:0,
     upquantity:0, MedicineName:'', message:'',   log:false, medinfo:[], downval:[],
      headers: [
     { text: 'Medicine Name', align: 'start',sortable: false, value: 'MedicineName',},
     { text: 'Brand', value: 'brand' },
     { text: 'Quantity', value: 'quantity' },
     { text: 'Amount', value: 'unitprice' },],
     stocktemp:[],
     
  }
},
methods:{   
 //add & table
 change(){
   if(this.medname!='' && this.quantity!='' && this.quantity!=0){
//  let arr=this.$store.state.Stock;
    // let arr=[];
// this.$store.state.Stock.forEach((ele)=>{
//   arr.push({MedicineName:ele.MedicineName, quantity:ele.quantity, unitprice:ele.unitprice})
// })

 let med = this.arr.find((i) => i.MedicineName == this.medname);
for(let cn of this.arr){
 if(cn.MedicineName==this.medname)
 {
  console.log(cn.quantity);
  console.log(this.quantity);
// if(cn.quantity!=0){
   if(Number(cn.quantity)>=this.quantity && this.quantity>=0){
    cn.quantity=Number(cn.quantity)-this.quantity;
     this.upquantity=cn.quantity-this.quantity;
     console.log("cn"+cn.quantity);
      console.log("this"+this.quantity);
     console.log("up"+this.upquantity);
    //  this.quantity>=this.upquantity && 
        if(cn.quantity>=0){
        this.snackbar=true;
        this.text="Added cart Succesfully";
        this.updated(med);
        this.stockpush();
        }else{
          this.snackbar=true;
        this.text=" Quantity"; 
        }
   }
   else{
     this.snackbar=true;
     this.text="Please add valid Quantity";
   }
//  }
}
}
}
else{
     this.snackbar=true;
     this.text="please fill All details";
}
this.medname='';
this.quantity=null;
},//stock view table
stockpush(){
for(let i of this.arr){
if(i.MedicineName==this.medname){
this.stocktemp.push({MedicineName:i.MedicineName, quantity:this.upquantity, unitprice:i.unitprice});
}
}
},
updated(med){        
   this.temp.push({MedicineName: med.MedicineName, brand:'', quantity: this.quantity,
   unitprice: 0,})
   for (let i of this.temp) {
   for (let j of this.$store.state.MedicineMaster) {
     if (i.MedicineName == j.MedicineName) {
         i.brand = j.Brand;
     }
   }
 }
 for(let a of this.arr){
   for (let i of this.temp) {
     if(a.MedicineName==i.MedicineName){
     i.unitprice = i.quantity * a.unitprice;
     }
   }  
}
     this.medinfo=this.temp;
     this.text1();    
     this.addcheck=true;    
for(let val of this.medinfo){
 this.downval.push({MedicineName:val.MedicineName, brand:val.brand, 
   quantity:val.quantity, unitprice:val.unitprice});
 }
},
text1(){
let sum=0;
for(let val of this.medinfo){
   sum=sum+val.unitprice;
   this.total=sum;
   this.billno=1000;
   this.billno++;
   this.gst=Number(this.total)*0.18;
   this.netpayble=Number(this.total)+Number(this.gst);
   this.date=this.convertDate(new Date());
}
},
//converting date standard
convertDate(date) {
   var dateObj = new Date(date);
   var month = (dateObj.getMonth() + 1).toString().padStart(2, '0'); 
   var day = dateObj.getDate().toString().padStart(2, '0');
   var year = dateObj.getFullYear();
   return day + '/' + month + '/' + year;
},
//check the array
save(){
if(this.addcheck){
   this.master();
   this.stockview();
   this.snackbar=true;
   this.text="Saved Succesfully";
   this.addcheck=false;
}
else{
 this.snackbar=true;
 this.text="First add the Medicines";
}
 },
 //stock view table
stockview(){
for(let i of this.$store.state.Stock){
for(let j of this.stocktemp){
 if(i.MedicineName==j.MedicineName){
   i.quantity=j.quantity;
 }
 else{
   this.snackbar=true;
   this.text="qunantiy not added";
 }
}
}
this.details();
},
master(){
   this.$store.state.billMaster.push({Billno:this.billno, Billdate: this.date, Billamount:this.total, Billgst:this.gst, Netprice: this.netpayble, UserId:this.userid})
     for(let a of this.$store.state.billMaster){
       console.log(a.Billno + a.Billdate + a.Billamount + a.Billgst + a.Netprice + a.UserId)
     }
 },
details(){
for(let a of this.$store.state.Stock){
   for (let i of this.temp) {
     if(a.MedicineName==i.MedicineName){
       this.unitdetails=a.unitprice;
     this.amountdetails = i.quantity * a.unitprice;
     this.$store.state.billDetails.push({Billno:this.billno, MedicineName:this.medname, quantity:this.quantity, unitprice:this.unitdetails, amount:this.amountdetails})
     }
   }  
}
}
},
mounted(){
let a=this.$store.state.MedicineMaster;
 a.forEach((element) => {
   this.list.push(element.MedicineName);
});

this.$store.state.Stock.forEach((ele)=>{
  this.arr.push({MedicineName:ele.MedicineName, quantity:ele.quantity, unitprice:ele.unitprice})
})
} 
}
</script> -->























<!-- <template>
  <v-container>
    <h1 class="green--text text-center">Bill Entry Details</h1>
    <v-divider></v-divider>
    <p class="green--text font-weight-bold mx-3 text-h5">Bill Entry</p>
<v-row justify="center">
  <v-expansion-panels accordion>
    <v-expansion-panel class="py-0">
      <v-expansion-panel-header class="green py-2 text-h5 white--text">ITEMS</v-expansion-panel-header>
      <v-expansion-panel-content>
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
   <v-card >
       <v-card-title class="ml-8 mt-5 primary--text text-h5 font-weight-medium">
          <prepage :medinfo="this.medinfo" :total="this.total" :gst="this.gst" :netpayble="this.netpayble"></prepage>
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
           <span class="grey--text">{{ total }}</span>
          </p> <v-spacer></v-spacer>
          <p class="font-weight-bold text-caption black--text">GST:
           <span class="grey--text">{{ Math.floor(gst) }}</span>
          </p> <v-spacer></v-spacer>
          <p class="font-weight-bold text-caption black--text">Net Payable:
           <span class="grey--text">{{ netpayble }}</span>
          </p> <v-spacer></v-spacer>
       </v-card-title>
<v-card-title><v-spacer></v-spacer>
     <down :downval="this.downval"></down>
   </v-card-title>
 <v-card-title class="ml-8 primary--text text-h5 font-weight-medium">
   <v-text-field class="red--text"  append-icon="mdi-magnify" label="Search" single-line hide-details></v-text-field>
 </v-card-title>
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
import prepage from "../Biller/Preview.vue";
import down from "../Biller/download.vue";
export default {
components:{
 prepage,
 down,
},
data(){
  return{
     medname:'',   quantity:'',  date:'',  Quantity:'',  billno:'', total:0,  gst:0,
     netpayble:0,  userid : this.$store.state.useractive.userid,  valid: true,
     Qty: null,
      numberRules: [ value => !!value || 'Qty is required', value => (value >= 0) || 
       'Qty must be positive',
       v => /^-?\d+$/.test(v) || 'Enter a valid integer number'],
     list:[],  temp:[],
     addcheck:false, text:'',  snackbar:false, colour: "green", amountdetails :0, unitdetails:0,
     upquantity:0, MedicineName:'', message:'',   log:false, medinfo:[], downval:[],
      headers: [
     { text: 'Medicine Name', align: 'start',sortable: false, value: 'MedicineName',},
     { text: 'Brand', value: 'brand' },
     { text: 'Quantity', value: 'quantity' },
     { text: 'Amount', value: 'unitprice' },],
     stocktemp:[],
     
  }
},
methods:{   
 //add & table
 change(){
   if(this.medname!='' && this.quantity!='' && this.quantity!=0){

//  let arr=this.$store.state.Stock;
    let arr=[];
this.$store.state.Stock.forEach((ele)=>{
  arr.push({MedicineName:ele.MedicineName, quantity:ele.upquantity, unitprice:ele.unitprice})
})

 let med = arr.find((i) => i.MedicineName == this.medname);
for(let cn of arr){
 if(cn.MedicineName==this.medname)
 {
   if(cn.quantity>=this.quantity && this.quantity>=0){
     this.upquantity=cn.quantity-this.quantity;
     console.log(this.upquantity);
     this.snackbar=true;
     this.text="Added cart Succesfully";
     this.updated(med, arr);
     this.stockpush(arr);
   }
   else{
     this.snackbar=true;
     this.text="Please add valid Quantity";
   }
 }
}
}
else{
     this.snackbar=true;
     this.text="please fill All details";
}
this.medname='';
this.quantity=null;
},//stock view table
stockpush(arr){
for(let i of arr){
if(i.MedicineName==this.medname){
this.stocktemp.push({MedicineName:i.MedicineName, quantity:this.upquantity, unitprice:i.unitprice});
}
}
},
updated(med, arr){        
   this.temp.push({MedicineName: med.MedicineName, brand:'', quantity: this.quantity,
   unitprice: 0,})
   for (let i of this.temp) {
   for (let j of this.$store.state.MedicineMaster) {
     if (i.MedicineName == j.MedicineName) {
         i.brand = j.Brand;
     }
   }
 }
 for(let a of arr){
   for (let i of this.temp) {
     if(a.MedicineName==i.MedicineName){
     i.unitprice = i.quantity * a.unitprice;
     }
   }  
}
     this.medinfo=this.temp;
     this.text1();    
     this.addcheck=true;    
for(let val of this.medinfo){
 this.downval.push({MedicineName:val.MedicineName, brand:val.brand, 
   quantity:val.quantity, unitprice:val.unitprice});
 }
},
text1(){
let sum=0;
for(let val of this.medinfo){
   sum=sum+val.unitprice;
   this.total=sum;
   this.billno=1000;
   this.billno++;
   this.gst=Number(this.total)*0.18;
   this.netpayble=Number(this.total)+Number(this.gst);
   this.date=this.convertDate(new Date());
}
},
//converting date standard
convertDate(date) {
   var dateObj = new Date(date);
   var month = (dateObj.getMonth() + 1).toString().padStart(2, '0'); 
   var day = dateObj.getDate().toString().padStart(2, '0');
   var year = dateObj.getFullYear();
   return day + '/' + month + '/' + year;
},
//check the array
save(){
if(this.addcheck){
   this.master();
   this.stockview();
   this.snackbar=true;
   this.text="Saved Succesfully";
   this.addcheck=false;
}
else{
 this.snackbar=true;
 this.text="First add the Medicines";
}
 },
 //stock view table
stockview(){
for(let i of this.$store.state.Stock){
for(let j of this.stocktemp){
 if(i.MedicineName==j.MedicineName){
   i.quantity=j.quantity;
 }
 else{
   this.snackbar=true;
   this.text="qunantiy not added";
 }
}
}
this.details();
},
master(){
   this.$store.state.billMaster.push({Billno:this.billno, Billdate: this.date, Billamount:this.total, Billgst:this.gst, Netprice: this.netpayble, UserId:this.userid})
     for(let a of this.$store.state.billMaster){
       console.log(a.Billno + a.Billdate + a.Billamount + a.Billgst + a.Netprice + a.UserId)
     }
 },
details(){
for(let a of this.$store.state.Stock){
   for (let i of this.temp) {
     if(a.MedicineName==i.MedicineName){
       this.unitdetails=a.unitprice;
     this.amountdetails = i.quantity * a.unitprice;
     this.$store.state.billDetails.push({Billno:this.billno, MedicineName:this.medname, quantity:this.quantity, unitprice:this.unitdetails, amount:this.amountdetails})
     }
   }  
}
}
},
mounted(){
let a=this.$store.state.MedicineMaster;
 a.forEach((element) => {
   this.list.push(element.MedicineName);
});
} 
}
</script> -->