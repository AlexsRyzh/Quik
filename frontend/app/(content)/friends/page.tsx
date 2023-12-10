import Friend from "@/component/(friend)/friend/Friend";
import img from '@/public/profile-img.png'
import FriendList from "@/component/(friend)/friend-list/FriendList";

const friends = [
    {
        img: img,
        name: "Piter",
        surname: "Jackson",
        userID: 1
    },
    {
        img: img,
        name: "Piter",
        surname: "Jackson",
        userID: 2
    },
    {
        img: img,
        name: "Piter",
        surname: "Jackson",
        userID: 3
    },
    {
        img: img,
        name: "Piter",
        surname: "Jackson",
        userID: 4
    },
]

export default function Friends() {
    return (
        <Friend friendList={friends}>
            <FriendList/>
        </Friend>
    )
}