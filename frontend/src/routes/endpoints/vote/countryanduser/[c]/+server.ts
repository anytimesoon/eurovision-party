import type {RequestHandler} from "@sveltejs/kit";
import {json} from "@sveltejs/kit";
import type {VoteModel} from "$lib/models/classes/vote.model";
import {voteGoEP} from "$lib/models/enums/endpoints.enum";

export const GET :RequestHandler = async ({fetch, params}): Promise<Response> => {
    const voteRes = await fetch(voteGoEP.BY_COUNTRY_AND_USER + params.c);
    let vote:VoteModel = await voteRes.json()
    return json(vote)
}