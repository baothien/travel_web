import {all} from 'redux-saga/effects';
import authSaga from '../modules/Authentication/authSaga';

export default function* rootSaga(){
    yield all([authSaga()])
}