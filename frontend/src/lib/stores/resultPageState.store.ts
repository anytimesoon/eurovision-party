import {writable} from "svelte/store";
import {ResultPageStateModel} from "$lib/models/classes/resultPageState.model";
import {voteCats} from "$lib/models/enums/categories.enum";

export const resultPageState = newResultPageState()

let currentState: ResultPageStateModel
resultPageState.subscribe(val => currentState = val)

function newResultPageState() {
    const {subscribe, update, set} = writable(new ResultPageStateModel())

    return {
        subscribe,
        update,
        set,
        reset: () => reset(),
        isDefault: (): boolean => isDefault(),
        hasUserSelected: (): boolean => hasUserSelected()
    }
}

function reset() {
    currentState.reset()
    resultPageState.set(currentState)
}

function isDefault(): boolean {
    return currentState.userId === "" && currentState.category === voteCats.TOTAL && !currentState.sortByDescending
}

function hasUserSelected(): boolean {
    return currentState.userId !== ""
}