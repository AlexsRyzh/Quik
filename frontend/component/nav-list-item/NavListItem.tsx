import Link from "next/link";
import styles from './nav-list-item.module.scss'
import clsx from "clsx";

interface Props {
    title: string,
    icon: string,
    href: string,
    className?: string,
}

export function NavListItem(props: Props) {

    const {
        href,
        title,
        icon,
        className
    } = props
    
    return (
        <Link href={href} className={clsx(
            styles.container,
            className
        )}>
            <span className={clsx(
                styles.materialSymbolsRounded,
                "material-symbols-rounded"
            )}>
                {icon}
            </span>
            <p>{title}</p>
        </Link>
    )
}