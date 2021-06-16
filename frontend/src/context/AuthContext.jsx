import React, {createContext, useContext, useEffect, useState} from "react";
import jwtDecode from "jwt-decode";
import {RefreshToken} from "../http/Accounts";

const AuthContext = createContext();

export const AuthProvider = (props) => {
    const [account, setAcc] = useState(() => {
        return JSON.parse(localStorage.getItem('account.data'));
    })

    // Wrapper for the setAcc state hook. This wrapper will update the local storage
    // prior to calling the setAcc hook.
    const setAccount = (accountInfo) => {
        let body = jwtDecode(accountInfo.access_token);
        accountInfo.username = body.username;

        localStorage.setItem('account.data', JSON.stringify(accountInfo));
        setAcc(accountInfo);
    }

    useEffect(() => {
        if(account === null) { return; }
        let body = jwtDecode(account.access_token);
        const now = Date.now();
        let difference = (body.exp * 1000) - now;

        difference -= 5000; // Call refresh 5 seconds before expiration.
        setTimeout(refreshToken, difference);
    }, [account])

    const refreshToken = () => {
        RefreshToken(account.refresh_token).then(resp => {
            if(resp.error.status !== 200) {
                setAcc(null);
                return;
            }
            setAccount(resp.body);
        })
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