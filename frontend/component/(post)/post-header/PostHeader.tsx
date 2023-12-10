'use client'

import styles from './post-header.module.scss'
import Image from "next/image";
import Link from "next/link";
import {useContext, useEffect, useState} from "react";
import {PostCardContext} from "@/component/(post)/post-card/PostCard";
import $api from "@/http/api";
import img from '@/public/profile-img.png'

export default function PostHeader() {


    const {userID} = useContext(PostCardContext)
    const [user, setUser] = useState<{
        name?: string,
        surname?: string,
        imgLink?: string
    }>({})
    useEffect(() => {
        const fetch = async () => {
            try {
                const res = await $api.get(`/users/${userID}`)

                setUser({
                    name: res.data.name,
                    surname: res.data.surname,
                    imgLink: res.data['img_link']
                })

            } catch (e) {
                console.log(e)
            }
        }

        fetch()

    }, []);

    return (
        <div>
            <Link href={`/profile/${userID}`} className={styles.header}>
                <Image src={user.imgLink === "" ? img : user.imgLink} alt={"Профиль"} className={styles.img}/>
                <div className={styles.info}>
                    <p>{user.name} {user.surname}</p>
                    {/*<p className={styles.date}>{author?.date.getDate()}:{author?.date.getMonth()}:{authordate.getFullYear()}</p>*/}
                </div>
            </Link>
        </div>

    )
}