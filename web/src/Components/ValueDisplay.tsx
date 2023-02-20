import {useMemo} from "react";

export default function ValueDisplay(props: { values: number[], answerValues: number[], valuesVisible: boolean, toggleVisibilityFn: () => void }) {
    const averageAnswer = useMemo(() => {
        const filtered = props.answerValues.filter((value) => value !== -1)
        if (filtered.length === 0) {
            return -1
        }

        const sum = filtered.reduce((value, start) => value + start, 0)

        return sum / filtered.length
    }, [props.answerValues])

    const closestFibNumber = useMemo(() => {
        return props.values.reduce(function (prev, curr) {
            return (Math.abs(curr - averageAnswer) < Math.abs(prev - averageAnswer) ? curr : prev);
        })
    }, [averageAnswer])

    const noAnswersOrValuesNotVisible = (averageAnswer === -1 || !props.valuesVisible)

    return (
        <div className={`flex justify-center py-4`}>
            <div className={"hover:cursor-pointer"} onClick={() => props.toggleVisibilityFn()}>
                <p className={"text-8xl select-none text-green-600 rounded-md px-8 py-2 " + (noAnswersOrValuesNotVisible ? "" : "shadow-lg")}>
                    {averageAnswer === -1 && props.valuesVisible && <span>x</span>}
                    {!props.valuesVisible && <span>🃏</span>}
                    {averageAnswer !== -1 && props.valuesVisible && closestFibNumber}
                </p>
            </div>
        </div>
    )
}
