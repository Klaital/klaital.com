import React, {Dispatch, FormEvent, SetStateAction, useMemo, useState} from 'react'
import './Login.css'
import PropTypes from "prop-types";
import {LoginClient} from "../../backend/client";

export function Login(props: {setToken: Dispatch<SetStateAction<string>>}) {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const backend = useMemo(() => new LoginClient("http://localhost:8180"), [])

    const handleSubmit = async (e:FormEvent) => {
        e.preventDefault()
        const resp = backend.register(username, email, password)
        resp.then(resp => {
            props.setToken(resp.sessionToken)
        })
    }

    return <>
        <div className="login-wrapper">
            <h1>Please Log In</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    <p>Username: </p>
                    <input type="text" onChange={e => setUsername(e.target.value)} />
                </label>
                <label>
                    <p>Email: </p>
                    <input type="text" onChange={e => setEmail(e.target.value)} />
                </label>
                <label>
                    <p>Password: </p>
                    <input type="password" onChange={e => setPassword(e.target.value)} />
                </label>
                <div>
                    <button type="submit">Submit</button>
                </div>
            </form>
        </div>
    </>
}
