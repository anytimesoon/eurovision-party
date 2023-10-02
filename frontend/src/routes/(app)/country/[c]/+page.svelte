<script lang="ts">
    import type {VoteModel} from "$lib/models/classes/vote.model";
    import type {CountryModel} from "$lib/models/classes/country.model";
    import type {ActionData, PageData} from './$types';
    import VoteForm from "$lib/components/forms/VoteForm.svelte";

    export let data:PageData
    export let form:ActionData
    let country:CountryModel = data.country
    let vote:VoteModel = data.vote


    $: country = data.country
    $: vote = data.vote
    $: if (form) {
        vote = form.vote
    }
</script>

<div class="h-full overflow-auto p-3">
    <h1 class="text-center">{country.flag} {country.name}</h1>

    <div class="text-center py-3">
        {#if country.songName || country.bandName}
            <h4>{country.songName} <span class="text-xs text-typography-grey">by</span> {country.bandName}</h4>
        {/if}
    </div>

    <VoteForm vote={vote} catName="Song" countrySlug={country.slug} />
    <VoteForm vote={vote} catName="Performance" countrySlug={country.slug} />
    <VoteForm vote={vote} catName="Costumes" countrySlug={country.slug} />
    <VoteForm vote={vote} catName="Staging and Props" countrySlug={country.slug} />


</div>