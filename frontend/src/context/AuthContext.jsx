import React, {createContext, useContext, useState} from "react";

const AuthContext = createContext();

export const AuthProvider = (props) => {
    const [account, setAccount] = useState(null)
    return <AuthContext.Provider value={{account, setAccount}} {...props} />
}

export const useAuth = () => {
    const account = useContext(AuthContext)

    if (!account) {
        throw new Error('useAuth must be used inside AuthProvider')
    }
    return account;
}