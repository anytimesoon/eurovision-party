import {voteSvelteEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import type {ResultModel} from "$lib/models/classes/result.model";
import {browser} from "$app/environment";

export async function load({fetch}) {
    const userId = browser && JSON.parse(localStorage.getItem("resultPageState")).userId
    let userPromise:Response
    if(userId === null || userId === "total"){
        userPromise = await fetch(voteSvelteEP.RESULTS)
    } else {
        userPromise = await fetch(voteSvelteEP.RESULTS + userId)
    }

    const results:ResponseModel<ResultModel[]> = await userPromise.json()

    return {
        results: results.body
    }
}