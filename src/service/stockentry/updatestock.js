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
    UpdateStockData(medicinename, brand, quantity, unitprice,  currentUser){
        let obj = {
            medicinename:   medicinename,
            brand:          brand,
            quantity:       quantity,
            unitprice:       unitprice,
            currentUser:    currentUser
        }
        return baseApiClient.post('/updateStock', obj) 
    },
    
}