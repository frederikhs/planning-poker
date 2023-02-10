import {useEffect, useState} from "react";

const api_host = process.env.REACT_APP_API_HOST as string
const ws_api_host = process.env.REACT_APP_WS_API_HOST as string

type Client = {
    id: string
    username: string
    value: string
}

export default function Lobby() {
    const [registered, setRegistered] = useState<boolean>(false);

    const [userId, setUserId] = useState<string | null>(null);
    const [clients, setClients] = useState<Client[]>([]);

    useEffect(() => {
        fetch(api_host + "/register", {
            credentials: "include"
        }).then((res) => {
            if (res.status === 201) {
                console.log("registered ok")
                setRegistered(true)
            } else {
                console.log("registered false")
                setRegistered(false)
            }
        })
    }, [])

    useEffect(() => {
        if (!registered) {
            console.log("not registered yet")
            return
        }

        const websocket = new WebSocket(ws_api_host + "/ws/lobby1");

        websocket.onopen = () => {
            console.log('connected');
        }

        websocket.onclose = () => {
            console.log('got closed');
            setClients([])
            setUserId("-1")
        }

        websocket.onerror = () => {
            console.log("error")
            setUserId("-1")
        }

        websocket.onmessage = (event) => {
            const data = JSON.parse(event.data);
            console.log(data);
            // setUserId(data.client_id)

            switch (data.event_type) {
                case "welcome_event":
                    setClients(data.clients)
                    setUserId(data.client.id)
                    break
                case "join_event":
                    setClients([...clients, data.client])
                    break
                case "leave_event":
                    setClients(clients.filter((c) => c.id !== data.client.id))
                    break
            }
        }

        return () => {
            websocket.close()
        }
    }, [registered])

    return (
        <main>
            <h1>Hello {userId}</h1>

            <br/>
            <h1>users also here:</h1>
            <div>
                {clients.map((client, index) => {
                    return <p key={index}>{client.id} {client.value}</p>
                })}
            </div>
        </main>
    )
}
