import {type Writable, writable} from "svelte/store";

export const socketStateStore =newSocketStateStore()

function newSocketStateStore() {
    const {subscribe, set} = writable(false)
    return {
        subscribe,
        set,
        isReady: (state:boolean) => set(state)
    }
}