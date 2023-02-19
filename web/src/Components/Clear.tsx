import {ArchiveBoxXMarkIcon} from "@heroicons/react/24/outline";

export default function Clear(props: { clearFn: () => void }) {
    return (
        <div onClick={() => props.clearFn()} className={"p-2 rounded-md shadow-lg hover:cursor-pointer"}>
            <ArchiveBoxXMarkIcon className={"w-8 h-8"}/>
        </div>
    )
}
