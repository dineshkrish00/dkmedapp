<template>
<div>
    <v-container>
        <v-btn class="green white--text" @click="downclick" >DOWNLOAD</v-btn> 
        <v-snackbar v-model="snackbar" class="" >{{ text }}
      <template v-slot:action="{ attrs }">
    <v-btn :color="colour" text v-bind="attrs" @click="snackbar = false">Close</v-btn></template>
    </v-snackbar>
</v-container>
</div>
</template>

<script>
import Papa from "papaparse";

export default{
    props:{
        downval:Array,
    },
    data(){
        return{
            snackbar:false,
            text:'',
            colour: "red",
        }
    },
    methods:{
    download(){
            const csvData = this.downval;
            const csvString = Papa.unparse(csvData);
            const blob = new Blob([csvString], { type: 'text/csv' });
            const link = document.createElement('a');
            link.href = window.URL.createObjectURL(blob);
            link.download = 'Billdetails.csv';
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
            this.snackbar=true
    },
    downclick(){
        if(this.downval!=''){
            this.download();
        }
        else{
            this.snackbar=true;
            this.text="Without data file doesn't download";
            this.colour="green";
        }
    }
    }
}

</script>