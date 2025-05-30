import React, { useState, useEffect } from 'react';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { handleSendMessage } from '../apis/sendmsg';
import { handleGetUsers } from '../apis/getusers';
import { setupWebSocket } from '../apis/sendmsg';
import { useNavigate } from 'react-router-dom';

function Chat() {
  const [token, setToken] = useState(null);
  const [users, setUsers] = useState([]);
  const [selectedUser, setSelectedUser] = useState(null);
  const [messages, setMessages] = useState([]);
  const [messageInput, setMessageInput] = useState('');
  const [searchQuery, setSearchQuery] = useState('');
  const [ws, setWs] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const storedToken = localStorage.getItem('accessToken');
    if (!storedToken || storedToken === 'undefined') {
      toast.error('Access token not found or invalid. Please log in.');
      navigate('/login');
      return;
    }
    setToken(storedToken);

    const fetchUsers = async () => {
      const result = await handleGetUsers(searchQuery, storedToken);
      if (result.success) {
        setUsers(result.data.users || []);
      } else {
        toast.error(result.error);
        if (result.error.toLowerCase().includes('token')) {
          localStorage.removeItem('accessToken');
          navigate('/login');
        }
      }
    };
    fetchUsers();

    const websocket = setupWebSocket(storedToken, (message) => {
      setMessages((prev) => [...prev, message]);
    });
    setWs(websocket);

    return () => {
      if (websocket) websocket.close();
    };
  }, [searchQuery, navigate]);

  const sendMessage = async () => {
    if (!selectedUser || !messageInput.trim()) {
      toast.error('Please select a user and enter a message');
      return;
    }
    const receiverId = parseInt(selectedUser.id);
    if (isNaN(receiverId)) {
      toast.error('Invalid receiver ID');
      return;
    }
    const result = await handleSendMessage(receiverId, messageInput, token);
    if (result.success) {
      setMessageInput('');
    } else {
      toast.error(result.error);
    }
  };

  const filteredUsers = users.filter((user) =>
    user.username.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <div className="h-screen flex flex-col bg-gray-100">
      <ToastContainer />
      <header className="bg-blue-600 text-white p-4 text-center text-2xl font-bold">
        ChatSphere
      </header>
      <div className="flex flex-1 overflow-hidden">
        <div className="w-full md:w-1/3 lg:w-1/4 bg-white border-r flex flex-col">
          <div className="p-4">
            <input
              type="text"
              placeholder="Search users..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="w-full p-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div className="flex-1 overflow-y-auto">
            {filteredUsers.map((user) => (
              <div
                key={user.id}
                onClick={() => {
                  setSelectedUser(user);
                  setMessages([]);
                }}
                className={`p-4 cursor-pointer hover:bg-gray-100 ${selectedUser?.id === user.id ? 'bg-blue-100' : ''}`}
              >
                {user.username}
              </div>
            ))}
          </div>
        </div>
        <div className="flex-1 flex flex-col">
          {selectedUser ? (
            <>
              <div className="bg-white p-4 border-b">
                <h2 className="text-xl font-semibold">{selectedUser.username}</h2>
              </div>
              <div className="flex-1 p-4 overflow-y-auto bg-gray-50">
                {messages.map((msg, index) => (
                  <div
                    key={index}
                    className={`mb-2 p-2 rounded-lg max-w-md ${parseInt(msg.sender_id) === parseInt(selectedUser.id)
                        ? 'bg-blue-500 text-white ml-auto'
                        : 'bg-gray-200 text-black'
                      }`}
                  >
                    {msg.content}
                  </div>
                ))}
              </div>
              <div className="p-4 bg-white border-t flex">
                <input
                  type="text"
                  value={messageInput}
                  onChange={(e) => setMessageInput(e.target.value)}
                  placeholder="Type a message..."
                  className="flex-1 p-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                  onKeyPress={(e) => e.key === 'Enter' && sendMessage()}
                />
                <button
                  onClick={sendMessage}
                  className="ml-2 bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
                >
                  Send
                </button>
              </div>
            </>
          ) : (
            <div className="flex-1 flex items-center justify-center text-gray-500">
              Select a user to start chatting
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default Chat;