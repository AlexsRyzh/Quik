import React from "react";
import styles from "./layout.module.scss"
import Image from "next/image";
import logo from '@/public/logo-max.png'
import Link from "next/link";

export default function AuthLayout({
    children
}: {
    children: React.ReactNode
}) {
    return (
        <div className={styles.outer}>
            <div className={styles.container}>
                <div className={styles.logoContainer}>
                    <Link
                        href={'/'}
                        className={styles.link}
                    >
                        <Image
                            src={logo}
                            alt={"logo"}
                            className={styles.img}
                        />
                        <h2 className={styles.title}>Quik</h2>
                    </Link>
                </div>
                <div className={styles.mainContainer}>
                    {children}
                </div>
            </div>
        </div>
    )
}
