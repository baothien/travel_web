import { User } from "./user";

export interface Notify{
    id: string;
    child_id:string;
    destination_id: string;
    from_user_id: string;
    from_user: User;
    to_user_id: string;
    type: string;
    title: string;
    body: string;
    is_read: boolean;
    is_important: boolean;
    created_at: string;
}