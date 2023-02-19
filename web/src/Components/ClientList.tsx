import {Client} from "../type";
import {CheckIcon, EllipsisHorizontalIcon} from "@heroicons/react/24/solid";
import {useMemo, useState} from "react";

export default function ClientList(props: { clients: Client[], thisClient: Client, setUsernameFn: (v: string) => void, valuesVisible: boolean }) {
    const excludeViewerClients = useMemo(() => {
        return props.clients.filter((client) => !client.viewer)
    }, [props.clients])

    const sortedClients = useMemo(() => {
        return props.clients.sort((a, b) => {
            if (a.client_id > b.client_id) {
                return -1
            } else if (a.client_id < b.client_id) {
                return 1
            } else {
                return 0
            }
        })
    }, [excludeViewerClients])

    const nonViewerClients = useMemo(() => {
        return sortedClients.filter((client) => !client.viewer)
    }, [sortedClients])

    return (
        <div className={"p-4 space-y-2"}>
            {!props.thisClient.viewer &&
                <DisplayClient client={props.thisClient} setUsernameFn={props.setUsernameFn} thisClient={true} valuesVisible={props.valuesVisible}/>
            }

            {nonViewerClients.map((client, index) => {
                return <DisplayClient key={index} client={client} setUsernameFn={() => {
                }} thisClient={false} valuesVisible={props.valuesVisible}/>
            })}

            {((nonViewerClients.length === 0 && props.thisClient.viewer) || props.clients.length === 0) && <p className={"text-center"}>
                Seems like there no one here to play. Invite others by sending them the link
                <br/>
                <span className={"font-bold"}>{window.location.href}</span>
            </p>}
        </div>
    )
}

function DisplayClient(props: { client: Client, setUsernameFn: (v: string) => void, thisClient: boolean, valuesVisible: boolean }) {
    const [username, setUsername] = useState<string>(props.client.username)

    return (
        <div className={`flex justify-between shadow-lg text-white ${props.thisClient ? 'bg-green-600' : 'bg-gray-600'} rounded-md text-3xl px-4 py-2`}>

            {props.thisClient && <input
                className={"bg-transparent block text-white focus:outline-none w-4/5"}
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                onKeyUp={() => props.setUsernameFn(username)}
            />
            }

            {!props.thisClient && <p className={"break-all w-4/5"}>{props.client.username}</p>}

            {props.client.value === -1 && <div className={"flex align-items animate-pulse"}>
                <EllipsisHorizontalIcon className={"w-5"}/>
            </div>}
            {(props.client.value !== -1 && !props.valuesVisible) && <div className={"flex align-items"}>
                <CheckIcon className={"w-5"}/>
            </div>}
            {props.client.value !== -1 && props.valuesVisible && <p>
                {props.client.value}
            </p>}
        </div>
    )
}
