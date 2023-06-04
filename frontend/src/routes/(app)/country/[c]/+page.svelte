<script lang="ts">
    import {VoteModel} from "$lib/models/classes/vote.model";
    import type {CountryModel} from "$lib/models/classes/country.model";
    import type { PageData } from './$types';

    export let data:PageData;
    let country:CountryModel = data.country
    let vote:VoteModel = data.vote

    const options:Array<{ label: string; value: number }>  = [{
        value: 1,
        label: '1'
    },{
        value: 2,
        label: '2'
    },{
        value: 3,
        label: '3'
    },{
        value: 4,
        label: '4'
    },{
        value: 5,
        label: '5'
    },{
        value: 6,
        label: '6'
    },{
        value: 7,
        label: '7'
    },{
        value: 8,
        label: '8'
    },{
        value: 9,
        label: '9'
    },{
        value: 10,
        label: '10'
    }]

    $: country = data.country
    $: vote = data.vote
</script>

    <h1>{country.flag} {country.name}</h1>

    <p>{country.songName} by {country.bandName}</p>




