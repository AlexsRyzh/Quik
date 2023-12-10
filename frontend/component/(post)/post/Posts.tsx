'use client'

import React, {useEffect, useState} from "react";
import $api from "@/http/api";
import PostCard from "@/component/(post)/post-card/PostCard";
import PostHeader from "@/component/(post)/post-header/PostHeader";
import PostText from "@/component/(post)/post-text/PostText";
import PostImage from "@/component/(post)/post-image/PostImage";
import styles from './post.module.scss'
import PostFooter from "@/component/(post)/post-footer/PostFooter";

export default function Posts() {

    const [ids, setIDs] = useState([])
    const [userIDs, setUseIDs] = useState([])

    useEffect(() => {
        const fetch = async () => {
            try {
                const res = await $api.get('/posts-ids')

                setIDs(res.data["ids"])
                setUseIDs(res.data['user_id'])

            } catch (e) {
                console.log(e)
            }
        }
        fetch()
    }, []);

    console.log(ids)

    return (
        <div className={styles.container}>
            {ids.map((value, index) => {
                return (
                    <PostCard value={{postID: value, userID: userIDs[index]}}>
                        <PostHeader/>
                        <PostText/>
                        <PostImage/>
                        <PostFooter/>
                    </PostCard>
                )
            })}
        </div>
    )
}