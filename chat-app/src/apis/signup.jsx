import { signup } from './apislist/api';
export const handleSignup = async (username,email, password) => {
    try {
        const response = await signup(username,email, password);
        return { success: true, data: response.data };
    } catch (error) {
        if (error.response) {
            return { success: false, error: error.response.data.error || 'Signup failed' };
        } else if (error.request) {
            return { success: false, error: 'No response from server. Please try again.' };
        } else {
            return { success: false, error: 'An unexpected error occurred.' };
        }
    }
};