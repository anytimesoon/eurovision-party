import { RequestModel } from "$lib/models/classes/request.model";
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

export async function sendPost<T, U>(endpoint : string, payload : T): Promise<ResponseModel<U>> {
    let resp = new ResponseModel<U>();
    let t = new TokenModel();
    tokenStore.subscribe(
        val => t = val
    )

    const request = new RequestModel<T>().build(t.token, payload);

    await fetch(endpoint, {
        method: "POST",
        body: JSON.stringify(request)
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