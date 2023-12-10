'use client'

import styles from './chat-message-list.module.scss'
import {useContext} from "react";
import {ChatPageContext} from "@/component/(chat)/chat-page/ChatPage";
import ChatMessage from "@/component/(chat)/chat-message/ChatMessage";

export default function ChatMessageList() {

    const {chatData} = useContext(ChatPageContext)

    return (
        <div className={styles.container}>
            {chatData?.messages.map((v) => {
                return (
                    <ChatMessage
                        message={v.message}
                        img={v.img}
                        from={v.from}
                        date={v.date}
                        my={v.from === 1}
                    />
                )
            })}
        </div>
    )
}