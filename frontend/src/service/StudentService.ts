import axios from "axios";
import { StudentInterface } from "../interfaces/UserInterface";
import { ApiResponse, ApiErrorResponse } from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { isResponseOk, getEmptyHeaderWithBearerToken } from "../utils/ApiUtil";

export const getAllStudents = async () => {
    const headers = getEmptyHeaderWithBearerToken();
    const path = `${appConfig.BACKEND_BASE_URL}/student`;
    const axios_res = await axios.get(path, { headers });
    const res = axios_res.data as ApiResponse<StudentInterface[]>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const getStudentById = async (id: string) => {
    const headers = getEmptyHeaderWithBearerToken();
    const path = `${appConfig.BACKEND_BASE_URL}/student/${id}`;
    const axios_res = await axios.get(path, { headers });
    const res = axios_res.data as ApiResponse<StudentInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const getStudentByUserId = async (id: string) => {
    const headers = getEmptyHeaderWithBearerToken();
    const path = `${appConfig.BACKEND_BASE_URL}/student/userId/${id}`;
    const axios_res = await axios.get(path, { headers });
    const res = axios_res.data as ApiResponse<StudentInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const getAllStudentByTeamId = async (id: string) => {
    const headers = getEmptyHeaderWithBearerToken();
    const path = `${appConfig.BACKEND_BASE_URL}/student/teamId/${id}`;
    const axios_res = await axios.get(path, { headers });
    const res = axios_res.data as ApiResponse<StudentInterface[]>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}



export const createStudent = async (name: string, email: string, password: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/student`;
    const payload = {
        Name: name,
        Email: email,
        Password: password
    };
    const axios_res = await axios.post(path, payload);
    const res = axios_res.data as ApiResponse<StudentInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}
// export const createStudent = async (student: StudentInterface) => {
//     const path = `${appConfig.BACKEND_BASE_URL}/student`;
//     const axios_res = await axios.post(path, student);
//     const res = axios_res.data as ApiResponse<StudentInterface>;
//     if(!isResponseOk(res)) {
//         throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
//     }
//     return res;
// }

export const updateStudentById = async (id: string, student: Partial<StudentInterface>) => {
    const headers = getEmptyHeaderWithBearerToken();
    const path = `${appConfig.BACKEND_BASE_URL}/student/${id}`;
    const axios_res = await axios.put(path, student, { headers });
    const res = axios_res.data as ApiResponse<StudentInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const deleteStudentById = async (id: string) => {
    const headers = getEmptyHeaderWithBearerToken();
    const path = `${appConfig.BACKEND_BASE_URL}/student/${id}`;
    const axios_res = await axios.delete(path, { headers });
    const res = axios_res.data as ApiResponse<StudentInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const updateStudentTeamById = async (id: string, student: Partial<StudentInterface>) => {
    const headers = getEmptyHeaderWithBearerToken();
    const path = `${appConfig.BACKEND_BASE_URL}/student/updateTeam/${id}`;
    const axios_res = await axios.put(path, student, { headers });
    const res = axios_res.data as ApiResponse<StudentInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const alreadyHaveTeam = (student: StudentInterface) => {
    return (student.team_id !== "")
}