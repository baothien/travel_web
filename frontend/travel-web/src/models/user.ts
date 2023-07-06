export interface UserToken{
    access_token: string;
    refresh_token: string;
    expired_time: number;
}

export interface User{
    id: string;
    user_name: string;
    user_type: string;
    full_name: string;
    phone: string;
    email: string;
    password: string;
    date_of_birth: string;
    avatar: string;
    token: UserToken;
    createdAt: string;
}