import {writable} from "svelte/store";
import {UserModel} from "$lib/models/classes/user.model";
import {browser} from "$app/environment";
import type {NewUserModel} from "$lib/models/classes/user.model";

const defaultUsers = new Map<string, UserModel>()

export const userStore = writable<Map<string, UserModel>>(browser && JSON.parse(localStorage.getItem("userStore") || JSON.stringify(defaultUsers)));

userStore.subscribe((val) => {
    browser && localStorage.setItem("userStore", JSON.stringify(val))
});

export const registeredUserStore = writable<NewUserModel[]>([])

export const currentUserStore = writable<UserModel>(new UserModel())