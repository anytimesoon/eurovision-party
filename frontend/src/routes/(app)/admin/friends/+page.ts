import {userSvelteEP} from "$lib/models/enums/endpoints.enum";
import type {ResponseModel} from "$lib/models/classes/response.model";
import {NewUserModel} from "$lib/models/classes/user.model";

export async function load({fetch}) {
    const userRes = await fetch(userSvelteEP.REGISTERED)
    const users:ResponseModel<NewUserModel[]> = await userRes.json()

    return {
        users: users.body
    }
}