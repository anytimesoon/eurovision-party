import {chatEP, countrySvelteEP, userSvelteEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import {CountryModel} from "$lib/models/classes/country.model";
import {redirect} from "@sveltejs/kit";
import type {UserModel} from "$lib/models/classes/user.model";
import type {LayoutServerLoad} from "../../../.svelte-kit/types/src/routes/(login)/$types";
export const ssr = false;

export const load:LayoutServerLoad =  ( async ({ fetch }) => {
    const countryRes = await fetch(countrySvelteEP.ALL)
    const countries: ResponseModel<CountryModel[]> = await countryRes.json()

    if (countries.error != "") {
        throw redirect(303, "/login")
    }

    const countryModels = countries.body.map((country):CountryModel => {
        return new CountryModel().deserialize(country)
    })

    const usersRes = await fetch(userSvelteEP.ALL)
    const users: ResponseModel<Map<string, UserModel>> = await usersRes.json()

    let socket = connectToSocket()

    return {
        countries: countryModels,
        socket: socket,
        users: users.body
    }
});

function connectToSocket() {
    let socket = new WebSocket(chatEP)

    socket.onopen = function () {
        console.log("You're connected. Welcome to the party!!!🎉")
    };

    socket.onclose = function (e) {
        console.log("Connection stopped. Attempting to reconnect")
        setTimeout(connectToSocket, Math.min(10000, 250))
    };

    return socket
}