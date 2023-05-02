import { ResponseModel } from "$lib/models/classes/response.model"
import { TokenModel } from "$lib/models/classes/token.model";
import { tokenStore } from '$lib/stores/token.store';

export async function sendGet<T>(endpoint : string): Promise<ResponseModel<T>> {
    let resp = new ResponseModel<T>();
    let t = new TokenModel();
    tokenStore.subscribe(
        val => t = val
    )

    await fetch(endpoint, {
        method: "GET",
        headers: {
            "Authorization": t.token
        }
    }).
    then(
        res => res.json() as Promise<ResponseModel<T>>
    ).
    then(
        json => {
            resp = json
            if (resp.token.token !== "") {
                tokenStore.set(resp.token)
            }
        }
    ).
    catch((e) => {
        console.log(e)
    })

    return resp
}

export async function sendCreateOrUpdate<T, U>(endpoint: string, payload: T, method = "PUT"): Promise<ResponseModel<U>> {
    let resp = new ResponseModel<U>();
    let t = new TokenModel();
    tokenStore.subscribe(
        val => t = val
    )

    await fetch(endpoint, {
        method: method,
        headers: {
            "Authorization": t.token
        },
        body: JSON.stringify(payload, replace)
    }).
    then(
        res => res.json() as Promise<ResponseModel<U>>
    ).
    then(
        json => {
            resp = json
            if (resp.token.token !== "") {
                tokenStore.set(resp.token)
            }
        }
    ).
    catch((e) => {
        console.log(e)
    });

    return resp;
}

function replace(key, value){
    return key === "score" ? +value : value;
}