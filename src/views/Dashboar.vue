<template>
    <div>
        <dashboard1 v-if="role=='Admin' ||  role=='Inventry'"></dashboard1>
        <dashboard2 v-if="role=='Manager' ||  role=='Biller'"></dashboard2>
    </div>
</template>

<script>
import dashboard1 from "../components/Dashboard/Dashboard.vue";
import dashboard2 from "../components/Dashboard/ManagerDB.vue";

export default{
  name: "dashboardname",
components:{
    dashboard1,
    dashboard2,
},
data(){
    return{
        role:this.$store.state.useractive.role,
    }
}
}
</script>