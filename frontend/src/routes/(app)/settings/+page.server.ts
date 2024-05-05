import type {Actions} from "./$types";
import {userGoEP} from "$lib/models/enums/endpoints.enum";
import type {ResponseModel} from "$lib/models/classes/response.model";
import type {UserModel} from "$lib/models/classes/user.model";

export const actions : Actions = {
    updateName: async ({fetch, request}) => {
        const fd = await request.formData()
        // @ts-ignore
        const user:UserModel = Object.fromEntries([...fd]) as UserModel

        const res = await fetch(userGoEP.UPDATE, {
            method: "PUT",
            body: JSON.stringify(user)
        })

        const userResp:ResponseModel<UserModel> = await res.json()

        return {
            hideNameForm: true,
            error: userResp.error,
            user: userResp.body
        }
    },
    updateImg: async ({fetch, request}) => {
        const fd = await request.formData()
        let user:UserModel
        let error:string

        const res = await fetch(userGoEP.UPDATE_IMAGE, {
            method: "PUT",
            body: fd
        })

        if (!res.ok) {
            error = "Could not process the image. Please try another."
        } else {
            const userResp = await res.json()
            user = userResp.body
        }

        return {
            hideNameForm: true,
            avatarUpdated: true,
            error,
            user: user
        }
    }
}