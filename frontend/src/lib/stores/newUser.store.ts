import type {NewUserModel} from "$lib/models/classes/newUser.model";
import {writable} from "svelte/store";

export const newUserStore = writable<NewUserModel[]>([])