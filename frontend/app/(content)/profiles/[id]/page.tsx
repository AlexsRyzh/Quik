import React from "react";
import Profile from "@/component/(profile)/profile/Profile";


export default function ProfilesID({params}: { params: { id: string } }) {
    return (
        <Profile id={params.id}/>
    )
}