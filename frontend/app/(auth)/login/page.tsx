'use client'

import {Input} from "@/component/input/Input";
import React, {useCallback} from "react";
import styles from './login.module.scss'
import {Button} from "@/component/button/Button";
import {useForm} from "react-hook-form";
import Link from "next/link";
import $api from "@/http/api";
import {useRouter} from "next/navigation";
import {toast} from "react-toastify";

interface FormInput {
    login: string
    password: string
}

export default function Login() {

    const {
        register,
        handleSubmit,
        formState: {errors}
    } = useForm<FormInput>()

    const router = useRouter()

    const onSubmit = useCallback(async (data: FormInput) => {
        try {
            const result = await $api.post('/auth/login', {...data})

            window.localStorage.setItem('access_token', result.data['access_token'])
            window.localStorage.setItem('refresh_token', result.data['refresh_token'])
            window.localStorage.setItem('user_id', result.data['user_id'])

            router.push('/')
        } catch (e) {
            console.log(e)

            toast.error("Не верный логин или пароль")
        }
    }, [])

    return (
        <div className={styles.container}>
            <h2 className={styles.title}>Вход</h2>
            <form
                className={styles.form}
                onSubmit={handleSubmit(onSubmit)}
            >
                <Input
                    placeholder="Логин"
                    {...register("login", {required: "Не может быть пустым"})}
                    errMessage={errors?.login?.message}
                />

                <Input
                    placeholder="Пароль"
                    {...register("password", {required: "Не может быть пустым"})}
                    errMessage={errors?.password?.message}
                />
                <Button>
                    Войти
                </Button>
            </form>
            <p
                className={styles.bottomText}>Нет аккаунта?
                <Link href={'/register'} className={styles.link}>Зарегистрироваться</Link>
            </p>
        </div>
    )
}