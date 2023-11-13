import axios from "axios";
import { ProjectRequestInterface } from "../interfaces/ProjectInterface";
import {
  ApiResponse,
  ApiErrorResponse,
} from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { getEmptyHeaderWithBearerToken, isResponseOk } from "../utils/ApiUtil";

export const getAllProjectRequests = async () => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project-request`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<ProjectRequestInterface[]>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const getProjectRequestById = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project-request/${id}`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<ProjectRequestInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const getAllProjectRequestByProjectId = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project-request/projectid/${id}`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<ProjectRequestInterface[]>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const createProjectRequest = async (
  projectRequest: ProjectRequestInterface
) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project-request`;
  const axios_res = await axios.post(path, projectRequest, { headers });
  const res = axios_res.data as ApiResponse<ProjectRequestInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const acceptProjectRequestById = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project-request/accept/${id}`;
  const axios_res = await axios.post(path, { headers });
  const res = axios_res.data as ApiResponse<ProjectRequestInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const rejectProjectRequestById = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project-request/reject/${id}`;
  const axios_res = await axios.post(path, { headers });
  const res = axios_res.data as ApiResponse<ProjectRequestInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const updateProjectRequestById = async (
  id: string,
  projectRequest: ProjectRequestInterface
) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project-request/${id}`;
  const axios_res = await axios.put(path, projectRequest, { headers });
  const res = axios_res.data as ApiResponse<ProjectRequestInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const deleteProjectRequestById = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/project-request/${id}`;
  const axios_res = await axios.delete(path, { headers });
  const res = axios_res.data as ApiResponse<ProjectRequestInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};
