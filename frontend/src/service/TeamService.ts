import axios from "axios";
import { TeamInterface } from "../interfaces/TeamInterface";
import { ApiResponse, ApiErrorResponse } from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { isResponseOk } from "../utils/ApiUtil";

export const getAllTeams = async () => {
    const path = `${appConfig.BACKEND_BASE_URL}/team`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<TeamInterface[]>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const getTeamById = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/team/${id}`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<TeamInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const createTeam = async (team: TeamInterface) => {
    const path = `${appConfig.BACKEND_BASE_URL}/team`;
    const axios_res = await axios.post(path, team);
    const res = axios_res.data as ApiResponse<TeamInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const updateTeamById = async (id: string, team: TeamInterface) => {
    const path = `${appConfig.BACKEND_BASE_URL}/team/${id}`;
    const axios_res = await axios.put(path, team);
    const res = axios_res.data as ApiResponse<TeamInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const deleteTeamById = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/team/${id}`;
    const axios_res = await axios.delete(path);
    const res = axios_res.data as ApiResponse<TeamInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

