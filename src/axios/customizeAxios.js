import axios from "axios";

axios.defaults.withCredentials = true;

const instance = axios.create({
    baseURL: 'http://localhost:8080',
});

instance.interceptors.response.use(
    function (response) {
        return response.data;
    },
    function (error) {
        let res = {};
        if (error.response) {
            res.data = error.response.data;
            res.status = error.response.status;
            res.headers = error.response.headers;
        } else if (error.request) {
            console.log(error.request);
        } else {
            console.log('Error', error.message);
        }
        return Promise.reject(res); // Reject the promise with the error response
    }
);

export default instance;