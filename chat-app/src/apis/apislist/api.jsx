import axios from 'axios';
import { signupapi } from '../constants/apiendpoints';
import { loginapi } from '../constants/apiendpoints';
import { sendMessageApi } from '../constants/apiendpoints';
import { getUsersApi } from '../constants/apiendpoints';

export const signup = (username, email, password) => {
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

export const sendMessage = (receiverId, content, token) => {
  const payload = {
    receiver_id: receiverId,
    content: content,
  };
  console.log('Sending payload to /sendmessage:', payload); // Debug log
  return axios.post(
    `${sendMessageApi}`,
    payload,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );
};

export const getUsers = (searchQuery, token) => {
  return axios.get(`${getUsersApi}`, {
    params: { user: searchQuery }, 
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};