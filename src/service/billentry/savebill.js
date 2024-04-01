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
    SaveBillData(
        medname,quantity,pricecart,billno,date,total_compute,net_compute,userid
    ){
        let obj = {
                medicinename:medname,
                quantity:quantity,
                unitprice:pricecart,
                billNo:billno,
                billDate:date,
                billAmount:total_compute,
                billGst: 18,
                netPrice:net_compute,
                currentUser:userid
        }
        return baseApiClient.post('/saveBillData', obj)
    },
    
}