import {writable} from "svelte/store";
import {UserModel} from "$lib/models/classes/user.model";
import {browser} from "$app/environment";

export const userStore = writable<Map<string, UserModel>>(new Map<string, UserModel>())

export const currentUser = writable<UserModel>(browser && JSON.parse(localStorage.getItem("currentUser") || JSON.stringify(new UserModel())))

currentUser.subscribe((val) => {
    browser && localStorage.setItem("currentUser", JSON.stringify(val))
})

export const botId = writable<string>(browser && localStorage.getItem("botUser") )

botId.subscribe((val) => {
    browser && localStorage.setItem("botUser", val)
})