import axios from './customizeAxios';

const fetchAllRooms = async () => {
    return await axios.get('/rooms');
}

const fetchRoomById = async (roomId) => {
    return await axios.get(`/rooms/${roomId}`);
}

const createRoom = async (room) => {
    return await axios.post('/rooms', room);
}

const updateRoom = async (room) => {
    return await axios.put(`/rooms/${room.id}`, room);
}

const deleteRoom = async (roomId) => {
    return await axios.delete(`/rooms/${roomId}`);
}

export {
    fetchAllRooms,
    fetchRoomById,
    createRoom,
    updateRoom,
    deleteRoom
}