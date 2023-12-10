import styles from './post-text.module.scss'
import {useContext, useEffect, useState} from "react";
import {PostCardContext} from "@/component/(post)/post-card/PostCard";
import $api from "@/http/api";


export default function PostText() {

    const {postID} = useContext(PostCardContext)

    const [text, setText] = useState("")

    useEffect(() => {
        const fetch = async () => {
            try {
                const res = await $api.get(`/post-text/${postID}`)

                setText(res.data)
            } catch (e) {
                console.log(e)
            }
        }

        fetch()
    }, []);

    return (
        <pre className={styles.text}>
            {text}
        </pre>
    )
}