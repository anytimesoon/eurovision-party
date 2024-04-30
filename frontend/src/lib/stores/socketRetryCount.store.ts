import {writable} from "svelte/store";

export const socketRetryCount = newSocketRetryCountStore()

function newSocketRetryCountStore() {
    const {subscribe, update, set} = writable(0)
    return {
        subscribe,
        set,
        update,
        increment: () => update((n) => n + 1),
        reset: () => set(0)
    }
}