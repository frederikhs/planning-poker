import {useEffect, useMemo, useState} from "react"
import {Event} from "../Events";
import {useParams} from "react-router-dom";
import ValuePicker from "../Components/ValuePicker";
import {Client} from "../type";
import ClientList from "../Components/ClientList";
import ValueDisplay from "../Components/ValueDisplay";

const api_host = process.env.REACT_APP_API_HOST as string
const ws_api_host = process.env.REACT_APP_WS_API_HOST as string

const fibNumbers = [0, 1, 3, 5, 8, 13, 21]

export default function Lobby() {
    let {lobby_id} = useParams()

    const [registered, setRegistered] = useState<boolean>(false)
    const [ws, setWs] = useState<WebSocket | null>(null)

    const [thisClient, setThisClient] = useState<Client | null>(null)
    const [clients, setClients] = useState<Client[]>([])

    const [valuesVisible, setValuesVisible] = useState<boolean>(false)
    const [pickedValue, setPickedValue] = useState<number | null>(null)

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

        const websocket = new WebSocket(ws_api_host + "/ws/" + lobby_id)
        setWs(websocket)

        websocket.onopen = () => {
            console.log('connected')
        }

        websocket.onclose = () => {
            console.log('got closed')
            setClients([])
            setThisClient(null)
        }

        websocket.onerror = () => {
            console.log("error")
            setClients([])
            setThisClient(null)
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
                    setThisClient(data.session)
                    addSessions(data.sessions)
                    setValuesVisible(data.visible)
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
                case Event.toggle_visibility_event:
                    setValuesVisible(data.visible)
                    break
            }
        }
    }

    const addSessions = (newSessions: Client[]) => {
        setClients([...clients, ...newSessions])
    }

    const addClient = (session: Client) => {
        setClients([...clients, session])
    }

    const updateClient = (session: Client) => {
        if (session.client_id === thisClient?.client_id) {
            setThisClient(session)
        } else {
            setClients(clients.map((c) => c.client_id === session.client_id ? session : c))
        }
    }

    const removeClient = (session: Client) => {
        setClients(clients.filter((s) => s.client_id !== session.client_id))
    }

    const answerValues = useMemo(() => {
        if (thisClient === null) {
            return [-1]
        }

        return [thisClient.value, ...clients.map((value) => value.value)]
    }, [clients, thisClient])

    const send = (object: Object) => {
        if (ws === null) {
            throw "no websocket connection"
        }

        console.log(object);

        ws.send(JSON.stringify(object))
    }

    const pick = (value: number) => {
        setPickedValue(value)
        send({
            event_type: "pick_event",
            value: value,
        })
    }

    const toggleVisibility = () => {
        send({event_type: "toggle_visibility_request_event"})
    }

    if (thisClient == null) {
        return <p>Loading</p>
    }

    return (
        <main>
            <ValueDisplay values={answerValues} valuesVisible={valuesVisible} toggleVisibilityFn={toggleVisibility}/>

            <ClientList clients={clients} thisClient={thisClient} valuesVisible={valuesVisible}/>

            <ValuePicker values={fibNumbers} pickFn={pick} pickedValue={pickedValue}/>
        </main>
    )
}
