import {writable} from "svelte/store";
import type {CountryModel} from "$lib/models/classes/country.model";
import {browser} from "$app/environment";

const defaultCountryList:CountryModel[] = new Array<CountryModel>();
// export const partCountryStore = writable<CountryModel[]>(browser && JSON.parse(localStorage.getItem("partCountryStore") || JSON.stringify(defaultCountryList)));
//
// partCountryStore.subscribe((val) => {
//     browser && localStorage.setItem("partCountryStore", JSON.stringify(val))
// });

export const partCountryStore = writable<CountryModel[]>(defaultCountryList)