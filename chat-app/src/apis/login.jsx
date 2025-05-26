import { login } from './apislist/api';
export const handleLogin = async (email, password) => {
  try {
    const response = await login(email, password);
    return { success: true, data: response.data };
  } catch (error) {
    // Always return a string error message
    let errorMsg = 'An unexpected error occurred.';
    if (error.response && error.response.data && error.response.data.error) {
      errorMsg = error.response.data.error;
    } else if (error.message) {
      errorMsg = error.message;
    }
    return { success: false, error: errorMsg };
  }
};