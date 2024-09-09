import axios from './customizeAxios';

const fetchMessagesOfRoom = async (roomId) => {
    return await axios.get(`/messages/room/${roomId}`);
}

const createMessage = async (message) => {
    return await axios.post('/messages', message);
}

export {
    fetchMessagesOfRoom,
    createMessage
}