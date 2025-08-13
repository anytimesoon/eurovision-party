import type {ResponseModel} from "$lib/models/classes/response.model";
import {sessionStore} from "$lib/stores/session.store";
import {goto} from "$app/navigation";
import {errorStore} from "$lib/stores/error.store";

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
    return await postOrPut<T, K>("POST", url, body) as Promise<T>
}

export async function postWithError<T, K>(url: string, body: K): Promise<ResponseModel<T>> {
    return await postOrPut<T, K>("POST", url, body, true) as Promise<ResponseModel<T>>
}

export async function put<T, K>(url: string, body: K): Promise<T> {
    return await postOrPut<T, K>("PUT", url, body) as Promise<T>
}

async function postOrPut<T, K>(method: string, url: string, body: K, withError: boolean = false): Promise<T | ResponseModel<T>> {
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
        if (withError) {
            return json
        } else {
            handleError(response.status, json.error)
        }
    }
}

function handleError(status: number, error: string) {
    errorStore.set(error)
    switch (status) {
        case 401:
            goto("/login")
            break;
        case 403:
            goto("/settings")
            break;
        default:
            console.error(error)
            goto("/error")
    }
}