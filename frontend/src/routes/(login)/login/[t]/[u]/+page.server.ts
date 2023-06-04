import {LoginModel} from "$lib/models/classes/login.model";
import {authEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import type {PageLoad} from ".$/types";
import {redirect} from "@sveltejs/kit";
import type {SessionModel} from "$lib/models/classes/session.model";

export const load =  ( async ({ params, fetch, cookies }) => {

    const payload = new LoginModel(params.t, params.u)

    const res = await fetch(authEP.LOGIN, {
        method: "POST",
        body: JSON.stringify(payload),
        mode: 'cors'
    });

    const login: ResponseModel<SessionModel> = await res.json()

    cookies.set("session", login.body.token, login.body.opts)


    if (login.error != "") {
        throw redirect(303, "/login")
    }

    return {
        currentUser: login.body.user
    }
}) satisfies PageLoad;