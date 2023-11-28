import {FormEvent, useMemo, useState} from "react";
import {LoginClient} from "./backend/client";
import {Header} from "./Header";

interface RegisterFormElements extends HTMLFormControlsCollection {
    usernameInput: HTMLInputElement,
    emailInput: HTMLInputElement,
    passwordInput: HTMLInputElement,
}
interface RegisterFormElement extends HTMLFormElement {
    readonly elements: RegisterFormElements
}
export function RegisterPage() {
    const [username, setUsername] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const backend = useMemo(() => new LoginClient("http://localhost:8180"), []);

    function handleRegister(event: FormEvent<RegisterFormElement>) {
        event.preventDefault()
        backend.register(username, email, password)
            .then(
                (resp) => {
                    console.log("Successfully logged in: " + resp.sessionToken);
                }
            )
            .catch(
                (error) => {
                    console.log("Failed to login: " + error)
                }
            )
    }

    return <>
        <Header />
        <form onSubmit={handleRegister}>
            <div className="form-item">
                <label htmlFor="usernameInput">Handle: </label>
                <input id="usernameInput"
                       value={username}
                       onChange={e => setUsername(e.target.value)}
                />
            </div>
            <div className="form-item">
                <label htmlFor="emailInput">Email: </label>
                <input id="emailInput"
                       value={email}
                       onChange={e => setEmail(e.target.value)}
                />
            </div>
            <div className="form-item">
                <label htmlFor="passwordInput">Password: </label>
                <input id="passwordInput"
                       type="password"
                       value={password}
                       onChange={e => setPassword(e.target.value)}
                />
            </div>
            <input type="submit" value="Register" />
        </form>
    </>
}