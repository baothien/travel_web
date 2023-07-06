import { authActions } from './authSlice';
import { call, fork, put, take } from "redux-saga/effects"



function* watchLoginFLow(){
    yield console.log('login flow')

    // while (true) {
    //     yield take(authActions.logoutSuccess.type);
    //     yield call(handleLogout);
    // }
}

function* authSaga(){
    yield fork(watchLoginFLow)
}

export default authSaga