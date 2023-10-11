import type {RequestHandler} from "@sveltejs/kit";
import type {CountryModel} from "$lib/models/classes/country.model";
import {countryGoEP} from "$lib/models/enums/endpoints.enum";
import {json} from "@sveltejs/kit";
import {building} from "$app/environment";

export const GET :RequestHandler = async ({fetch, cookies}): Promise<Response> => {
    if (!building) {
        const countryRes: Response = await fetch(countryGoEP.ALL)

        let countries: Array<CountryModel> = await countryRes.json()
        return json(countries)
    }
}