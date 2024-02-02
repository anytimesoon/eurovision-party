import {type Writable, writable} from "svelte/store";

export const socketStateStore:Writable<boolean> =newSocketStateStore()

function newSocketStateStore(): Writable<boolean> {
    const {subscribe, set} = writable(false)
    return {
        subscribe,
        isReady: (state:boolean) => set(state)
    }
}