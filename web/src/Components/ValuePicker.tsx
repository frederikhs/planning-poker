import {XMarkIcon} from "@heroicons/react/24/solid";

export default function ValuePicker(props: { pickFn: (v: number) => void, values: number[], pickedValue: number | null }) {
    return (
        <div className="fixed bottom-0 bg-gray-200 w-screen p-4 border-t-2 border-gray-600">
            <div className="grid sm:flex sm:justify-center grid-rows-2 grid-flow-col gap-4">
                <button onClick={() => props.pickFn(-1)} className={"value-button value-button-orange lowercase"}>
                    x
                </button>
                {props.values.map((value, index) => {
                    return <button key={index} onClick={() => props.pickFn(value)}
                                   className={`value-button ${value === props.pickedValue ? 'value-button-green' : 'value-button-gray'}`}>{value}</button>
                })}
            </div>

            {/*<div className="grid sm:flex sm:justify-center grid-rows-1 grid-flow-col gap-4 mt-4">*/}
            {/*    <button onClick={() => props.pickFn(-1)} className={"value-button value-button-orange"}>*/}
            {/*        <XMarkIcon className={"w-5"}/>*/}
            {/*    </button>*/}
            {/*</div>*/}
        </div>
    )
}
