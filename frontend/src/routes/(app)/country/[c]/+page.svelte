<script lang="ts">
    import {VoteModel} from "$lib/models/classes/vote.model";
    import type {CountryModel} from "$lib/models/classes/country.model";
    import type {ActionData, PageData} from './$types';
    import VoteForm from "$lib/components/forms/VoteForm.svelte";

    export let data:PageData
    export let form:ActionData
    let country:CountryModel = data.country
    let vote:VoteModel = data.vote

    const updateVote = (form) => {
        if(form != null) {
            vote = form.vote
        }
    }


    $: country = data.country
    $: vote = data.vote
    $: updateVote(form)
</script>


<h1>{country.flag} {country.name}</h1>

{#if country.songName || country.bandName}
    <h3>{country.songName}</h3>
    <h5>{country.bandName}</h5>
{/if}

<h4>Song</h4>
<VoteForm countrySlug={country.slug} bind:score={vote.song} cat="song"/>

<h4>Performance</h4>
<VoteForm countrySlug={country.slug} bind:score={vote.performance} cat="performance"/>

<h4>Costume</h4>
<VoteForm countrySlug={country.slug} bind:score={vote.costume} cat="costume"/>

<h4>Staging and Props</h4>
<VoteForm countrySlug={country.slug} bind:score={vote.props} cat="props"/>

