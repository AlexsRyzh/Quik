'use client'

import styles from './chat-input.module.scss'
import UploadButton from "@/component/upload-button/UploadButton";
import {useState} from "react";
import Textarea from "@/component/(post)/textarea/Textarea";
import {Button} from "@/component/button/Button";
import ImageView from "@/component/image-view/ImageView";
import img from '@/public/profile-img.png'
import clsx from "clsx";


export default function ChatInput() {

    const [image, setImage] = useState<File>()

    return (
        <div className={styles.container}>
            <div className={styles.innerContainer}>
                <UploadButton
                    onChange={(e) => {
                        if (e.target.files) {
                            setImage(e.target.files[0]);
                        }
                    }}
                />
                <div className={styles.textImg}>
                    <Textarea
                        placeholder={"Написать сообщения"}
                    />
                    {image && (
                        <ImageView imgLink={img} onClick={() => setImage(undefined)}/>
                    )}
                </div>
                <Button className={styles.btn}>
                    <span className={clsx(
                        "material-symbols-rounded",
                        styles.icon
                    )}>
                        send
                    </span>
                </Button>
            </div>
        </div>
    )
}
