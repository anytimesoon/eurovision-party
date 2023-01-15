import {writable} from "svelte/store";
import type {UserModel} from "$lib/models/classes/user.model";

const users = new Map<string, UserModel>()

export const userStore = writable<Map<string, UserModel>>(users);

export function updateUserStore(users : UserModel[]) {
    for (let i = 0; i<users.length; i++){
        const u = users[i]
        userStore.update((users:Map<string, UserModel>) => {
            users.set(u.id, u)
            return users
        })
    }
}