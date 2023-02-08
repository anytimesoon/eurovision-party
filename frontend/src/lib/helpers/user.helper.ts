import {browser} from "$app/env";
import type {IUser} from "$lib/models/interfaces/iuser.interface";
import {UserModel} from "$lib/models/classes/user.model";

export function currentUser(): IUser | null {
    if (browser) {
        const dummyUser:UserModel = new UserModel()
        return JSON.parse(localStorage.getItem("me") || JSON.stringify(dummyUser))
    } else {
        return null
    }
}