import axios from "axios";
import { ApiLoginResponse, ApiErrorResponse, ApiResponse } from "../interfaces/ApiResponseInterface";
import appConfig from "../configs/config";
import { isResponseOk } from "../utils/ApiUtil";
import { addHoursToDate } from "../utils/DateTimeUtil";
import { StudentInterface, ProfessorInterface } from "../interfaces/UserInterface";

export const login = async (email: string, password: string) => {
    const path = `${appConfig.BACKEND_BASE_URL}/auth/login`;
    let axios_res = await axios.post(path, {
        email: email,
        password: password,
    });
    const res = axios_res.data as ApiLoginResponse;
    if(!isResponseOk(res)) {
        throw new ApiErrorResponse(res.code, res.error ?? "Unknown error");
    }
    if(res.data?.token != null) {
        sessionStorage.setItem("accessToken", res.data.token)
        sessionStorage.setItem("token_expires", addHoursToDate(new Date(), 1).toString())

        sessionStorage.setItem("userId", res.data.user.id)
        sessionStorage.setItem("role", res.data.user.role)
    }

    const studentPath = `${appConfig.BACKEND_BASE_URL}/student/userId/${res.data?.user.id}`;
    const professorPath = `${appConfig.BACKEND_BASE_URL}/professor/userId/${res.data?.user.id}`;

    switch(res.data?.user.role) {
        case "Student": {
            axios_res = await axios.get(studentPath);
            let studentRes = axios_res.data as ApiResponse<StudentInterface>;
            if(!isResponseOk(studentRes)) {
                throw new ApiErrorResponse(studentRes.code, studentRes.error ?? "Unknown error");
            }
            sessionStorage.setItem("studentId", studentRes.data?.id ?? "");
            break;
        }
        case "Professor": {
            axios_res = await axios.get(professorPath);
            let professorRes = axios_res.data as ApiResponse<ProfessorInterface>;
            if(!isResponseOk(professorRes)) {
                throw new ApiErrorResponse(professorRes.code, professorRes.error ?? "Unknown error");
            }
            sessionStorage.setItem("professorId", professorRes.data?.id ?? "");
            break;
        }
        default:
            throw new ApiErrorResponse("unknown_role", "Unknown role");
    }
    return res;
}