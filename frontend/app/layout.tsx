import type {Metadata} from 'next'
import React from "react";
import {Header} from "@/component/header/Header";
import Content from "@/component/content/Content";
import './reset.scss'
import './global.scss'
import {ToastContainer} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import AuthContext from "@/contex/auth/AuthContext";


export const metadata: Metadata = {
    title: 'Quik',
    description: 'Quik - Социальная сеть',
}

export default function RootLayout({
    children
}: {
    children: React.ReactNode
}) {


    return (
        <html lang="en">
            <body>
                <Header/>
                <AuthContext>
                    <Content>
                        {children}
                    </Content>
                    <ToastContainer autoClose={1000}/>
                </AuthContext>
            </body>
        </html>
    )
}
