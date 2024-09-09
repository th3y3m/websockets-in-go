import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { createMessage, fetchMessagesOfRoom } from '../axios/messageAxios';
import './BoxChat.css'; // Import the CSS file

const BoxChat = () => {
    const { id: roomId } = useParams();
    const [message, setMessage] = useState('');
    const [messages, setMessages] = useState([]);
    const [socket, setSocket] = useState(null);

    const handleSendMessage = async () => {
        if (socket && message) {
            socket.send(message);
            await createMessage({
                "room_id": parseInt(roomId, 10), // Ensure roomId is an integer
                "user_id": parseInt(roomId, 10),
                "message": message
            });
            setMessage('');
        } else {
            alert('Please join a room and type a message.');
        }
    };

    useEffect(() => {
        const fetchMessages = async () => {
            try {
                const response = await fetchMessagesOfRoom(roomId);
                if (response && Array.isArray(response.messages)) {
                    setMessages(response.messages);
                } else {
                    console.error('Expected an array but got:', response);
                }
            } catch (error) {
                console.error('Error fetching messages:', error);
            }
        };

        if (roomId) {
            fetchMessages();

            const ws = new WebSocket(`ws://localhost:8080/ws?room=${roomId}`);

            ws.onopen = () => {
                console.log(`Connected to room: ${roomId}`);
            };

            ws.onmessage = (event) => {
                try {
                    const parsedData = JSON.parse(event.data);
                    setMessages((prevMessages) => [...prevMessages, parsedData]);
                } catch (e) {
                    setMessages((prevMessages) => [...prevMessages, { message: event.data }]);
                }
            };

            ws.onclose = () => {
                console.log(`Disconnected from room: ${roomId}`);
            };

            setSocket(ws);

            return () => {
                ws.close();
            };
        }
    }, [roomId]);

    const handleBack = () => {
        window.history.back();
    };

    const handleKeyDown = (e) => {
        if (e.key === 'Enter') {
            handleSendMessage();
        }
    };

    return (
        <div className="box-chat">
            <button className="back-button" onClick={handleBack}>Back</button>
            <div className="chat-window">
                <div className="messages" id="chat">
                    {messages.map((msg, index) => (
                        <div key={index} className="message">{msg.message}</div>
                    ))}
                </div>
                <div className="input-area">
                    <input
                        type="text"
                        value={message}
                        onChange={(e) => setMessage(e.target.value)}
                        onKeyDown={handleKeyDown} // Add this line
                        placeholder="Type a message..."
                        className="message-input"
                    />
                    <button className="send-button" onClick={handleSendMessage}>Send</button>
                </div>
            </div>
        </div>
    );
};

export default BoxChat;