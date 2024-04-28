import {writable} from "svelte/store";
import {browser} from "$app/environment";
import {ResultPageStateModel} from "$lib/models/classes/resultPageState.model";

let defaultState: ResultPageStateModel = new ResultPageStateModel()

export const resultPageState = writable<ResultPageStateModel>(
    browser && JSON.parse(
        localStorage.getItem("resultPageState") ||
        JSON.stringify(defaultState)
    )
)

resultPageState.subscribe((val) => {
    browser && localStorage.setItem("resultPageState", JSON.stringify(val))
})