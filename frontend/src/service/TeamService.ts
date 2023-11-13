import axios from "axios";
import {
  TeamCreateInterface,
  TeamInterface,
} from "../interfaces/TeamInterface";
import {
  ApiResponse,
  ApiErrorResponse,
} from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { getEmptyHeaderWithBearerToken, isResponseOk } from "../utils/ApiUtil";

export const getAllTeams = async () => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/team`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<TeamInterface[]>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const getTeamById = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/team/${id}`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<TeamInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const createTeam = async (teamCreate: TeamCreateInterface) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/team`;
  const axios_res = await axios.post(path, teamCreate, { headers });
  const res = axios_res.data as ApiResponse<TeamCreateInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const updateTeamById = async (id: string, team: TeamInterface) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/team/${id}`;
  const axios_res = await axios.put(path, team, { headers });
  const res = axios_res.data as ApiResponse<TeamInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const deleteTeamById = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/team/${id}`;
  const axios_res = await axios.delete(path, { headers });
  const res = axios_res.data as ApiResponse<TeamInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const getTeamByUserId = async (userId: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/team/user_id/${userId}`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<TeamInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};
