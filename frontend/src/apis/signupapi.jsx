import axios from 'axios';
const API_URL = 'http://localhost:8005'; // Updated to remove trailing /signup
export const signup = async (userData) => {
  try {
    const response = await axios.post(`${API_URL}/signup`, userData);
    return response.data;
  } catch (error) {
    console.error('Signup error:', error.response?.data || error); // Log error for debugging
    throw error.response?.data || { message: 'An error occurred during signup' };
  }
};