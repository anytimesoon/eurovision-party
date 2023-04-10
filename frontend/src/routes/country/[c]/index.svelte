<script type="ts">
    import {partCountryStore} from "$lib/stores/partCountry.store";
    import Radio from "$lib/components/RadioButton.svelte";
    import { page } from "$app/stores";
    import {CountryModel} from "$lib/models/classes/country.model";

    const country = $partCountryStore.find( (c:CountryModel) => {
        if (c.slug === $page.params.c) {
            return c
        }
    });



    const options = [{
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
</script>

{#await country}
<h1>{country.flag} {country.name}</h1>

<h3>Song</h3>
<p>{country.songName} by {country.bandName}</p>

<h3>Vote!</h3>
<Radio {options} fontSize={16} legend='Score the quality of the song'/>
<Radio {options} fontSize={16} legend='Score the quality of the performance'/>
<Radio {options} fontSize={16} legend='Score the quality of the costumes'/>
<Radio {options} fontSize={16} legend='Score the quality of the props'/>
    {/await}