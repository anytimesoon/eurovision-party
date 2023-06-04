import {authEP, countrySvelteEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import {CountryModel} from "$lib/models/classes/country.model";
import {redirect} from "@sveltejs/kit";
export const ssr = false;

export const load =  ( async ({ fetch }) => {
    const countryRes = await fetch(countrySvelteEP.ALL)
    const countries: ResponseModel<CountryModel[]> = await countryRes.json()

    if (countries.error != "") {
        throw redirect(303, "/login")
    }

    return {
        // ...data,
        countries: countries.body
    }
}) satisfies LayoutServerLoad;