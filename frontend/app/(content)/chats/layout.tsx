import React from "react";
import styles from './chats-layout.module.scss'

interface Props {
    children: React.ReactNode
}

export default function Layout(props: Props) {

    const {children} = props

    return (
        <div className={styles.container}>
            {children}
        </div>
    )
}