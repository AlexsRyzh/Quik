import TextareaAutosize from 'react-textarea-autosize';
import styles from './textarea.module.scss'
import clsx from "clsx";

interface Props {
    placeholder?: string,
    className?: string,
    onChange?: any,
    value: string
}


export default function Textarea(props: Props) {

    const {className, onChange, ...otherProps} = props

    return (
        <TextareaAutosize
            className={
                clsx(
                    styles.textarea,
                    className
                )
            }
            onChange={onChange}
            {...otherProps}
        />
    )

}