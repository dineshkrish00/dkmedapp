<template>
    <div>
        <stockentry></stockentry>
    </div>
</template>

<script>
import stockentry from "../components/Inventory/StockEntry.vue";

export default{
  name: "stockentryname",

    components:{
        stockentry,
    },

}

</script>