import { PaginateData, ResponseMessage } from './../models/common';
import { axiosClient } from './axiosClient';
import { ChildReview, CreateReviewParams, ReplyReviewParams, ResponseReplyReview, Review } from './../models/review';
import { PaginateParams, Response } from '../models/common';
const reviewApi = {
    postReview(params: CreateReviewParams) : Promise<any>{
        const url = '/place-service/review/create'
        return axiosClient.post(url, params)
    },

    getReviews(placeId:string, paginate: PaginateParams): Promise<Response<PaginateData<Review>>>{
        const url = `/place-service/review/list/${placeId}`
        return axiosClient.get(url, {
            params: {
                page: paginate.page,
                limit: paginate.limit
            }
        })
    },

    postReplyReview(params:ReplyReviewParams): Promise<ResponseMessage<ResponseReplyReview>>{
        const url = `/place-service/review/child/create/${params.parent_id}`
        return axiosClient.post(url, params)
    },

    getReplyReview(reviewId: string): Promise<Response<PaginateData<ChildReview>>>{
        const url = `/place-service/child-review/list/${reviewId}`
        return axiosClient.get(url, {
            params: {
                page: 1,
                limit: 100
            }
        })
    }

}

export default reviewApi