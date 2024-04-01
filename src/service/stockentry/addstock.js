import axios from 'axios'

const baseApiClient = axios.create(
    {
        baseURL : 'http://localhost:29091',
        withCredentials : false, // this is the default
        headers : {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        },
    }
)

export default{
    AddStockData(medicinename, brand,  currentUser){
        let obj = {
            medicinename: medicinename,
            brand: brand,
            currentUser:currentUser
        }
        return baseApiClient.post('/addStock', obj) 
    },
    
}