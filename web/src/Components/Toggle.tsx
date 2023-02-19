import {ArchiveBoxXMarkIcon, TvIcon} from "@heroicons/react/24/outline";
import {useNavigate} from "react-router-dom";
import {ChevronDoubleLeftIcon} from "@heroicons/react/20/solid";
import {TvIcon as TvIconSolid} from "@heroicons/react/24/solid";

export function ViewerToggle(props: { toggleFn: () => void, active: boolean }) {
    return (
        <div onClick={() => props.toggleFn()} className={"p-2 rounded-md shadow-lg hover:cursor-pointer"}>
            {props.active && <TvIconSolid className={"w-8 h-8"}/>}
            {!props.active && <TvIcon className={"w-8 h-8"}/>}
        </div>
    )
}

export function ClearToggle(props: { clearFn: () => void }) {
    return (
        <div onClick={() => props.clearFn()} className={"p-2 rounded-md shadow-lg hover:cursor-pointer"}>
            <ArchiveBoxXMarkIcon className={"w-8 h-8"}/>
        </div>
    )
}

export function GoToLobbyToggle() {
    const navigate = useNavigate()

    return (
        <div onClick={() => navigate("/")} className={"p-2 rounded-md shadow-lg hover:cursor-pointer"}>
            <ChevronDoubleLeftIcon className={"w-8 h-8"}/>
        </div>
    )
}
