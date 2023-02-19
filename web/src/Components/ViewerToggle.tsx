import {ComputerDesktopIcon} from "@heroicons/react/24/solid";
import {TvIcon} from "@heroicons/react/24/outline";

export default function ViewerToggle(props: { toggleFn: () => void }) {
    return (
        <div onClick={() => props.toggleFn()} className={"p-2 rounded-md shadow-lg hover:cursor-pointer"}>
            <TvIcon className={"w-8 h-8"}/>
        </div>
    )
}
