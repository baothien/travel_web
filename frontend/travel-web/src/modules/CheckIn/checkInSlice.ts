import { RootState } from '../../app/store';
import { PayloadAction, createSlice } from '@reduxjs/toolkit';

interface CheckInState{
    bgImgURL: string;
    capturedImgUrl: string;
    currentPlaceId: string;
}

const initialState:CheckInState = {
    bgImgURL: '',
    capturedImgUrl: '',
    currentPlaceId: '',
}

const checkInSlice = createSlice({
    name: 'checkIn',
    initialState,
    reducers:{
        setBgImgURL(state, action){
            state.bgImgURL = action.payload;
            return state
        },

        setCapturedImgUrl(state, action){
            state.capturedImgUrl = action.payload
            return state
        },

        setCurrentPlaceId(state, action) {
            state.currentPlaceId = action.payload
            return state
        }
        
    }
})

//actions
export const checkInActions = checkInSlice.actions

//selector
export const selectBgImgURL = (state: RootState) => state.checkIn.bgImgURL
export const selectCapturedImgUrl = (state: RootState) => state.checkIn.capturedImgUrl
export const selectCurrentPlaceId = (state: RootState) => state.checkIn.currentPlaceId

//reducer
const checkInReducer  = checkInSlice.reducer
export default checkInReducer