import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { RootState } from "../../app/store";
import { User } from "../../models/user";
import { stat } from "fs";

interface AuthState {
  isLoggedIn: boolean;
  currentUser: User | null;
}

const initialState: AuthState = {
  isLoggedIn: false,
  currentUser: null,
};

export const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    loginSuccess: (state, action: PayloadAction<User>) => {
      state.currentUser = action.payload;
      state.isLoggedIn = true;
    },

    logoutSuccess: (state) => {
      state.currentUser = null;
      state.isLoggedIn = false;
    },

    updateUser: (state, action: PayloadAction<any>) => {
      if (state.currentUser)
        state.currentUser = {
          ...state.currentUser,
          avatar: action.payload.avatar,
          date_of_birth: action.payload.date_of_birth,
          email: action.payload.email,
          full_name: action.payload.full_name,
          phone: action.payload.phone,
        };
    },
  },
});

//actions
export const authActions = authSlice.actions;

//selectors
export const selectUser = (state: RootState) => state.auth.currentUser;
export const selectIsLoggedIn = (state: RootState) => state.auth.isLoggedIn;

export const selectAccessToken = (state: RootState) =>
  state.auth.currentUser?.token.access_token;

export const selectRefreshToken = (state: RootState) =>
  state.auth.currentUser?.token.refresh_token;

export const selectExpiredToken = (state: RootState) =>
  state.auth.currentUser?.token.expired_time;

//reducer
const authReducer = authSlice.reducer;
export default authReducer;
