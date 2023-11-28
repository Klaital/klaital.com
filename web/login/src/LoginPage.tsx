import {FormEvent, useMemo, useState} from "react";
import {LoginClient} from "./backend/client";

interface LoginFormElements extends HTMLFormControlsCollection {
    emailInput: HTMLInputElement,
    passwordInput: HTMLInputElement,
}
interface LoginFormElement extends HTMLFormElement {
    readonly elements: LoginFormElements
}
export function LoginPage() {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const backend = useMemo(() => new LoginClient("http://localhost:8180"), []);

    function handleLogin(event: FormEvent<LoginFormElement>) {
        event.preventDefault()
        backend.login(email, password)
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
        <form onSubmit={handleLogin}>
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
            <input type="submit" value="Login" />
        </form>
    </>
}