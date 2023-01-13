import {writable} from "svelte/store";
import type {UserModel} from "$lib/models/classes/user.model";

export const userStore = writable<UserModel[]>([]);