import { login } from './apislist/api';
export const handleLogin = async (email, password) => {
    try {
        const response = await login(email, password);
        return { success: true, data: response.data };
    } catch (error) {
        if (error.response) {
            return { success: false, error: error.response.data.error || 'Login failed' };
        } else if (error.request) {
            return { success: false, error: 'No response from server. Please try again.' };
        } else {
            return { success: false, error: 'An unexpected error occurred.' };
        }
    }
};