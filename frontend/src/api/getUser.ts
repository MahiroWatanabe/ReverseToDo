import axios,{AxiosInstance} from "axios";

const authInstance:AxiosInstance = axios.create({
    baseURL:"http://localhost:8080",
    headers: {
        'Content-Type': 'application/json',
    },
})

export {authInstance} 