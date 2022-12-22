import { RequestModel } from "$lib/models/classes/request.model";
import { ResponseModel } from "$lib/models/classes/response.model"
import { TokenModel } from "$lib/models/classes/token.model";
import { tokenStore } from '$lib/stores/token.store';

export function sendGet<T, U>(endpoint : string, payload : T) {
    fetch(endpoint, {
        method: "GET",
        body: JSON.stringify(payload),
    }).
    then(
        res => {
            res.json() as Promise<U>
            return res
        }
    ).
    catch((e) => {
        console.log(e)
    })
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
        body: JSON.stringify(request),
    }).
    then(
        res => res.json() as Promise<ResponseModel<U>>
    ).
    then(
        json =>resp = json
    ).
    catch((e) => {
        console.log(e)
    });

    return resp;
}