import { UserInterface } from "./UserInterface";

export interface ApiResponse<T = any> {
    code: string;
    data?: T;
    message?: string;
    error?: string;
}

export class ApiErrorResponse extends Error {
    code: string;
    message: string;

    constructor(code: string, message: string) {
        super(message);
        this.code = code;
        this.message = message;
    }
}

export interface ApiLoginResponse extends ApiResponse{
    code: string;
    message?: string;
    data?: UserInterface;
    token?: string;
    errors?: string;
}