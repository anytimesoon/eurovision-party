<script lang="ts">
    import {VoteModel} from "$lib/models/classes/vote.model";
    import type {CountryModel} from "$lib/models/classes/country.model";
    import type {ActionData, PageData} from './$types';
    import {currentUser} from "$lib/stores/user.store";
    import {voteOptions} from "$lib/models/classes/voteOptions.model";
    import { enhance } from '$app/forms';

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
<form method="POST" action="?/vote" use:enhance>
    <input name="countrySlug" type="hidden" value={country.slug}>
    <input name="cat" type="hidden" value="song">
    <input name="userId" type="hidden" value={$currentUser.id}>
    {#each voteOptions as { key, label }}
        <input type="radio" bind:group={vote.song} value={key} name="score" on:click={(e) => {
            e.target.parentElement.requestSubmit()
        }}/>
        <label>{label}</label>
    {/each}
</form>

<h4>Performance</h4>
<form method="POST" action="?/vote" use:enhance>
    <input name="countrySlug" type="hidden" value={country.slug}>
    <input name="cat" type="hidden" value="performance">
    <input name="userId" type="hidden" value={$currentUser.id}>
    {#each voteOptions as { key, label }}
        <input type="radio" bind:group={vote.performance} value={key} name="score" on:click={(e) => {
            e.target.parentElement.requestSubmit()
        }}/>
        <label>{label}</label>
    {/each}
</form>

<h4>Costume</h4>
<form method="POST" action="?/vote" use:enhance>
    <input name="countrySlug" type="hidden" value={country.slug}>
    <input name="cat" type="hidden" value="costume">
    <input name="userId" type="hidden" value={$currentUser.id}>
    {#each voteOptions as { key, label }}
        <input type="radio" bind:group={vote.costume} value={key} name="score" on:click={(e) => {
            e.target.parentElement.requestSubmit()
        }}/>
        <label>{label}</label>
    {/each}
</form>

<h4>Staging and Props</h4>
<form method="POST" action="?/vote" use:enhance>
    <input name="countrySlug" type="hidden" value={country.slug}>
    <input name="cat" type="hidden" value="props">
    <input name="userId" type="hidden" value={$currentUser.id}>
    {#each voteOptions as { key, label }}
        <input type="radio" bind:group={vote.props} value={key} name="score" on:click={(e) => {
            e.target.parentElement.requestSubmit()
        }}/>
        <label>{label}</label>
    {/each}
</form>
