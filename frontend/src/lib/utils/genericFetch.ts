import type {ResponseModel} from "$lib/models/classes/response.model";
import {redirect} from "@sveltejs/kit";
import {sessionStore} from "$lib/stores/session.store";

let session:string
sessionStore.subscribe(val => session = val)

export async function get(url: string): Promise<any> {
    const response = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
            "Authorization": session
        }
    })
    const json: ResponseModel<any> = await response.json()

    if (response.ok) {
        return json.body
    } else {
        handleError(response.status, json.error)
    }
}

export async function post<T, K>(url: string, body: K): Promise<T> {
    return await postOrPut<T, K>("POST", url, body)
}

export async function put<T, K>(url: string, body: K): Promise<T> {
    return await postOrPut<T, K>("PUT", url, body)
}

async function postOrPut<T, K>(method: string, url: string, body: K): Promise<T> {
    const response = await fetch(url, {
        method: method,
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
            "Authorization": session
        },
        body: JSON.stringify(body)
    })
    const json: ResponseModel<T> = await response.json()
    if (response.ok) {
        return json.body
    } else {
        handleError(response.status, json.error)
    }
}

function handleError(status: number, error: string) {
    switch (status) {
        case 401:
        case 403:
            redirect(status, "/login")
            break;
        default:
            console.error(error)
            redirect(status, "/error")
    }
}