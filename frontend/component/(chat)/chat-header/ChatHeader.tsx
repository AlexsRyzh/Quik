'use client'

import styles from './chat-header.module.scss'
import {useContext} from "react";
import {ChatPageContext} from "@/component/(chat)/chat-page/ChatPage";
import Link from "next/link";
import Image from "next/image";
import {ButtonLink} from "@/component/button-link/ButtonLink";
import clsx from "clsx";

export default function ChatHeader() {

    const {chatData} = useContext(ChatPageContext)

    return (
        <div className={styles.container}>
            <ButtonLink href={'/chats'} className={styles.button}>
                <span className={clsx(
                    "material-symbols-rounded",
                    styles.buttonBack
                )}>
                    arrow_back_ios_new
                </span>
                Назад
            </ButtonLink>
            <Link href={'/profiles/1'} className={styles.info}>
                {chatData?.userTo.name} {chatData?.userTo.surname}
            </Link>
            <Link href={'/profiles/1'}>
                <Image src={chatData?.userTo.imgLink || ""} alt={"Фото"} className={styles.img}/>
            </Link>
        </div>
    )

}