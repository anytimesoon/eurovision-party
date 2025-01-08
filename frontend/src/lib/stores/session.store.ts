import {writable} from "svelte/store";
import {browser} from "$app/environment";

export const sessionStore = writable<string>(
    browser && (localStorage.getItem("session") || "")
)

sessionStore.subscribe((val) => {
    browser && localStorage.setItem("session", val)
})