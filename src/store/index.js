import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {



   useractive:[{
      userid:null,
      role:null,
   }],

    Login:[
      {userid:'FT104', pass:'12345', role:'Biller' },
      {userid:'FT102', pass:'12345', role:'Manager' },
      {userid:'FT101', pass:'12345', role:'Admin' },
      {userid:'FT103', pass:'12345', role:'Inventry' },
    ],
    log_history:[],

    MedicineMaster:[
       {MedicineName:'vitaminA', Brand:'A-brand'},
       {MedicineName:'vitaminB', Brand:'B-brand'},
       {MedicineName:'vitaminC', Brand:'C-brand'}
    ],
    Stock:[
       {MedicineName:'vitaminA', quantity:'100', unitprice:'5'},
       {MedicineName:'vitaminB', quantity:'200', unitprice:'10'},
       {MedicineName:'vitaminC', quantity:'300', unitprice:'15'},
    ],
    billMaster:[
       {Billno:'121', Billdate:'07/12/2023', Billamount:'20', Billgst:'20', Netprice:'20', UserId:"id"},
       {Billno:'122', Billdate:'08/12/2023', Billamount:'20', Billgst:'20', Netprice:'20', UserId:"id"},
       {Billno:'122', Billdate:'08/12/2023', Billamount:'20', Billgst:'20', Netprice:'20', UserId:"id"},
    ],

    billDetails:[
       {Billno:'121', MedicineName:'vitaminA', quantity:'100', unitprice:'5', amount:'15'},
       {Billno:'122', MedicineName:'vitaminB', quantity:'1200', unitprice:'6', amount:'20'},
       {Billno:'122', MedicineName:'vitaminC', quantity:'1200', unitprice:'6', amount:'20'},
       
    ]
},
mutations: {
   updateuser(state,value){
      state.useractive.role=value;
   },

 addMedicine(state, medicine) {
    state.MedicineMaster.push(medicine);
   },

   stockUpdate(state, values) {
    state.Stock.push(values);
   },
   
//   addUserID(state, value) {
//     state.Login.push({userid: value,});
//   },
//   addPass(state, value) {
//     state.Login.push({pass: value});
//   },
//   addRole(state, value) {
//     state.Login.push({role: value});
//   },
},
actions: {},
modules: {},
});

