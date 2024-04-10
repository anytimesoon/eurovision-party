import {LoginModel} from "$lib/models/classes/login.model";
import {authEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import type {PageLoad} from ".$/types";
import {redirect} from "@sveltejs/kit";
import type {SessionModel} from "$lib/models/classes/session.model";
import {browser} from "$app/environment";

export const load =  ( async ({ params, fetch }) => {

    const payload = new LoginModel(params.t, params.u)

    const res = await fetch(authEP.LOGIN, {
        method: "POST",
        body: JSON.stringify(payload)
    });

    const login: ResponseModel<SessionModel> = await res.json()
    const hasLoggedIn:boolean = browser && localStorage.getItem("visited") || false

    if (!hasLoggedIn) {
        browser && localStorage.setItem("visited", true)
    }

    if (login.error != "") {
        throw redirect(303, "/login")
    }

    return {
        currentUser: login.body.user,
        botId: login.body.botId,
        hasLoggedIn,
        sessionToken: login.body.token,
        loginToken: params.t
    }
}) satisfies PageLoad;