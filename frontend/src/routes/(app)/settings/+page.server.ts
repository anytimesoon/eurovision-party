import type {Actions} from "./$types";
import {voteGoEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import {ResultModel} from "$lib/models/classes/result.model";

export const actions : Actions = {
    showNameForm: async ({fetch, request}) => {
        // const fd = await request.formData()

        return {
            hideNameForm: false
        }
    },
    updateName: async ({fetch, request}) => {
        const fd = await request.formData()
        const user = Object.fromEntries(fd)

        console.log(user)

        return {
            hideNameForm: true
        }
    }
}