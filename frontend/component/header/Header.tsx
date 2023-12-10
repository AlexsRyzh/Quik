'use client'

import styles from './header.module.scss'
import Link from "next/link";
import logo from '@/public/logo.png'
import Image from "next/image";
import {Button} from "@/component/button/Button";

export function Header() {
    

    return (
        <header className={styles.container}>
            <div className={styles.innerContainer}>
                <Link
                    href={"/"}
                    className={styles.logoContainer}
                >
                    <Image
                        src={logo}
                        alt={"Лого"}
                        className={styles.logo}
                    />
                    <h3 className={styles.logoText}>Quik</h3>
                </Link>
                <Button>
                    Выйти
                </Button>
            </div>
        </header>
    )
}