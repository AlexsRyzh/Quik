'use client'

import styles from './post-card.module.scss'
import React, {createContext, useState} from "react";

interface ContextType {
    postID?: number,
    userID?: number,
    version?: number,
    setVersion?: any,
}

export const PostCardContext = createContext<ContextType>({})

interface Props {
    children: React.ReactNode
    value?: {
        postID: number,
        userID: number
    }
}

export default function PostCard(props: Props) {

    const {children, value} = props

    const [version, setVersion] = useState(0)


    return (
        <PostCardContext.Provider value={{...value, version, setVersion}}>
            <div className={styles.card}>
                {children}
            </div>
        </PostCardContext.Provider>
    )
}