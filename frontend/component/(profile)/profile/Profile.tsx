'use client'

import React, {useEffect, useState} from "react";
import $api from "@/http/api";
import PostCard from "@/component/(post)/post-card/PostCard";
import PostHeader from "@/component/(post)/post-header/PostHeader";
import PostText from "@/component/(post)/post-text/PostText";
import PostImage from "@/component/(post)/post-image/PostImage";
import styles from './profile.module.scss'
import PostFooter from "@/component/(post)/post-footer/PostFooter";
import ProfileInfo from "@/component/(profile)/profile-info/ProfileInfo";

interface Props {
    id: string
}

export default function Profile(props: Props) {

    const {id} = props

    const [ids, setIDs] = useState([])
    const [userIDs, setUseIDs] = useState([])

    useEffect(() => {
        const fetch = async () => {
            try {
                const res = await $api.get('/posts-ids-my')

                setIDs(res.data["ids"])
                setUseIDs(res.data['user_id'])

            } catch (e) {
                console.log(e)
            }
        }
        fetch()
    }, []);


    return (
        <>
            <ProfileInfo id={id}/>
            <div className={styles.postContainer}>
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
        </>
    )
}