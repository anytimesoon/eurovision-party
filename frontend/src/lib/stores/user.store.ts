import {writable} from "svelte/store";
import {UserModel} from "$lib/models/classes/user.model";

export const currentUser = writable<UserModel>(browser && JSON.parse(localStorage.getItem("currentUser") || JSON.stringify(new UserModel())))

currentUser.subscribe((val) => {
    browser && localStorage.setItem("currentUser", JSON.stringify(val))
});

export const userStore = writable<Map<string, UserModel>>({})
export const activeUserStore = writable<Map<string, UserModel>>({})