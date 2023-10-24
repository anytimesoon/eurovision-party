import {LoginModel} from "$lib/models/classes/login.model";
import {authEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import type {PageLoad} from ".$/types";
import {json, redirect, RequestHandler} from "@sveltejs/kit";
import type {SessionModel} from "$lib/models/classes/session.model";

export const prerender = false

export const GET :RequestHandler = async ({fetch, cookies, params}): Promise<Response> => {
    const payload = new LoginModel(params.token, params.userId)

    const res = await fetch(authEP.LOGIN, {
        method: "POST",
        body: JSON.stringify(payload),
        mode: 'cors'
    });

    const login: ResponseModel<SessionModel> = await res.json()
    cookies.set("session", login.body.token, login.body.opts)
    const hasLoggedIn:boolean = cookies.get("visited") || false

    if (!hasLoggedIn && login.error != "") {
        cookies.set('visited', 'true', { path: '/' })
    }

    return json(login)
}