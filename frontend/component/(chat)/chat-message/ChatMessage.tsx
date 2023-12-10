import styles from './chat-message.module.scss'
import Image, {StaticImageData} from "next/image";
import svg1 from '@/public/Vectormessage-suffix.svg'
import svg2 from '@/public/Vector.svg'
import clsx from "clsx";

interface Props {
    message: string,
    img?: StaticImageData,
    from: number,
    date: Date,
    my: boolean
}

export default function ChatMessage(props: Props) {

    const {
        message,
        img,
        my
    } = props

    return (
        <div className={clsx(
            styles.container,
            my && styles.myContainer
        )}>
            {!my && (
                <Image src={svg1} alt={'Хвостик'}/>
            )}
            <div className={clsx(
                styles.blockMessage,
                my && styles.myMessage
            )}>
                <pre className={styles.message}>
                    {message}
                </pre>
                {img && (
                    <Image src={img} alt={""} className={styles.img}/>
                )}
            </div>
            {my && (
                <Image src={svg2} alt={'Хвостик'}/>
            )}
        </div>
    )
}