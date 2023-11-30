import React, {useState} from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css';
import {Dashboard} from '../Dashboard/Dashboard';
import {Preferences} from '../Preferences/Preferences';
import {Login} from "../Login/Login";
import {Header} from "../Header/Header";
import useLoginData from "./useLogin";
function App() {
    const { loginData, setLogin } = useLoginData()

    if (!loginData) {
        // return <Login setToken={setSessionToken} />
        return <>
            <Header />
            <Login setToken={setLogin}/>
        </>
    }
    return (
        <div className="wrapper">
            <Header />
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Dashboard />} />
                    <Route path="/preferences" element={<Preferences />} />
                </Routes>
            </BrowserRouter>
        </div>
    )
}

export default App;
