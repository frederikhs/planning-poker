import {Client} from "../type";
import {CheckIcon, EllipsisHorizontalIcon} from "@heroicons/react/24/solid";
import {useMemo} from "react";

export default function ClientList(props: { clients: Client[], thisClient: Client, valuesVisible: boolean }) {
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
    }, [props.clients])

    return (
        <div className={"p-4 space-y-2"}>
            <DisplayClient client={props.thisClient} thisClient={true} valuesVisible={props.valuesVisible}/>

            {sortedClients.map((client, index) => {
                return <DisplayClient key={index} client={client} thisClient={false} valuesVisible={props.valuesVisible}/>
            })}
        </div>
    )
}

function DisplayClient(props: { client: Client, thisClient: boolean, valuesVisible: boolean }) {
    return (
        <div className={`flex justify-between shadow-lg text-white ${props.thisClient ? 'bg-green-600' : 'bg-gray-600'} rounded-md text-3xl px-4 py-2`}>
            <p>{props.client.username}</p>
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
