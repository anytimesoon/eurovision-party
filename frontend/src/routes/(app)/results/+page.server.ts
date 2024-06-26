import type {Actions} from "./$types";
import {voteGoEP} from "$lib/models/enums/endpoints.enum";
import type {ResponseModel} from "$lib/models/classes/response.model";
import type {ResultModel} from "$lib/models/classes/result.model";

export const actions : Actions = {
    getUserResults: async ({fetch, request}) => {
        const fd = await request.formData()
        const user = Object.fromEntries(fd)

        let userPromise:Response
        if(user.id === ""){
            userPromise = await fetch(voteGoEP.RESULTS)
        } else {
            userPromise = await fetch(voteGoEP.RESULTS + user.id)
        }

        const userResultJSON:ResponseModel<ResultModel[]> = await userPromise.json()
        const userId = user.id.toString()
        return {
            success: true,
            results: userResultJSON.body,
            selection: userId
        }
    }
}