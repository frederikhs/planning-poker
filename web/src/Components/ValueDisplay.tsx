import {useMemo} from "react";

export default function ValueDisplay(props: { values: number[], valuesVisible: boolean, toggleVisibilityFn: () => void }) {
    const averageAnswer = useMemo(() => {
        const filtered = props.values.filter((value) => value !== -1)
        if (filtered.length === 0) {
            return -1
        }

        const sum = filtered.reduce((value, start) => value + start, 0)

        return Math.round(sum / filtered.length)
    }, [props.values])

    const noAnswersOrValuesNotVisible = (averageAnswer === -1 || !props.valuesVisible)

    return (
        <div className={`flex justify-center py-4`}>
            <div className={"hover:cursor-pointer"} onClick={() => props.toggleVisibilityFn()}>
                <p className={"text-8xl select-none text-green-600 rounded-md px-8 py-2 " + (noAnswersOrValuesNotVisible ? "" : "shadow-lg")}>
                    {averageAnswer === -1 && props.valuesVisible && <span>x</span>}
                    {!props.valuesVisible && <span>🃏</span>}
                    {averageAnswer !== -1 && props.valuesVisible && averageAnswer}
                </p>
            </div>
        </div>
    )
}
