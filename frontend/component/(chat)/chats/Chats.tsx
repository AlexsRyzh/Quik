'use client'

import React, {createContext} from "react";
import styles from './chats.module.scss'
import {StaticImageData} from "next/image";

interface ChatItem {
    idUser: number,
    name: string,
    img: StaticImageData,
    surname: string,
    lastMessage: {
        text: string,
        date: Date
    }

}

interface ContextType {
    chatList?: ChatItem[]
}

export const ChatListContext = createContext<ContextType>({})

interface Props {
    chatList: ChatItem[],
    children: React.ReactNode
}

export default function Chats(props: Props) {

    const {
        chatList,
        children
    } = props

    return (
        <ChatListContext.Provider value={{chatList}}>
            <div className={styles.container}>
                {children}
            </div>
        </ChatListContext.Provider>
    )
}