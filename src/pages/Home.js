import React, { useEffect, useState } from 'react';
import { fetchAllRooms } from '../axios/roomAxios';
import { useNavigate } from 'react-router-dom';
import './Home.css'; // Import the CSS file

const Home = () => {
    const [data, setData] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchRooms = async () => {
            try {
                const response = await fetchAllRooms();
                if (response && Array.isArray(response.rooms)) {
                    setData(response.rooms);
                } else {
                    console.error('Expected an array but got:', response);
                }
            } catch (error) {
                console.error('Error fetching rooms:', error);
            }
        };
        fetchRooms();
    }, []); // Empty dependency array means this effect runs once on mount

    const handleConnectRoom = (roomId) => {
        navigate(`/room/${roomId}`);
    };

    return (
        <div className="home-container">
            <h1>Available Rooms</h1>
            {data.length > 0 ? (
                <ul className="room-list">
                    {data.map((room) => (
                        <li key={room.RoomId} className="room-item" onClick={() => handleConnectRoom(room.RoomId)}>
                            {room.RoomName}
                        </li>
                    ))}
                </ul>
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
};

export default Home;