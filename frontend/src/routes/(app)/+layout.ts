import {countryGoEP, userGoEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import {CountryModel} from "$lib/models/classes/country.model";
import {redirect} from "@sveltejs/kit";
import type {UserModel} from "$lib/models/classes/user.model";
import type {LayoutServerLoad} from "./$types";
import {browser} from "$app/environment";
export const ssr = false;

export const load:LayoutServerLoad =  ( async ({ fetch }) => {
    console.log(browser && localStorage.getItem("session"))
    const countryRes = await fetch(countryGoEP.ALL, {
        headers: {
            "Authorization": browser && localStorage.getItem("session")
        }
    })
    const countries: ResponseModel<CountryModel[]> = await countryRes.json()

    if (countries.error != "") {
        throw redirect(303, "/login")
    }

    const countryModels = countries.body.map((country):CountryModel => {
        return new CountryModel().deserialize(country)
    })

    const usersRes = await fetch(userGoEP.ALL, {
        headers: {
            "Authorization": browser && localStorage.getItem("session")
        }
    })
    const users: ResponseModel<Map<string, UserModel>> = await usersRes.json()

    return {
        countries: countryModels,
        users: users.body
    }
});
