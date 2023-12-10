'use client'

import styles from './post-footer.module.scss'
import ToggleButton from "@/component/toggle-button/ToggleButton";
import {useCallback, useContext, useEffect, useState} from "react";
import Comment from "@/component/(post)/comment/Comment";
import CommentInput from "@/component/(post)/comment-input/CommentInput";
import $api from "@/http/api";
import {PostCardContext} from "@/component/(post)/post-card/PostCard";
import {toast} from "react-toastify";

export default function PostFooter() {

    const {postID} = useContext(PostCardContext)

    const [like, setLike] = useState({count: 0, isLike: false})
    const [comment, setComment] = useState({
        count: 0,
        comments: [],
    })
    const [commentIsOpen, setCommentIsOpen] = useState(false)
    const [refreshComment, setRefreshComment] = useState(false)


    const fetchLike = useCallback(async () => {
        try {
            const res = await $api.get(`/post-is-like/${postID}`)
            const resCount = await $api.get(`/posts/${postID}/count-likes`)
            setLike({count: resCount.data.count, isLike: res.data['is_like']})
        } catch (e) {
            console.log(e)
        }
    }, [])

    const fetchComment = useCallback(async () => {
        try {
            const resCount = await $api.get(`/posts/${postID}/count-comment`)
            const resComments = await $api.get(`/posts/${postID}/comments`)

            setComment({count: resCount.data.count, comments: resComments.data})
        } catch (e) {
            console.log(e)
        }

    }, [])


    useEffect(() => {
        fetchLike()
    }, []);

    useEffect(() => {
        fetchComment()
    }, [refreshComment]);

    const toggleLike = useCallback(async () => {
        try {
            if (!like.isLike) {
                let res = await $api.post(`/likes/${postID}`)
                setLike(prev => ({count: prev.count + 1, isLike: true}))
                toast.success("Лайк успешно поставлен")
            } else {
                let res = await $api.delete(`/likes/${postID}`)
                setLike(prev => ({count: prev.count - 1, isLike: false}))
                toast.success("Лайк успешно убран")
            }

            await fetchLike()
        } catch (e) {
            toast.error("Не удалось поставить лайк")
            console.log(e)
        }
    }, [like])

    return (
        <div className={styles.container}>
            <div className={styles.buttonContainer}>
                <ToggleButton
                    icon={"favorite"}
                    activeClassName={styles.activeLike}
                    noActiveClassName={styles.noActieLike}
                    value={like.isLike}
                    onClick={toggleLike}
                    count={like.count}
                />
                <ToggleButton
                    icon={"comment"}
                    activeClassName={styles.activeComment}
                    noActiveClassName={styles.noActiveComment}
                    value={commentIsOpen}
                    onClick={() => {
                        setCommentIsOpen((prev: boolean) => !prev)
                    }}
                    count={comment.count}
                />
            </div>
            {commentIsOpen && (
                <>
                    {comment.comments?.map((data: any) => {

                        console.log(data)
                        return (
                            <Comment {...data} />
                        )
                    })}
                    <CommentInput
                        refreshComment={() => setRefreshComment(prev => !prev)}
                    />
                </>
            )}
        </div>
    )
}