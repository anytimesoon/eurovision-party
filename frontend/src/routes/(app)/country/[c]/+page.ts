import {CountryModel} from "$lib/models/classes/country.model";
import {VoteModel} from "$lib/models/classes/vote.model";
import {countryEP, voteEP} from "$lib/models/enums/endpoints.enum";
import {get} from "$lib/utils/genericFetch";

export async function load({params}) {
    const country: CountryModel = await get(countryEP.FIND_ONE + params.c)
        .then(result => CountryModel.deserialize(result));

    const vote: VoteModel = await get(voteEP.BY_COUNTRY_AND_USER + params.c)
        .then(result => VoteModel.deserialize(result))

    return {
        country: country,
        vote: vote
    }
}