import axios from "axios";
import { ProfessorInterface } from "../interfaces/UserInterface";
import {
  ApiResponse,
  ApiErrorResponse,
} from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { getEmptyHeaderWithBearerToken, isResponseOk } from "../utils/ApiUtil";

export const getAllProfessors = async () => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/professor`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<ProfessorInterface[]>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const getProfessorById = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/professor/${id}`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<ProfessorInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const getProfessorByUserId = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/professor/userId/${id}`;
  const axios_res = await axios.get(path, { headers });
  const res = axios_res.data as ApiResponse<ProfessorInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const createProfessor = async (
  name: string,
  email: string,
  password: string,
  profile: string
) => {
  const path = `${appConfig.BACKEND_BASE_URL}/professor`;
  const payload = {
    Name: name,
    Email: email,
    Password: password,
    Profile: profile,
  };
  const axios_res = await axios.post(path, payload);
  const res = axios_res.data as ApiResponse<ProfessorInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};
// export const createProfessor = async (professor: ProfessorInterface) => {
//     const path = `${appConfig.BACKEND_BASE_URL}/professor`;
//     const axios_res = await axios.post(path, professor);
//     const res = axios_res.data as ApiResponse<ProfessorInterface>;
//     if(!isResponseOk(res)) {
//         throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
//     }
//     return res;
// }

export const updateProfessorById = async (
  id: string,
  professor: Partial<ProfessorInterface>
) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/professor/${id}`;
  const axios_res = await axios.put(path, professor, { headers });
  const res = axios_res.data as ApiResponse<ProfessorInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};

export const deleteProfessorById = async (id: string) => {
  const headers = getEmptyHeaderWithBearerToken();
  const path = `${appConfig.BACKEND_BASE_URL}/professor/${id}`;
  const axios_res = await axios.delete(path, { headers });
  const res = axios_res.data as ApiResponse<ProfessorInterface>;
  if (!isResponseOk(res)) {
    throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
  }
  return res;
};
