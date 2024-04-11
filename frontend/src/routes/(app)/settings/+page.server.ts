import type {Actions} from "./$types";
import {userGoEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import {UserModel} from "$lib/models/classes/user.model";

export const actions : Actions = {
    updateName: async ({fetch, request}) => {
        const fd = await request.formData()
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
        console.log(fd)
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

        console.log(user)
        console.log(error)

        return {
            hideNameForm: true,
            avatarUpdated: true,
            error,
            user: user
        }
    }
}