import { User } from "./user";

export interface CreateReviewParams{
    description: string;
    place_id: string;
    rate: number;
    review_img:Array<{
        name: string;
        url: string | undefined
    }>
}

export interface Review{
    id: string;
    place_id: string;
    user_id: string;
    user: User;
    rate: number;
    description: string;
    review_img?: Array<{
        name: string;
        url: string | undefined
    }>;
    child_review: ChildReview[];
    created_at: string;
}

export interface ChildReviewImg{
    id: string;
    child_review_id: string;
    name: string;
    url: string;
}
export interface ChildReview{
    id: string;
    parent_id: string;
    user_id: string;
    user: User;
    description: string;
    review_img?: ChildReviewImg[];
    created_at: string;
}

export interface ResponseReplyReview{
    id: string;
    parent_id: string;
    user_id: string;
    description: string;
    created_at: string;
}

export interface ReplyReviewParams{
    description: string;
    parent_id: string;
    review_img?: Array<{
        name: string;
        url: string | undefined
    }>;
}

