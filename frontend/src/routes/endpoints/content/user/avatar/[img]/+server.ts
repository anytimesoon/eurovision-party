import type {RequestHandler} from "@sveltejs/kit";
import {staticGoEP} from "$lib/models/enums/endpoints.enum";

export const GET :RequestHandler = async ({fetch, params}): Promise<Response> => {
    return await fetch(staticGoEP.IMG+params.img)
}