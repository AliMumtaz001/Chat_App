import { sendMessage } from './apislist/api';
import { toast } from 'react-toastify';

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

export const setupWebSocket = (token, onMessageReceived) => {
  const ws = new WebSocket('ws://localhost:8004/protected/ws'); 

  ws.onopen = () => {
    console.log('WebSocket connected');
    // Send authentication token if required by your backend
    ws.send(JSON.stringify({ token }));
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

// export const setupWebSocket = (token, onMessageReceived) => {
//   let ws;
//   let reconnectAttempts = 0;
//   const maxReconnectAttempts = 5;
//   const reconnectInterval = 3000; // 3 seconds

//   const connect = () => {
//     ws = new WebSocket('ws://localhost:8080/ws');

//     ws.onopen = () => {
//       console.log('WebSocket connected');
//       reconnectAttempts = 0; // Reset reconnect attempts on successful connection
//       // Send authentication token
//       ws.send(JSON.stringify({ token }));
//     };

//     ws.onmessage = (event) => {
//       try {
//         const message = JSON.parse(event.data);
//         if (message.error) {
//           toast.error(message.error); // Display errors like "Invalid token"
//           ws.close();
//         } else {
//           onMessageReceived(message);
//         }
//       } catch (error) {
//         console.error('Error parsing WebSocket message:', error);
//       }
//     };

//     ws.onclose = () => {
//       console.log('WebSocket disconnected');
//       if (reconnectAttempts < maxReconnectAttempts) {
//         setTimeout(() => {
//           console.log(`Reconnecting WebSocket (attempt ${reconnectAttempts + 1})...`);
//           reconnectAttempts++;
//           connect();
//         }, reconnectInterval);
//       } else {
//         toast.error('Failed to connect to WebSocket after multiple attempts.');
//       }
//     };

//     ws.onerror = (error) => {
//       console.error('WebSocket error:', error);
//     };
//   };

//   connect();
//   return ws;
// };