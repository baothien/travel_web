import { ResponseMessage } from '../models/common';
import {axiosClient} from './axiosClient';

const profileApi = {
    updateProfile(params:any) : Promise<ResponseMessage<any>>{
        const url = '/user-service/user/profile'
        return axiosClient.patch(url, params)
    },

    changePassword(params:any):Promise<any>{
        const url='/user-service/user/change-pass'
        return axiosClient.patch(url, params)
    }
}

export default profileApi