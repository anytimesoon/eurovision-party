import { writable } from "svelte/store";
import type { TokenModel } from "$lib/models/classes/token.model";
  
export const tokenStore = writable<TokenModel>();

export const setToken = (token:TokenModel) => {
    tokenStore.set(token)
}