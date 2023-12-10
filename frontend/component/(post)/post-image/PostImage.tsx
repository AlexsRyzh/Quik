'use client'

import styles from './post-image.module.scss'
import {useContext, useEffect, useState} from "react";
import {PostCardContext} from "@/component/(post)/post-card/PostCard";
import $api from "@/http/api";
import {FILE_UPLOAD} from "@/const/const";

export default function PostImage() {

    const {postID} = useContext(PostCardContext)

    const [imgName, setImgName] = useState("")

    useEffect(() => {
        const fetch = async () => {
            try {
                const res = await $api.get(`/post-img/${postID}`)

                setImgName(res.data)
            } catch (e) {
                console.log(e)
            }
        }
        fetch()
    }, []);

    return (
        <>
            {imgName !== "" ? (
                <div className={styles.container}>

                    <img src={FILE_UPLOAD + imgName} className={styles.img} alt={"Фото"}/>
                </div>
            ) : null}
        </>

    )
}