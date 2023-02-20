export default function ValuePicker(props: { pickFn: (v: number) => void, values: number[], pickedValue: number | null }) {
    return (
        // <div className="fixed bg-white shadow-xl bottom-0 w-screen p-4 border-gray-600">
        <div className="bg-white w-screen p-4 border-gray-600 border-t-4 border-gray-600">
            <div className="grid sm:flex sm:justify-center grid-cols-3 grid-flow-dense gap-4">
                <button onClick={() => props.pickFn(-1)} className={"value-button value-button-gray lowercase"}>
                    x
                </button>
                {props.values.map((value, index) => {
                    return <button key={index} onClick={() => props.pickFn(value)}
                                   className={`shadow-lg value-button ${value === props.pickedValue ? 'value-button-green' : 'value-button-gray'}`}>{value}</button>
                })}
            </div>
        </div>
    )
}
