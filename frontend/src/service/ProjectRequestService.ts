import axios from "axios";
import { ProjectRequestInterface } from "../interfaces/ProjectInterface";
import { ApiResponse, ApiErrorResponse } from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { isResponseOk } from "../utils/ApiUtil";

export const getAllProjectRequests = async () => {
    const path = `${appConfig.BACKEND_BASE_URL}/projectRequest`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<ProjectRequestInterface[]>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const getProjectRequestById = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/projectRequest/${id}`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<ProjectRequestInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const createProjectRequest = async (projectRequest: ProjectRequestInterface) => {
    const path = `${appConfig.BACKEND_BASE_URL}/projectRequest`;
    const axios_res = await axios.post(path, projectRequest);
    const res = axios_res.data as ApiResponse<ProjectRequestInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const updateProjectRequestById = async (id: string, projectRequest: ProjectRequestInterface) => {
    const path = `${appConfig.BACKEND_BASE_URL}/projectRequest/${id}`;
    const axios_res = await axios.put(path, projectRequest);
    const res = axios_res.data as ApiResponse<ProjectRequestInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const deleteProjectRequestById = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/projectRequest/${id}`;
    const axios_res = await axios.delete(path);
    const res = axios_res.data as ApiResponse<ProjectRequestInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

