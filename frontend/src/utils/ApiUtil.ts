import { ApiResponse } from "../interfaces/ApiResponseInterface";

export const getErrorMessage = (error: unknown): string => {
    let errorMessage = "";
    if (error instanceof Error) {
        errorMessage = error.message;
    } else if (typeof error === "string") {
        errorMessage = error;
    } else {
        errorMessage = "Unknown error";
    }
    return errorMessage;
}

export const isResponseOk = (response: ApiResponse): boolean => {
    if (response.code.slice(0, 2) === "20") {
        return true;
    } else {
        return false;
    }
}