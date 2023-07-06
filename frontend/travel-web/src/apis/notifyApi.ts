import { PaginateData, PaginateParams } from './../models/common';
import { Response } from '../models/common';
import { axiosClient } from './axiosClient';
import { Notify } from '../models/notify';


const notifyApi = {
    getNotifyList(paginate: PaginateParams):Promise<Response<PaginateData<Notify>>>{
        const url = '/notify-service/notify/list'
        return axiosClient.get(url, {
            params:{
                page: paginate.page,
                limit: paginate.limit
            }
        })
    },

    patchReadNotify(id:string){
        const url = `/notify-service/notify/is-read/${id}`
        return axiosClient.patch(url, {
            "is_read": true
        })
    },

    getNotifyCount(){
        const url = `/notify-service/notify/count`
        return axiosClient.get(url)
    }
}

export default notifyApi