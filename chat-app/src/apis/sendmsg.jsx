import { sendMessage } from './apislist/api';
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

let ws;

export const setupWebSocket = (onMessageReceived) => {
  if (ws && ws.readyState !== WebSocket.CLOSED) {
    console.log('WebSocket already connected');
    return ws;
  }

  const sockettoken = localStorage.getItem('accessToken');
  console.log('WebSocket token:', sockettoken);
  ws = new WebSocket(getsocketApi + 'protected/ws');

  ws.onopen = () => {
    console.log('WebSocket connected');
    ws.send(JSON.stringify({ action: 'connect', message: "hello", token: sockettoken }));
  };

  ws.onmessage = (event) => {
    const message = JSON.parse(event.data);
    onMessageReceived(message);
  };

  ws.onclose = () => {
    console.log('WebSocket disconnected');
    setTimeout(() => {
      console.log('Attempting to reconnect...');
      setupWebSocket(onMessageReceived);
    }, 3000); 
  };

  ws.onerror = (error) => {
    console.error('WebSocket error:', error);
  };

  return ;
};