'use client'

import {ChatListContext} from "@/component/(chat)/chats/Chats";
import {useContext} from "react";
import styles from './chat-list.module.scss'
import Image from "next/image";
import Link from "next/link";

export default function ChatsList() {

    const {chatList} = useContext(ChatListContext)

    return (
        <>
            {chatList?.map((v) => {

                return (
                    <Link href={`/chats/${v.idUser}`} className={styles.container}>
                        <Image src={v.img} alt={"Фото"} className={styles.img}/>
                        <div className={styles.info}>
                            <div className={styles.dateAndFio}>
                                <p>{v.name} {v.surname}</p>
                                <p>{v.lastMessage.date.getDate()}.{v.lastMessage.date.getMonth()}</p>
                            </div>
                            <p className={styles.message}>{v.lastMessage.text.slice(0, 50)}</p>
                        </div>
                    </Link>
                )
            })}
        </>
    )
}