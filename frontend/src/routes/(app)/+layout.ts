import {countrySvelteEP, userSvelteEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import {CountryModel} from "$lib/models/classes/country.model";
import {redirect} from "@sveltejs/kit";
import type {UserModel} from "$lib/models/classes/user.model";
import type {LayoutServerLoad} from "./$types";
export const ssr = false;

export const load:LayoutServerLoad =  ( async ({ fetch }) => {
    const countryRes = await fetch(countrySvelteEP.ALL)
    const countries: ResponseModel<CountryModel[]> = await countryRes.json()

    if (countries.error != "") {
        redirect(303, "/login");
    }

    const countryModels = countries.body.map((country):CountryModel => {
        return new CountryModel().deserialize(country)
    })

    const usersRes = await fetch(userSvelteEP.ALL)
    const users: ResponseModel<Map<string, UserModel>> = await usersRes.json()

    return {
        countries: countryModels,
        users: users.body
    }
});
