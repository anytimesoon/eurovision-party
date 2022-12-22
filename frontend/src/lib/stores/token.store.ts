import { writable } from "svelte/store";
import type { IToken } from "$lib/models/interfaces/itoken.interface";
import { browser } from "$app/env";
import { TokenModel } from "$lib/models/classes/token.model";

const defaultToken:TokenModel = new TokenModel

export const tokenStore = writable<IToken>(browser && JSON.parse(localStorage.getItem("tokenStore") || JSON.stringify(defaultToken)));

tokenStore.subscribe((val) => {
    browser && localStorage.setItem("tokenStore", JSON.stringify(val))
});