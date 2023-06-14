import React, { useEffect, useMemo, useState } from 'react'
import { Event } from '../Events'
import { useParams } from 'react-router-dom'
import ValuePicker from '../Components/ValuePicker'
import { type Client } from '../type'
import ClientList from '../Components/ClientList'
import ValueDisplay from '../Components/ValueDisplay'
import { ClearToggle, GoToLobbyToggle, ViewerToggle } from '../Components/Toggle'
import Error from '../Components/Error'

const API_HOST = process.env.REACT_APP_API_HOST as string
const WS_API_HOST = process.env.REACT_APP_WS_API_HOST as string

const fibNumbers = [0, 0.5, 1, 2, 3, 5, 8, 13]

export default function Lobby (): JSX.Element {
  const { lobbyId } = useParams()
  const [registered, setRegistered] = useState<boolean>(false)
  const [ws, setWs] = useState<WebSocket | null>(null)
  const [thisClient, setThisClient] = useState<Client | null>(null)
  const [clients, setClients] = useState<Client[]>([])
  const [valuesVisible, setValuesVisible] = useState<boolean>(false)

  useEffect(() => {
    fetch(API_HOST + '/register', {
      credentials: 'include'
    }).then((res) => {
      if (res.status === 201) {
        console.log('registered ok')
        setRegistered(true)
      } else {
        console.log('registered false')
        setRegistered(false)
      }
    }).catch((e) => {
      console.error(e)
    })
  }, [])

  const wsConnect = (): WebSocket => {
    const websocket = new WebSocket(`${WS_API_HOST}/ws/${lobbyId as string}`)

    websocket.onopen = () => {
      console.log('connected')
    }

    websocket.onclose = () => {
      console.log('got closed')
      setClients([])
      setThisClient(null)
    }

    websocket.onerror = () => {
      console.log('error')
      setClients([])
      setThisClient(null)
    }

    return websocket
  }

  useEffect(() => {
    if (!registered) {
      console.log('not registered yet')
      return
    }

    const websocket = wsConnect()
    setWs(websocket)

    return () => {
      setWs(null)
      websocket.close()
    }
  }, [registered, lobbyId])

  if (ws !== null) {
    ws.onmessage = (event) => {
      const data = JSON.parse(event.data)
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
        case Event.clear_lobby_event:
          clearClientValues()
          break
      }
    }
  }

  const addSessions = (newSessions: Client[]): void => {
    setClients([...clients, ...newSessions])
  }

  const addClient = (session: Client): void => {
    setClients([...clients, session])
  }

  const updateClient = (session: Client): void => {
    if (session.client_id === thisClient?.client_id) {
      setThisClient(session)
    } else {
      setClients(clients.map((c) => c.client_id === session.client_id ? session : c))
    }
  }

  const removeClient = (session: Client): void => {
    setClients(clients.filter((s) => s.client_id !== session.client_id))
  }

  const clearClientValues = (): void => {
    if (thisClient !== null) {
      thisClient.value = -1
      setThisClient(thisClient)
    }

    setValuesVisible(false)
    setClients(clients.map((c) => {
      c.value = -1
      return c
    }))
  }

  // this memo is here because there is a leak if client join and we will then filter out the clients with duplicate client ids
  const noDuplicateClients = useMemo(() => {
    return clients.filter((value, index, self) => self.map((c) => c.client_id).indexOf(value.client_id) === index)
  }, [clients])

  const answerValues = useMemo(() => {
    if (thisClient === null) {
      return [-1]
    }

    const otherClients = noDuplicateClients.filter((client) => !client.viewer).map((value) => value.value)

    if (thisClient.viewer) {
      return otherClients
    } else {
      return [thisClient.value, ...otherClients]
    }
  }, [noDuplicateClients, thisClient])

  const send = (object: any): void => {
    if (ws === null) {
      return
    }

    ws.send(JSON.stringify(object))
  }

  const pick = (value: number): void => {
    if (thisClient === null) {
      return
    }

    send({
      event_type: Event.pick_event,
      value
    })
  }

  const updateUsername = (username: string): void => {
    send({
      event_type: Event.choose_username_event,
      username
    })
  }

  const toggleVisibility = (): void => {
    send({
      event_type: Event.toggle_visibility_request_event
    })
  }

  const clearValues = (): void => {
    send({
      event_type: Event.clear_lobby_event
    })
  }

  const toggleViewer = (): void => {
    if (thisClient === null) {
      return
    }

    send({
      event_type: Event.toggle_viewer_request_event,
      viewer: !thisClient.viewer
    })
  }

  if (thisClient == null) {
    return <Error/>
  }

  return (
        <main className={'flex flex-col h-max'}>
            <div className={'flex justify-between sm:justify-start m-4'}>
                <GoToLobbyToggle/>
                {(!thisClient.viewer && valuesVisible) && <ClearToggle clearFn={clearValues}/>}
                <ViewerToggle toggleFn={toggleViewer} active={thisClient.viewer}/>
            </div>

            <ValueDisplay values={fibNumbers} answerValues={answerValues} valuesVisible={valuesVisible} toggleVisibilityFn={toggleVisibility}/>

            <ClientList clients={noDuplicateClients} thisClient={thisClient} valuesVisible={valuesVisible} setUsernameFn={updateUsername}/>

            {!thisClient.viewer && <ValuePicker values={fibNumbers} pickFn={pick} pickedValue={thisClient.value}/>}
        </main>
  )
}
