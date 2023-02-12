import {useEffect, useState} from "react";

const api_host = process.env.REACT_APP_API_HOST as string
const ws_api_host = process.env.REACT_APP_WS_API_HOST as string

type Session = {
    client_id: string
    username: string
    value: string
}

export default function Lobby() {
    const [registered, setRegistered] = useState<boolean>(false);

    const [userId, setUserId] = useState<string | null>(null);
    const [sessions, setSessions] = useState<Session[]>([]);

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

        const websocket = new WebSocket(ws_api_host + "/ws");

        websocket.onopen = () => {
            console.log('connected');
        }

        websocket.onclose = () => {
            console.log('got closed');
            setSessions([])
            setUserId("-1")
        }

        websocket.onerror = () => {
            console.log("error")
            setUserId("-1")
        }

        websocket.onmessage = (event) => {
            const data = JSON.parse(event.data);
            console.log(data);
            switch (data.event_type) {
                case "welcome_event":
                    setSessions(data.sessions)
                    setUserId(data.session.client_id)
                    break
                case "join_event":
                    setSessions([...sessions, data.session])
                    break
                case "leave_event":
                    setSessions(sessions.filter((c) => c.client_id !== data.session.client_id))
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
                {sessions
                    .filter((client) => client.client_id !== userId)
                    .map((client, index) => {
                    return <p key={index}>{client.client_id} {client.value}</p>
                })}
            </div>
        </main>
    )
}
