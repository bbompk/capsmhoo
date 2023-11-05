import axios from "axios";
import { StudentInterface } from "../interfaces/UserInterface";
import { ApiResponse, ApiErrorResponse } from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { isResponseOk } from "../utils/ApiUtil";

export const getAllStudents = async () => {
    const path = `${appConfig.BACKEND_BASE_URL}/student`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<StudentInterface[]>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const getStudentById = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/student/${id}`;
    const axios_res = await axios.get(path);
    const res = axios_res.data as ApiResponse<StudentInterface>;
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

export const updateStudentById = async (id: string, student: StudentInterface) => {
    const path = `${appConfig.BACKEND_BASE_URL}/student/${id}`;
    const axios_res = await axios.put(path, student);
    const res = axios_res.data as ApiResponse<StudentInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

export const deleteStudentById = async (id: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/student/${id}`;
    const axios_res = await axios.delete(path);
    const res = axios_res.data as ApiResponse<StudentInterface>;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    return res;
}

