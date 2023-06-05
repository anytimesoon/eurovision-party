import {chatEP, countrySvelteEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import {CountryModel} from "$lib/models/classes/country.model";
import {redirect} from "@sveltejs/kit";
import type {CommentModel} from "$lib/models/classes/comment.model";
export const ssr = false;

export const load =  ( async ({ fetch }) => {
    const countryRes = await fetch(countrySvelteEP.ALL)
    const countries: ResponseModel<CountryModel[]> = await countryRes.json()

    if (countries.error != "") {
        throw redirect(303, "/login")
    }

    const countryModels = countries.body.map((country):CountryModel => {
        return new CountryModel().deserialize(country)
    })

    const socket = new WebSocket(chatEP);

    let timeout = 250;

    socket.onopen = function () {
        console.log("You're connected. Welcome to the party!!!🎉");
        timeout = 250;
    };

    socket.onclose = function (e) {
        console.log('Socket is closed. Reconnect will be reattempted in ' + timeout + "milliseconds. " + e.reason);
        // setTimeout(connectToSocket, Math.min(10000, timeout += timeout));
    };

    return {
        countries: countryModels,
        socket: socket
    }
}) satisfies LayoutServerLoad;