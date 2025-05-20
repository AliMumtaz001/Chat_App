import axios from 'axios';
import { signupapi } from '../constants/apiendpoints';
import { loginapi } from '../constants/apiendpoints';
export const signup = (username,email, password) => {
  return axios.post(`${signupapi}`, {
    username,
    email,
    password,
  });
};

export const login = (email, password) => {
  return axios.post(`${loginapi}`, {
    email,
    password,
  });
};