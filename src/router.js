import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import BoxChat from './pages/BoxChat';

const Router = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/room/:id" element={<BoxChat />} />
            </Routes>
        </BrowserRouter>
    );
};

export default Router;