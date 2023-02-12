import {useEffect, useState} from "react"
import {Event} from "./Events";

const api_host = process.env.REACT_APP_API_HOST as string
const ws_api_host = process.env.REACT_APP_WS_API_HOST as string

const fibNumbers = [0, 1, 2, 3, 5, 8, 13, 21, 25]

type Session = {
    client_id: string
    username: string
    value: string
}

export default function Lobby() {
    const [registered, setRegistered] = useState<boolean>(false)
    const [ws, setWs] = useState<WebSocket | null>(null)

    const [userId, setUserId] = useState<string | null>(null)
    const [sessions, setSessions] = useState<Session[]>([])

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

        const websocket = new WebSocket(ws_api_host + "/ws")
        setWs(websocket)

        websocket.onopen = () => {
            console.log('connected')
        }

        websocket.onclose = () => {
            console.log('got closed')
            setSessions([])
            setUserId("-1")
        }

        websocket.onerror = () => {
            console.log("error")
            setSessions([])
            setUserId("-1")
        }

        return () => {
            websocket.close()
        }
    }, [registered])

    if (ws !== null) {
        ws.onmessage = (event) => {
            const data = JSON.parse(event.data)
            console.log(data)
            switch (data.event_type) {
                case Event.welcome_event:
                    setUserId(data.session.client_id)
                    addSessions(data.sessions)
                    break
                case Event.join_event:
                    addClient(data.session)
                    break
                case Event.leave_event:
                    removeClient(data.session)
                    break
                case Event.session_change_event:
                    updateClient(data.session)
                    break
            }
        }
    }

    const addSessions = (newSessions: Session[]) => {
        setSessions([...sessions, ...newSessions])
    }

    const addClient = (session: Session) => {
        setSessions([...sessions, session])
    }

    const updateClient = (session: Session) => {
        // console.log(session, sessions)
        setSessions(sessions.map((c) => c.client_id === session.client_id ? session : c))
    }

    const removeClient = (session: Session) => {
        setSessions(sessions.filter((s) => s.client_id !== session.client_id))
    }

    const pick = (value: number) => {
        if (ws !== null) {
            ws.send(JSON.stringify({
                event_type: "pick_event",
                value: value,
            }))
        }
    }

    return (
        <main>
            <h1>Hello {userId}</h1>

            <br/>
            <h1>users also here:</h1>
            <div>
                {sessions.filter((client) => client.client_id !== userId).map((client, index) => {
                    return <p key={index}>{client.client_id} {client.value}</p>
                })}
            </div>

            {fibNumbers.map((n, index) => <Value key={index} value={n} pickFn={pick}/>)}
        </main>
    )
}

function Value(props: { pickFn: (v: number) => void, value: number }) {
    return <button
        onClick={() => props.pickFn(props.value)}
        className="rounded-lg bg-emerald-600 px-4 py-1.5 text-base font-semibold text-white shadow-sm ring-1 ring-emerald-600 hover:bg-emerald-700 hover:ring-emerald-700"
    >{props.value}</button>
}
