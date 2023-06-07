import {json, RequestHandler} from "@sveltejs/kit";
import {userGoEP} from "$lib/models/enums/endpoints.enum";
import type {UserModel} from "$lib/models/classes/user.model";

export const GET :RequestHandler = async ({fetch}): Promise<Response> => {
    const userRes:Response = await fetch(userGoEP.ALL)

    let users:Array<UserModel> = await userRes.json()
    return json(users)
}