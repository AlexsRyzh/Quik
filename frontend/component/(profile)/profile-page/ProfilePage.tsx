'use client'

import React, {createContext} from "react";

interface ContextType {

}

export const ProfilePageContext = createContext<ContextType>({})

interface Props {
    value: ContextType,
    children: React.ReactNode
}

export default function ProfilePage(props: Props) {

    const {value, children} = props

    return (
        <ProfilePageContext.Provider value={value}>
            {children}
        </ProfilePageContext.Provider>
    )
}