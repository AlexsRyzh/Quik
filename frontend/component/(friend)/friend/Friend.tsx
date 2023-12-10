'use client'

import React, {createContext} from "react";
import {StaticImageData} from "next/image";
import styles from './friend-list.module.scss'

interface Friend {
    img: StaticImageData,
    name: string,
    surname: string,
    userID: number
}

interface ContextType {
    friendList?: Friend[]
}

export const FriendListContext = createContext<ContextType>({})


interface Props {
    friendList: Friend[],
    children: React.ReactNode,
}

export default function Friend(props: Props) {

    const {friendList, children} = props

    return (
        <FriendListContext.Provider value={{friendList}}>
            <div className={styles.container}>
                {children}
            </div>
        </FriendListContext.Provider>
    )

}