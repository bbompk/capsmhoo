import axios from "axios";
import { NotificationInterface } from "../interfaces/NotiInterface";
import { ApiResponse, ApiErrorResponse } from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { isResponseOk } from "../utils/ApiUtil";

export const getAllNotificationsByRoleId = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/notification/${id}`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<NotificationInterface[]>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const readAllNotificationsByRoleId = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/notification/${id}`;
    const axios_res = await axios.post(path);
    const res = axios_res.data as ApiResponse<null>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}