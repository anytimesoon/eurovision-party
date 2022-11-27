import { writable } from "svelte/store";
import type { CountryModel } from "$lib/models/classes/country.model";

// type State = {
//     countries: Array<CountryModel>;
//     error?: string;
// };
  
export const countryStore = writable<CountryModel[]>([]);

export const getCountries = async () => {
    let c = fetch("http://localhost:8080/api/country").
    then((res) => res.json() as Promise<Array<CountryModel>>).
    then(data => {
        countryStore.set(data);
    })
}