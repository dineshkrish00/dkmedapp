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
    AddBillData(MedicineName, Quantitiy, User){
        let obj = {
            medicinename: MedicineName,
            quantity: Quantitiy,
            currentUser: User
        }
        return baseApiClient.post('/addBillData', obj)
    },
    
}