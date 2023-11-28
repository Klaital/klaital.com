import {LoginResponse, RegisterResponse} from "./login";

export class LoginClient {
    rootUrl: string;

    constructor(rootUrl: string) {
        this.rootUrl = rootUrl;
    }

    async login(email: string, password: string): Promise<LoginResponse> {
        const resp = await fetch(this.rootUrl + "/api/login/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                email: email,
                password: password,
            })
        })
        return await resp.json();
    }

    async register(username: string, email: string, password: string): Promise<RegisterResponse> {
        const resp = await fetch(this.rootUrl + "/api/login/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                username: username,
                email: email,
                password: password,
            })
        })
        return await resp.json();
    }
}