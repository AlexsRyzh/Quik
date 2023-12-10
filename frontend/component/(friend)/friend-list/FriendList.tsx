'use client'


import {useContext} from "react";
import {FriendListContext} from "@/component/(friend)/friend/Friend";
import styles from './friend-list.module.scss'
import Image from "next/image";
import Link from "next/link";

export default function FriendList() {

    const {friendList} = useContext(FriendListContext)

    return (
        <>
            {friendList?.map((v) => {
                return (
                    <Link
                        href={`/profiles/${v.userID}`}
                        className={styles.container}
                    >
                        <Image src={v.img} alt={"Фото"} className={styles.img}/>
                        <p className={styles.fio}>{v.name} {v.surname}</p>
                    </Link>
                )
            })}
        </>
    )
}