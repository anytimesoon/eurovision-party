import { writable } from "svelte/store";
import type { CountryModel } from "$lib/models/classes/country.model";
  
export const countryStore = writable<CountryModel[]>([]);
export const activeCountryStore = writable<CountryModel[]>([]);

export const updateCountry = async (country:CountryModel) => {
    fetch('http://localhost:8080/api/country',{
        method: "PUT",
        mode: 'cors',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(this)
    }).then(response => response.json())
    .then(data => {
        if(data.participating){
            activeCountryStore.update(data)
        }
    })
    .catch((err) => {
       console.log(err)
    })
}

export const getCountries = async () => {
    let c = fetch("http://localhost:8080/api/country").
    then((res) => res.json() as Promise<Array<CountryModel>>).
    then(data => {
        countryStore.set(data);
    })
}

export const getActiveCountries = async () => {
    let c = fetch("http://localhost:8080/api/country/active").
    then((res) => res.json() as Promise<Array<CountryModel>>).
    then(data => {
        countryStore.set(data);
    })
}