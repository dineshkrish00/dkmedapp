<template>
    <div>
        <sales></sales>
        <inventry></inventry>
    </div>
</template>

<script>
import sales from "../Dashboard/todaysales.vue";
import inventry from "../Dashboard/inventorydb.vue";

export default{
    components:{
        sales,
        inventry,
    },
}

</script>