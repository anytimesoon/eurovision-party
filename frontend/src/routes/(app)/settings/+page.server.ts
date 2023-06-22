import type {Actions} from "./$types";
import {userGoEP, voteGoEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import {ResultModel} from "$lib/models/classes/result.model";
import type {UserModel} from "$lib/models/classes/user.model";

export const actions : Actions = {
    showNameForm: async ({fetch, request}) => {

        return {
            hideNameForm: false,
            hideImgForm: true
        }
    },
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
            hideImgForm: true,
            error: userResp.error,
            user: user
        }
    },
    showImgForm: async ({fetch, request}) => {

        return {
            hideNameForm: true,
            hideImgForm: false
        }
    },
    updateImg: async ({fetch, request}) => {
        const fd = await request.formData()
        const user:UserModel = Object.fromEntries([...fd]) as UserModel

        const res = await fetch(userGoEP.UPDATE, {
            method: "PUT",
            body: JSON.stringify(user)
        })

        const userResp:ResponseModel<UserModel> = await res.json()

        return {
            hideNameForm: true,
            hideImgForm: true,
            error: userResp.error,
            user: user
        }
    }
}