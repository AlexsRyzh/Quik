import axios from "axios";

const $api = axios.create({
    baseURL: 'http://localhost:3000',
});

$api.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem('access_token')}`
    return config
})

$api.interceptors.response.use((config) => {
    return config
}, async (error) => {
    const originalRequest = error.config
    console.log(error)
    if (error.response?.status == 401) {
        try {
            const res = await axios.get(
                `http://localhost:3000/auth/refresh}`,
                {headers: {Authorization: `Bearer ${localStorage.getItem('refresh_token')}`}}
            )
            localStorage.setItem('access_token', res.data['access_token'])
            localStorage.setItem('refresh_token', res.data['refresh_token'])
            window.localStorage.setItem('user_id', res.data['user_id'])
            return $api.request(originalRequest)
        } catch (e) {

            localStorage.removeItem('access_token')
            localStorage.removeItem('refresh_token')
            window.localStorage.removeItem('user_id')
            console.log("Не авторизован")
        }
    }
})

export default $api