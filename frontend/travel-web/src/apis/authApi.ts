import {
  LoginParams,
  RegisterParams,
  Response,
  ResponseMessage,
} from "../models/common";
import { User } from "../models/user";
import { axiosClient, axiosClientRefreshToken } from "./axiosClient";

const authApi = {
  postRegister(params: RegisterParams): Promise<ResponseMessage<User>> {
    const url = "/user-service/user/register";
    return axiosClient.post(url, params);
  },

  postLogin(params: LoginParams): Promise<Response<User>> {
    const url = "/user-service/user/login";
    return axiosClient.post(url, params);
  },

  getRefreshToken(refreshToken:string): Promise<Response<User>>{
    const url = '/user-service/user/refresh-token'
    return axiosClientRefreshToken.get(url,{
      headers:{
        "Authorization": `Bearer ${refreshToken}`
      }
    })
  }
};

export default authApi;
