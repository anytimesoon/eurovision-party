import {writable} from "svelte/store";
import {browser} from "$app/environment";

export const loginURI = writable<string>(browser && localStorage.getItem("loginURI") || "")

loginURI.subscribe((val) => {
    browser && localStorage.setItem("loginURI", val)
})