import { PaginateData, PaginateParams } from './../models/common';
import { FavPlaceParams } from './../models/place';
import { Response } from "../models/common";
import { Place } from "../models/place";
import { axiosClient } from "./axiosClient";

const placesApi = {
  getPlaces(): Promise<Response<Place[]>> {
    const url = "/place-service/place/list";
    return axiosClient.get(url);
  },

  getPlace(id: string): Promise<Response<Place>> {
    const url = `/place-service/place/detail/${id}`;
    return axiosClient.get(url);
  },

  postFavorite(params:FavPlaceParams):Promise<any>{
    const url = 'place-service/place/favorite'
    return axiosClient.post(url, params)
  },

  getFavoriteList(params: PaginateParams): Promise<Response<PaginateData<Place>>>{
    const url = 'place-service/place/favorite/list'
    return axiosClient.get(url, {
      params
    })
  }
};

export default placesApi;
