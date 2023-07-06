import { Review } from "./review";

export interface PlaceType{
    id: string;
    code: string;
    name: string;
}

export interface PlaceImg{
    id: string;
    place_id: string;
    name: string;
    url: string;
}

export interface Place{
    id: string;
    thumbnail: string;
    name: string;
    place_type_id?: string;
    place_type: PlaceType;
    lat: number;
    lng: number;
    address: string;
    place_img: PlaceImg[];
    review?: Review[];
    created_at?: string;
    updated_at?: string;
}

export interface FavPlaceParams{
    is_favorite: boolean;
    place_id: string;
}