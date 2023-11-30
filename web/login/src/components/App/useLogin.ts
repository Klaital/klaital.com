import {useState} from "react";
import {LoginResponse} from "../../backend/login";

export default function useLoginData() {
    const getUser = (): LoginResponse | null => {
        const userRaw =  localStorage.getItem('userLogin')
        if (userRaw) {
            return JSON.parse(userRaw)
        } else {
            return null
        }
    }
    const [loginData, setLoginData] = useState(getUser())

    const saveLogin = (user: LoginResponse) => {
        localStorage.setItem('userLogin', JSON.stringify(user))
        setLoginData(user)
    }

    return {
        setLogin: saveLogin,
        loginData
    }
}
