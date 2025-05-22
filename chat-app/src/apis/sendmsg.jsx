import { sendMessage } from './apislist/api';
import { toast } from 'react-toastify';
import {getsocketApi} from './constants/apiendpoints';

export const handleSendMessage = async (receiverId, content, token) => {
  try {
    const response = await sendMessage(receiverId, content, token);
    return { success: true, data: response.data };
  } catch (error) {
    if (error.response) {
      return { success: false, error: error.response.data.error || 'Failed to send message' };
    } else if (error.request) {
      return { success: false, error: 'No response from server. Please try again.' };
    } else {
      return { success: false, error: 'An unexpected error occurred.' };
    }
  }
};

const sockettoken = localStorage.getItem('token');
export const setupWebSocket = (sockettoken, onMessageReceived) => {
  console.log('WebSocket token:', sockettoken);
  const ws = new WebSocket(getsocketApi+'protected/ws'); 

  ws.onopen = () => {
    console.log('WebSocket connected');
    ws.current.send(JSON.stringify({action: 'connect', message:"hello"}));
    ws.send(JSON.stringify({ sockettoken }));
  };

  ws.onmessage = (event) => {
    const message = JSON.parse(event.data);
    onMessageReceived(message);
  };

  ws.onclose = () => {
    console.log('WebSocket disconnected');
  };

  ws.onerror = (error) => {
    console.error('WebSocket error:', error);
  };

  return ws;
};