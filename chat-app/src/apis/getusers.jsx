import { getUsers } from './apislist/api';
import { toast } from 'react-toastify';

export const handleGetUsers = async (searchQuery, token) => {
  try {
    const response = await getUsers(searchQuery, token);
    console.log('Get Users Response:', response.data); // Debug log
    return { success: true, data: response.data };
  } catch (error) {
    if (error.response) {
      return { success: false, error: error.response.data.error || 'Failed to fetch users' };
    } else if (error.request) {
      return { success: false, error: 'No response from server. Please try again.' };
    } else {
      return { success: false, error: 'An unexpected error occurred.' };
    }
  }
};