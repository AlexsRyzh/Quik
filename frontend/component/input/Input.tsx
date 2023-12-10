'use client'

import {ChangeEventHandler, forwardRef} from "react";
import clsx from "clsx";
import styles from './input.module.scss'

interface Props {
    type?: string,
    value?: string,
    onChange?: ChangeEventHandler,
    className?: string,
    errMessage?: string,
    placeholder: string
}

export const Input = forwardRef(function Input(props: Props, ref: any) {

    const {
        type,
        value,
        onChange,
        className,
        errMessage,
        ...otherProps
    } = props


    return (
        <div className={styles.container}>
            <input
                className={clsx(
                    styles.input,
                    errMessage && styles.inputErr,
                    className
                )}
                value={value}
                onChange={onChange}
                type={type && "text"}
                ref={ref}
                {...otherProps}
            />
            <p className={styles.err}>{errMessage || ""}</p>
        </div>
    )
})