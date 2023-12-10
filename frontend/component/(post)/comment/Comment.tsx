import styles from './comment.module.scss'
import Image from "next/image";
import img from '@/public/profile-img.png'

interface Props {
    name: string,
    surname: string,
    message: string,
    user_img_link: string
}

export default function Comment(props: Props) {

    const {
        name,
        surname,
        message,
        user_img_link
    } = props

    return (
        <div className={styles.container}>
            <Image src={user_img_link === "" ? img : user_img_link} alt={'Фото'} className={styles.img}/>
            <div className={styles.content}>
                <p className={styles.fio}>{name} {surname}</p>
                <pre className={styles.comment}>
                    {message}
                </pre>
            </div>
        </div>
    )
}