'use client'

import styles from './profile-info.module.scss'
import Image from "next/image";
import {useEffect, useState} from "react";
import bg from '@/public/background.png'
import ProfileImage from "@/component/(profile)/profile-image/ProfileImage";
import {Button} from "@/component/button/Button";
import clsx from "clsx";
import {useRouter} from "next/navigation";
import $api from "@/http/api";
import {User} from "@/const/const";

interface Props {
    id: string
}

export default function ProfileInfo(props: Props) {

    const {id} = props

    const [user, setUser] = useState<User>({})


    useEffect(() => {
        const fetch = async () => {
            try {
                const res = await $api.get('/users-me')

                setUser(res.data)
            } catch (e) {
                console.log(e)
            }
        }

        fetch()
    }, []);

    const router = useRouter()

    return (
        <div className={styles.container}>
            <Image src={bg} alt={""} className={styles.img}/>
            <div className={styles.bthInfo}>
                <div className={styles.info}>
                    <ProfileImage src={user.img_link}/>
                    <div className={styles.fioTag}>
                        <p>{user.name} {user.surname}</p>
                        <p className={styles.tag}>@{user.tag}</p>
                    </div>
                </div>
                <div className={styles.btnContainer}>
                    {user.id !== Number(id) && (
                        <Button onClick={() => router.push('/chats/1')}>
                            Сообщение
                        </Button>
                    )}
                    {user.id !== Number(id) && (
                        <Button className={styles.btn}>
                            <span className={clsx(
                                "material-symbols-rounded",
                                styles.icon
                            )}>
                                {true ? "person_remove" : "person_add"}
                            </span>
                        </Button>
                    )}
                    {user.id === Number(id) && (
                        <Button className={styles.btn}>
                            <span className={clsx(
                                "material-symbols-rounded",
                                styles.icon
                            )}>
                                settings
                            </span>
                        </Button>
                    )}
                </div>
            </div>

        </div>
    )
}
