import { RootState } from './../../app/store';
import { PayloadAction } from '@reduxjs/toolkit';
import {  createSlice } from '@reduxjs/toolkit';

interface NotifyState{
    count: number;
    cmtNotifiedId: string;
}

const initialState:NotifyState = {
    count: 0,
    cmtNotifiedId: '',

}

export const notifySlice = createSlice({
    name: 'notify',
    initialState,
    reducers: {
        setCount(state, action: PayloadAction<any>){
            state.count = action.payload
            return state
        },
        increaseCount(state){
            state.count++
            return state
        },
        decreaseCount(state){
            state.count--
            return state
        },

        setCmtNotifiedId(state, action: PayloadAction<string>){
            state.cmtNotifiedId = action.payload
            return state
        }
    }
})

//actions
export const notifyActions = notifySlice.actions

//selector
export const selectNotifyCount = (state: RootState) => state.notify.count
export const selectCmtNotifiedId = (state: RootState) => state.notify.cmtNotifiedId

//reducer
const notifyReducer = notifySlice.reducer 
export default notifyReducer