import {readable, writable} from "svelte/store";
import type { CountryModel } from "$lib/models/classes/country.model";
import {browser} from "$app/env";

const defaultCountryList:CountryModel[] = [];
export const countryStore = writable<CountryModel[]>([]);

export const partCountryStore = readable<CountryModel[]>(browser && JSON.parse(localStorage.getItem("partCountryStore") || JSON.stringify(defaultCountryList)));

partCountryStore.subscribe((val) => {
    browser && localStorage.setItem("partCountryStore", JSON.stringify(val))
});