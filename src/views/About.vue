download(){
    // Your CSV data (replace this with your actual data)
 const csvData = this.stock

 // Convert the data array to a CSV string
 const csvString = Papa.unparse(csvData);

 // Create a Blob from the CSV string
 const blob = new Blob([csvString], { type: 'text/csv' });

 // Create a temporary link element
 const link = document.createElement('a');
 link.href = window.URL.createObjectURL(blob);

 // Set the filename for the download
 link.download = 'stock.csv';

 // Append the link to the document
 document.body.appendChild(link);

 // Trigger the click event to initiate the download
 link.click();

 // Remove the link from the document
 document.body.removeChild(link);
 this.snackbar=true
 }
 
 
 
 import Papa from 'papaparse';