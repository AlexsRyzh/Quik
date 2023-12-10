import Chats from "@/component/(chat)/chats/Chats";
import ChatsList from "@/component/(chat)/chats-list/ChatsList";
import img from '@/public/profile-img.png'

const chatList = [
    {
        idUser: 1,
        name: "Piter",
        img: img,
        surname: "Jackson",
        lastMessage: {
            text: "Горы всегда притягивали людей своей величественностью, красотой и загадочностью. Каждая горная вершина - это уникальный мир",
            date: new Date()
        }
    },
    {
        idUser: 2,
        name: "Piter",
        img: img,
        surname: "Jackson",
        lastMessage: {
            text: "Горы всегда притягивали людей своей величественностью, красотой и загадочностью. Каждая горная вершина - это уникальный мир",
            date: new Date()
        }
    },
    {
        idUser: 3,
        name: "Piter",
        img: img,
        surname: "Jackson",
        lastMessage: {
            text: "Горы всегда притягивали людей своей величественностью, красотой и загадочностью. Каждая горная вершина - это уникальный мир",
            date: new Date()
        }
    },
]

export default function ChatListPage() {
    return (
        <Chats chatList={chatList}>
            <ChatsList/>
        </Chats>
    )
}