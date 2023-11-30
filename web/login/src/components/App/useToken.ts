import {useState} from "react";

export default function useToken() {
    const getToken = () => {
        return localStorage.getItem('sessionToken')
    }
    const [token, setToken] = useState(getToken())

    const saveToken = (userToken: string) => {
        localStorage.setItem('sessionToken', userToken)
        setToken(userToken)
    }

    return {
        setToken: saveToken,
        token
    }
}
