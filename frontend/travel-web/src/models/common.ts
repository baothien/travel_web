import { User } from "./user";

export interface LoginParams{
    user_name: string;
    password: string;
}

export interface RegisterParams{
    email: string,
    password: string,
    avatar: "",
    date_of_birth: "",
    full_name: "",
    phone: ""
}

export interface ResponseMessage<T>{
    message: string;
    data: T;
}

export interface Response<T>{
    data: T
}

export interface ImageUploaded{
    full_path: string;
    id: string;
    name: string;
    path: string;
    size: number;
    type: string;
}

export interface PaginateParams{
    page: number;
    limit: number;
}

export interface PaginateData<T>{
    limit: number;
    page: number;
    sort: string;
    total_rows: number;
    total_pages: number;
    rows: T[]
}
