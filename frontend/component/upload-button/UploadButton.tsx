'use client'

import styles from './uplaod-file.module.scss'
import clsx from "clsx";
import {ChangeEvent} from "react";

interface Props {
    onChange: (e: ChangeEvent<HTMLInputElement>) => void
}

export default function UploadButton(props: Props) {

    const {
        onChange
    } = props

    return (
        <label className={styles.label}>
            <input
                type="file"
                accept="image/png,.jpg,.jpeg"
                className={styles.container}
                onChange={onChange}
            />
            <span className={clsx(
                "material-symbols-rounded",
                styles.icon
            )}>
                attach_file
            </span>
        </label>
    )
}