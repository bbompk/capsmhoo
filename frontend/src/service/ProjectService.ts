import axios from "axios";
import { ProjectInterface } from "../interfaces/ProjectInterface";
import { ApiResponse, ApiErrorResponse } from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { isResponseOk } from "../utils/ApiUtil";

export const getAllProjects = async () => {
    const path = `${appConfig.BACKEND_BASE_URL}/project`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<ProjectInterface[]>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const getProjectById = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/project/${id}`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<ProjectInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const createProject = async (project: ProjectInterface) => {
    const path = `${appConfig.BACKEND_BASE_URL}/project`;
    const axios_res = await axios.post(path, project);
    const res = axios_res.data as ApiResponse<ProjectInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const updateProjectById = async (id: string, project: ProjectInterface) => {
    const path = `${appConfig.BACKEND_BASE_URL}/project/${id}`;
    const axios_res = await axios.put(path, project);
    const res = axios_res.data as ApiResponse<ProjectInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const deleteProjectById = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/project/${id}`;
    const axios_res = await axios.delete(path);
    const res = axios_res.data as ApiResponse<ProjectInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

