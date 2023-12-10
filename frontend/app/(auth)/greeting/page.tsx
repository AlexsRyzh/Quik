import {ButtonLink} from "@/component/button-link/ButtonLink";
import styles from './page.module.scss'

export default function Greeting() {
    return (
        <div className={styles.container}>
            <h2 className={styles.title}>
                Добро пожаловать
            </h2>
            <ButtonLink
                href={'/login'}
            >
                Вход
            </ButtonLink>
            <ButtonLink
                href={'/register'}
            >
                Регистрация
            </ButtonLink>
        </div>
    )
}