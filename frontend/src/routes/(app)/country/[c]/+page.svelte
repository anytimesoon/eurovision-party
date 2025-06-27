<script lang="ts">
    import type {VoteModel} from "$lib/models/classes/vote.model";
    import type {PageData} from './$types';
    import VoteForm from "$lib/components/forms/VoteForm.svelte";

    interface Props {
        data: PageData;
    }

    let { data }: Props = $props();
    let country = $derived(data.country)
    let vote = $derived(data.vote)

    const updateVote = (newVote: VoteModel) => {
        vote = newVote
    }
</script>

<div class="h-full overflow-auto p-3">
    <h1 class="text-center">
        <span class="font-sans">
            {country.flag}
        </span>
        {country.name}
    </h1>

    <div class="text-center py-3">
        {#if country.songName || country.bandName}
            <h4>{country.songName} <span class="text-xs text-typography-grey">by</span> {country.bandName}</h4>
        {/if}
    </div>

    <VoteForm vote={vote} catName="Song" countrySlug={country.slug} updateVote={updateVote}/>
    <VoteForm vote={vote} catName="Performance" countrySlug={country.slug} updateVote={updateVote}/>
    <VoteForm vote={vote} catName="Costumes" countrySlug={country.slug} updateVote={updateVote}/>
    <VoteForm vote={vote} catName="Staging and Props" countrySlug={country.slug} updateVote={updateVote}/>


</div>