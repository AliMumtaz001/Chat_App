import React, { useState, useEffect, useRef } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { handleSendMessage, setupWebSocket } from '../apis/sendmsg';

const ChatSection = ({ selectedUser }) => {
    const [messages, setMessages] = useState([]);
    const [messageInput, setMessageInput] = useState('');
    const chatBoxRef = useRef(null);
    const wsRef = useRef(null);
    const token = localStorage.getItem('accessToken');

    useEffect(() => {
        if (!token || !selectedUser) return;

        const ws = setupWebSocket(token, (receivedMessage) => {
            console.log("WebSocket message received:", receivedMessage);
            setMessages((prev) => [...prev, { ...receivedMessage, timestamp: new Date().toLocaleTimeString() }]);
        });
        wsRef.current = ws;

        return () => {
            if (wsRef.current) {
                wsRef.current.close();
            }
        };
    }, [token, selectedUser]);

    useEffect(() => {
        if (chatBoxRef.current) {
            chatBoxRef.current.scrollTop = chatBoxRef.current.scrollHeight;
        }
    }, [messages]);

    const handleSend = async () => {
        if (messageInput.trim() && selectedUser) {
            const result = await handleSendMessage(selectedUser.id, messageInput, token);
            if (result.success) {
                setMessages((prev) => [
                    ...prev,
                    { message: messageInput, timestamp: new Date().toLocaleTimeString() }
                ]);
                setMessageInput('');
            } else {
                console.error('Send failed:', result.error);
            }
        }
    };

    const handleLogout = () => {
        localStorage.removeItem('accessToken');
        console.log('Logged out');
        navigate('/login');
    };

    return (
        <div className="flex-1 p-4">
            <button
                onClick={handleLogout}
                className="mb-4 p-2 bg-red-600 text-white rounded hover:bg-red-700 "
            >
                Logout
            </button>
            {selectedUser && <h2 className="text-xl font-bold mb-4">{selectedUser.username}</h2>}
            <div ref={chatBoxRef} className="lg:h-[650px] xl:h-[650px] overflow-y-auto border p-2 mb-4">
                {messages.map((msg, index) => (
                    <div
                        key={index}
                        className={`mb-2 ${msg.sender_id === selectedUser.id ? 'text-left' : 'text-right'
                            }`}
                    >
                        <span
                            className={`inline-block p-2 rounded-lg ${msg.sender_id === selectedUser.id ? 'bg-gray-300' : 'bg-blue-500 text-white'
                                }`}
                        >
                            {msg.message || msg.content}
                        </span>
                        <span className="text-xs text-gray-500">{msg.timestamp}</span>
                    </div>
                ))}
            </div>
            <div className="flex">
                <input
                    type="text"
                    value={messageInput}
                    onChange={(e) => setMessageInput(e.target.value)}
                    onKeyPress={(e) => e.key === 'Enter' && handleSend()}
                    className="flex-1 p-2 border rounded-l-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="Type a message..."
                />
                <button
                    onClick={handleSend}
                    className="p-2 bg-blue-600 text-white rounded-r-lg hover:bg-blue-700"
                >
                    Send
                </button>
            </div>
        </div>
    );
};

export default ChatSection;