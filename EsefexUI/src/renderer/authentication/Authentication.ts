import { AxiosRequestConfig } from "axios";

const authHeaders: AxiosRequestConfig = {
    headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + localStorage.getItem('accessToken')
    }
};

export { authHeaders };