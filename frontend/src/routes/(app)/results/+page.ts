import {userSvelteEP, voteSvelteEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import type {ResultModel} from "$lib/models/classes/result.model";
import type {UserModel} from "$lib/models/classes/user.model";

export async function load({fetch}) {
    const res = await fetch(voteSvelteEP.RESULTS)
    const results:ResponseModel<ResultModel[]> = await res.json()

    return {
        results: results.body
    }
}