import axios from "axios";
import { UserInterface } from "../interfaces/UserInterface";
import { ApiResponse, ApiErrorResponse } from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { isResponseOk } from "../utils/ApiUtil";

export const getAllUsers = async () => {
    const path = `${appConfig.BACKEND_BASE_URL}/user`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<UserInterface[]>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const getUserById = async (_id: string) => {
    const id = _id.replace(/"/g, '');
    const path = `${appConfig.BACKEND_BASE_URL}/user/${id}`;
    const axios_res = await axios.get(path);
    // const res = axios_res.data as ApiResponse<UserInterface>;
    const res: ApiResponse<UserInterface> = {
        code: axios_res.status.toString(),
        data: axios_res.data.data,
    };
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}
  
export const createUser = async (user: UserInterface) => {
    const path = `${appConfig.BACKEND_BASE_URL}/user`;
    const axios_res = await axios.post(path, user);
    const res = axios_res.data as ApiResponse<UserInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const updateUserById = async (id: string, user: UserInterface) => {
    const path = `${appConfig.BACKEND_BASE_URL}/user/${id}`;
    const axios_res = await axios.put(path, user);
    const res = axios_res.data as ApiResponse<UserInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const deleteUserById = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/user/${id}`;
    const axios_res = await axios.delete(path);
    const res = axios_res.data as ApiResponse<UserInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

