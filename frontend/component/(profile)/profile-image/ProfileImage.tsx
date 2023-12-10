'use client'

import Image from "next/image";
import styles from './profile-image.module.scss'
import clsx from "clsx";
import img from '@/public/profile-img.png'

interface Props {
    src?: string,
    online?: boolean
}

export default function ProfileImage(props: Props) {

    const {src, online} = props

    return (
        <div className={styles.container}>
            <Image src={src === "" ? img : src || ""} alt={""} className={styles.img}/>
            <div className={styles.status}>
                <div className={clsx(
                    styles.offline,
                    online && styles.online,
                )}></div>
            </div>
        </div>
    )
}