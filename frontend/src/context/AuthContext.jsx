import React, {createContext, useContext, useState} from "react";
import jwtDecode from "jwt-decode";

const AuthContext = createContext();

export const AuthProvider = (props) => {
    const [account, setAcc] = useState(localStorage.getItem('account.data'))

    const setAccount = (accountInfo) => {
        let body = jwtDecode(accountInfo.access_token);
        accountInfo.username = body.username;
        localStorage.setItem('account.data', JSON.stringify(accountInfo))
        setAcc(accountInfo)
    }
    return <AuthContext.Provider value={{account, setAccount}} {...props} />
}

export const useAuth = () => {
    const account = useContext(AuthContext)

    if (!account) {
        throw new Error('useAuth must be used inside AuthProvider')
    }
    return account;
}