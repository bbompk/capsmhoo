import axios from "axios";
import { ApiLoginResponse, ApiErrorResponse } from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { isResponseOk } from "../utils/ApiUtil";
import { addHoursToDate } from "../utils/DateTimeUtil";

export const login = async (email: string, password: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/login`;
    const axios_res = await axios.post(path, {
        email: email,
        password: password,
    });
    const res = axios_res.data as ApiLoginResponse;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    if(res.token != null) {
        localStorage.setItem("accessToken", res.token)
        localStorage.setItem("token_expires", addHoursToDate(new Date(), 1).toString())
    }
    return res;
}