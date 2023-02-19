import {ArchiveBoxXMarkIcon} from "@heroicons/react/24/outline";

export default function Clear(props: { enable: boolean, clearFn: () => void }) {
    if (!props.enable) {
        return null
    }

    return (
        <div onClick={() => props.clearFn()} className={"fixed top-0 right-0 m-4 p-2 rounded-md shadow-lg hover:cursor-pointer"}>
            <ArchiveBoxXMarkIcon className={"w-8 h-8"}/>
        </div>
    )
}
