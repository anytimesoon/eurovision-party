import type {CountryModel} from "$lib/models/classes/country.model";
import {countrySvelteEP, voteSvelteEP} from "$lib/models/enums/endpoints.enum";
import type {VoteModel} from "$lib/models/classes/vote.model";
import type {ResponseModel} from "$lib/models/classes/response.model";

export async function load({fetch, params}) {
    const countryRes = await fetch(countrySvelteEP.FIND_ONE + params.c);

    const country:ResponseModel<CountryModel> = await countryRes.json()

    const voteRes = await fetch(voteSvelteEP.BY_COUNTRY_AND_USER + params.c)

    const vote:ResponseModel<VoteModel> = await voteRes.json()

    return {
        country: country.body,
        vote: vote.body
    }
}