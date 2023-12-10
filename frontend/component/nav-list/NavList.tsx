'use client'

import styles from './nav-list.module.scss'
import {NavListItem} from "@/component/nav-list-item/NavListItem";
import {usePathname} from "next/navigation";
import clsx from "clsx";
import {v4 as uuidv4} from 'uuid';
import {useEffect, useState} from "react";
import {User} from "@/const/const";
import $api from "@/http/api";


export function NavList() {
    const pathname = usePathname()

    const [user, setUser] = useState<User>({})

    const linkList = [
        {name: 'Моя страница', href: `/profiles/${user.id}`, icon: 'home'},
        {name: 'Новости', href: '/', icon: 'feed'},
        {name: 'Друзья', href: '/friends', icon: 'group'},
        {name: 'Мессенджер', href: '/chats', icon: 'chat'},
    ]

    useEffect(() => {
        const fetch = async () => {
            try {
                const res = await $api.get('/users-me')
                setUser({...res.data})
            } catch (e) {
                console.log(e)
            }
        }
        fetch()
    }, []);

    return (
        <nav className={styles.container}>
            {linkList.map((item) => {

                return (
                    <NavListItem
                        key={uuidv4()}
                        href={item.href}
                        title={item.name}
                        icon={item.icon}
                        className={clsx(pathname === item.href ? styles.activeNavLink : '')}
                    />
                )
            })}
        </nav>
    )
}