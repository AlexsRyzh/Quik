import React from "react";
import clsx from "clsx";
import styles from './button.module.scss'

interface Props {
    className?: string,
    children?: React.ReactNode,
    onClick?: any
}

export function Button(props: Props) {

    const {
        children,
        className,
        ...otherProps
    } = props

    return (
        <button
            className={clsx(
                styles.button,
                className
            )}
            {...otherProps}
        >
            {children}
        </button>
    )
}