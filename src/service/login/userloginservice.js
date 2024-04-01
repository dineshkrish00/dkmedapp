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
    GetData(userId, password){
        let obj = {
            userId: userId,
            password: password
        }
        return baseApiClient.post('/getLogin', obj) 
    },
    LogoutData(userId){
        let detail={
            userId:userId
        }
        return baseApiClient.post('/getLogout', detail)
    }
}