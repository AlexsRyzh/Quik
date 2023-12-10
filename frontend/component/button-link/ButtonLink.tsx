import React from "react";
import clsx from "clsx";
import styles from "./button-linke.module.scss";
import Link from "next/link";

interface Props {
    className?: string,
    href: string,
    children?: React.ReactNode,
}

export function ButtonLink(props: Props) {

    const {
        href,
        children,
        className,
        ...otherProps
    } = props

    return (
        <Link
            className={clsx(
                styles.button,
                className
            )}
            href={href}
            {...otherProps}
        >
            {children}
        </Link>
    )
}