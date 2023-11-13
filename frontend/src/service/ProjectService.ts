import axios from "axios";
import { ProjectInterface } from "../interfaces/ProjectInterface";
import {
  ApiResponse,
  ApiErrorResponse,
} from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { isResponseOk, getEmptyHeaderWithBearerToken } from "../utils/ApiUtil";

export const getAllProjects = async () => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<ProjectInterface[]>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const getProjectById = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project/${id}`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<ProjectInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const getProjectByTeamId = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project/teamId/${id}`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<ProjectInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const getProjectByProfessorId = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project/professorId/${id}`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<ProjectInterface[]>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const createProject = async (project: ProjectInterface) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project`;
  const axios_res = await axios.post(path, project, { headers });
  const res = axios_res.data as ApiResponse<ProjectInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const updateProjectById = async (
  id: string,
  project: ProjectInterface
) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project/${id}`;
  const axios_res = await axios.put(path, project, { headers });
  const res = axios_res.data as ApiResponse<ProjectInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const deleteProjectById = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project/${id}`;
  const axios_res = await axios.delete(path, { headers });
  const res = axios_res.data as ApiResponse<ProjectInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};
