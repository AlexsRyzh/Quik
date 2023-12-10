'use client'

import styles from './toggle-button.module.scss'
import clsx from "clsx";

interface Props {
    activeClassName: string,
    noActiveClassName: string,
    icon: string,
    value: boolean,
    onClick: any,
    count: number
}

export default function LikesButton(props: Props) {

    const {
        activeClassName,
        noActiveClassName,
        icon,
        value,
        onClick,
        count
    } = props
    

    return (
        <button
            onClick={onClick}
            className={clsx(
                styles.button,
                noActiveClassName,
                value && activeClassName
            )}
        >
            <span className={clsx(
                "material-symbols-rounded",
                styles.icon,
            )}>
                {icon}
            </span>
            <p className={styles.text}>{count}</p>
        </button>
    )
}