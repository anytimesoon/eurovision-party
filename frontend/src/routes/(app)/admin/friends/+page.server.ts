import type {Actions} from "./$types";
import {authEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import type {NewUserModel} from "$lib/models/classes/user.model";
import {UserModel} from "$lib/models/classes/user.model";

export const prerender = false

export const actions : Actions = {
    register: async ({fetch, request}) => {
        const fd = await request.formData()
        const newUser:NewUserModel = Object.fromEntries([...fd]) as NewUserModel;
        const resProm = await fetch(authEP.REGISTER, {
            method: "POST",
            body: JSON.stringify(newUser)
        })

        const res:ResponseModel<UserModel> = await resProm.json()

        return {
            success: true,
            user: res.body,
            error: res.error
        }
    }
}