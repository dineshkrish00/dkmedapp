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
    AddUserData(userId,  password,  role, currentUser){
        let obj = {
            userId: userId,
            password: password,
            role: role,
            currentUser: currentUser
        }
        return baseApiClient.post('/addUser', obj) 
    },
    
}