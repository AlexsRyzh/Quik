'use client'

import React, {createContext, useEffect, useState} from "react";
import {usePathname, useRouter} from "next/navigation";

interface ContextType {
    userID?: number
}

export const AuthContextProvider = createContext<ContextType>({})

interface Props {
    children: React.ReactNode
}

const openLink = [
    "/login",
    "/greeting",
    "/register"
]
export default function AuthContext(props: Props) {

    const {children} = props

    const router = useRouter()
    const pathname = usePathname()

    const access_token = global?.localStorage?.getItem('access_token');

    const [value, setValue] = useState({})

    useEffect(() => {
        if (!access_token && !openLink.includes(pathname)) {
            router.push('/greeting')
        }

        if (access_token && openLink.includes(pathname)) {
            router.push('/')
        }

    }, [pathname, access_token]);

    return (
        <AuthContextProvider.Provider value={{...value}}>
            {children}
        </AuthContextProvider.Provider>
    )
}