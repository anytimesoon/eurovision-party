import type {CountryModel} from "$lib/models/classes/country.model";
import {countryGoEP, voteGoEP} from "$lib/models/enums/endpoints.enum";
import {VoteModel} from "$lib/models/classes/vote.model";
import type {ResponseModel} from "$lib/models/classes/response.model";
import {browser} from "$app/environment";

export async function load({fetch, params}) {
    const countryRes = await fetch(countryGoEP.FIND_ONE + params.c, {
        headers: {
            "Authorization": browser && localStorage.getItem("session")
        }
    });

    const country:ResponseModel<CountryModel> = await countryRes.json()

    const voteRes = await fetch(voteGoEP.BY_COUNTRY_AND_USER + params.c, {
            headers: {
                "Authorization": browser && localStorage.getItem("session")
            }
        })

    const res:ResponseModel<VoteModel> = await voteRes.json()
    const vote = new VoteModel().deserialize(res.body)

    return {
        country: country.body,
        vote: vote
    }
}