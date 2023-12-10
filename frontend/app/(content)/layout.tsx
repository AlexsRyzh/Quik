'use client'

import React from "react";
import styles from './content-layout.module.scss'
import {NavList} from "@/component/nav-list/NavList";
import MiniProfile from "@/component/mini-profile/MiniProfile";

interface Props {
    children: React.ReactNode
}

export default function RootLayout(props: Props) {


    const {children} = props

    return (
        <div className={styles.container}>
            <section className={styles.nav}>
                <MiniProfile/>
                <NavList/>
            </section>
            <section className={styles.content}>
                {children}
            </section>
        </div>
    )
}