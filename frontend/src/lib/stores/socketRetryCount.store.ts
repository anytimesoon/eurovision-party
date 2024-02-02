import {type Writable, writable} from "svelte/store";

export const socketRetryCount = newSocketRetryCountStore()

function newSocketRetryCountStore(): Writable<number> {
    const {subscribe, update, set} = writable(0)
    return {
        subscribe,
        increment: () => update((n) => n + 1),
        reset: () => set(0)
    }
}