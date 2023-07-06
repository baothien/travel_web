import { PaginateData, ResponseMessage } from '../models/common';
import {axiosClient} from './axiosClient';

const checkInApi = {
    postCheckInImg(params: any): Promise<any> {
        const url = "/place-service/checkin/create";
        return axiosClient.post(url, params);
    },

    getCheckInList():Promise<any>{
        const url='/place-service/checkin/list'
        return axiosClient.get(url)
    }
}

export default checkInApi