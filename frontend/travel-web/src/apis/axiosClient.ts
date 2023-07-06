
import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import queryString from "query-string";
import authApi from './authApi';
import { store } from '../app/store';
import { toast } from 'react-toastify';

const baseURL = "https://travel-api.huytx.com/stag/";

export const axiosClient = axios.create({
  baseURL: baseURL,
  headers: {
    "Content-Type": "application/json",
  },
  paramsSerializer: (params) => queryString.stringify(params),
});

axiosClient.interceptors.request.use(
  function (config: any) {
   const localWeb = localStorage.getItem('persist:trave-web')

   if(localWeb){
      const isLogged = store.getState().auth.isLoggedIn
      const accessToken = store.getState().auth.currentUser?.token.access_token

      if(isLogged){
        config.headers.Authorization = `Bearer ${accessToken}`
      }
   }
    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

axiosClient.interceptors.response.use(
  function (response: AxiosResponse) {
    return response.data;
  },

  async function (error) {
    // const navigate = useNavigate()

    const originalConfig = error.config

    if (error.response) {
      const errorMessage = error.response.data.error_info.message[0]

      if(errorMessage === 'Token is expired'){
        const localWeb = localStorage.getItem('persist:trave-web')

        if(localWeb){
          const refreshToken = store.getState().auth.currentUser?.token.refresh_token
          const state:any = store.getState()

          if(refreshToken){
            authApi.getRefreshToken(refreshToken).then(res => {
              state.auth.currentUser.token.access_token = res.data.token.access_token

              localStorage.setItem('persist:trave-web', JSON.stringify(state))

            }).catch(err => {
              toast.error('Phiên đăng nhập đã hết hạn')
              // destroyToken();
              // this.router.push("/login");
              state.auth.currentUser = null
              state.auth.isLoggedIn = false

              // navigate("/")
              
              localStorage.setItem('persist:trave-web', JSON.stringify(state))
              return Promise.reject(err);
            })
          }
        }
        return axiosClient(originalConfig)
      }
      
    return Promise.reject(error);
    }
  }
);

export const axiosClientRefreshToken = axios.create({
  baseURL: baseURL,
  headers: {
    "Content-Type": "application/json",
  },
})

axiosClientRefreshToken.interceptors.response.use(
  function (response: AxiosResponse) {
    return response.data;
  },

  function (error) {
    if (error.response) {
    return Promise.reject(error);
  }
  }
);