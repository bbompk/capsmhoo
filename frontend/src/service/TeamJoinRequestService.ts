import axios from "axios";
import { TeamJoinRequestInterface } from "../interfaces/TeamInterface";
import { ApiResponse, ApiErrorResponse } from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { isResponseOk } from "../utils/ApiUtil";

export const getAllTeamJoinRequests = async () => {
    const path = `${appConfig.BACKEND_BASE_URL}/team-join-request`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<TeamJoinRequestInterface[]>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const getTeamJoinRequestById = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/team-join-request/${id}`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<TeamJoinRequestInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const getAllTeamJoinRequestByTeamId = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/team-join-request/${id}`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<TeamJoinRequestInterface[]>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const createTeamJoinRequest = async (teamJoinRequest: TeamJoinRequestInterface) => {
    const path = `${appConfig.BACKEND_BASE_URL}/team-join-request`;
    const axios_res = await axios.post(path, teamJoinRequest);
    const res = axios_res.data as ApiResponse<TeamJoinRequestInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const updateTeamJoinRequestById = async (id: string, teamJoinRequest: TeamJoinRequestInterface) => {
    const path = `${appConfig.BACKEND_BASE_URL}/team-join-request/${id}`;
    const axios_res = await axios.put(path, teamJoinRequest);
    const res = axios_res.data as ApiResponse<TeamJoinRequestInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const deleteTeamJoinRequestById = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/team-join-request/${id}`;
    const axios_res = await axios.delete(path);
    const res = axios_res.data as ApiResponse<TeamJoinRequestInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

