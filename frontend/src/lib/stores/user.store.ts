import {writable} from "svelte/store";
import type {UserModel} from "$lib/models/classes/user.model";
import {browser} from "$app/env";

const defaultUsers = new Map<string, UserModel>()

export const userStore = writable<Map<string, UserModel>>(browser && JSON.parse(localStorage.getItem("userStore") || JSON.stringify(defaultUsers)));

userStore.subscribe((val) => {
    browser && localStorage.setItem("userStore", JSON.stringify(val))
});
