import type {CountryModel} from "$lib/models/classes/country.model";
import {sendCreateOrUpdate} from "$lib/helpers/sender.helper";
import {countryEP} from "$lib/models/enums/endpoints.enum";
import {countryStore} from "$lib/stores/country.store";

let countries:CountryModel[];
countryStore.subscribe( val => {
    countries = val;
})

export function updateCountryParticipation(country:CountryModel) {
    country.participating = !country.participating;

    send(country);
}

export function updateCountrySong(country:CountryModel, song:string) {
    country.songName = song;

    send(country);
}

export function updateCountryBand(country:CountryModel, band:string) {
    country.bandName = band;

    send(country);
}

function send(country:CountryModel) {
    sendCreateOrUpdate<CountryModel, CountryModel>(countryEP.UPDATE, country).then( data => country = data.body);

    const i:number = countries.findIndex((c:CountryModel) => c.id === country.id);
    countries[i] = country
    countryStore.set( countries );
}