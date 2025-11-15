import type {NewUserModel} from "$lib/models/classes/newUser.model";
import {writable} from "svelte/store";

export const newUserStore = newNewUserStore()

function newNewUserStore() {
    const {subscribe, update, set} = writable(new Array<NewUserModel>())

    return {
        subscribe,
        update,
        set,
        sortUsers: () => sortNewUsers(),
        addAndSort: (user: NewUserModel) => addAndSortNewUsers(user),
    }
}

function sortNewUsers() {
    newUserStore.update(users => {
        return sortingAlgo(users)
    })
}

function addAndSortNewUsers(user: NewUserModel) {
    newUserStore.update(users => {
        users.push(user)
        return sortingAlgo(users)
    })
}

function sortingAlgo(users: NewUserModel[]) {
    return users.sort((a, b) => (a.name > b.name) ? 1 : ((b.name > a.name) ? -1 : 0))
}