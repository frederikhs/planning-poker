import React, {useState} from "react";

const api_host = process.env.REACT_APP_API_HOST as string
const ws_api_host = process.env.REACT_APP_WS_API_HOST as string

type Client = {
    id: string
    username: string
    value: string
}

export default function Lobby() {
    const [userId, setUserId] = useState<string | null>(null);
    const [clients, setClients] = useState<Client[] | null>(null);

    React.useEffect(() => {
        const websocket = new WebSocket(ws_api_host + "/ws/lobby1");

        websocket.onopen = () => {
            console.log('connected');
        }

        websocket.onmessage = (event) => {
            const data = JSON.parse(event.data);
            console.log(data);
            setUserId(data.client_id)
            setClients(data.clients)
        }

        return () => {
            websocket.close()
        }
    }, [])

    return (
        <main>
            <h1>Hello {userId}</h1>

            <div>
                {clients?.map((client, index) => {
                    return <p key={index}>{client.id} {client.value}</p>
                })}
            </div>
        </main>
    )
}
