<script lang="ts">
    import type {VoteModel} from "$lib/models/classes/vote.model";
    import type {CountryModel} from "$lib/models/classes/country.model";
    import type {ActionData, PageData} from './$types';
    import {currentUser} from "$lib/stores/user.store";
    import {voteOptions} from "$lib/models/classes/voteOptions.model";
    import { enhance } from '$app/forms';

    export let data:PageData
    export let form:ActionData
    let country:CountryModel = data.country
    let vote:VoteModel = data.vote
    let songForm:HTMLFormElement
    let perfForm:HTMLFormElement
    let costForm:HTMLFormElement
    let propForm:HTMLFormElement

    const updateVote = (form) => {
        if(form != null) {
            vote = form.vote
        }
    }


    $: country = data.country
    $: vote = data.vote
    $: updateVote(form)
</script>

<div class="h-full overflow-auto p-3">
    <h1 class="text-center">{country.flag} {country.name}</h1>

    <div class="text-center py-3">
        {#if country.songName || country.bandName}
            <h3>{country.songName}</h3>
            <p>performed by</p>
            <h5>{country.bandName}</h5>
        {/if}
    </div>

    <div class="py-5">
        <h4 class="text-center"><i class="fa-solid fa-music"></i> Song {vote.song}/10</h4>
        <form method="POST" action="?/vote" use:enhance bind:this={songForm}>
            <input name="countrySlug" type="hidden" value={country.slug}>
            <input name="cat" type="hidden" value="song">
            <input name="userId" type="hidden" value={$currentUser.id}>
            <div class="flex flex-row-reverse items-center mx-auto justify-between">
                {#each voteOptions.reverse() as { key, label }}
                    <input id="song-star-{label}" class="hidden peer" type="radio" bind:group={vote.song} value={key} name="score" on:click={() => {
                        songForm.requestSubmit()
                    }}/>
                    <label
                            class="text-yellow-400 ease-in-out duration-300 text-4xl
                            hover:text-yellow-500
                            before:content-['\2606']
                            peer-checked:before:content-['\2605']"
                            for="song-star-{label}"></label>
                {/each}
            </div>
        </form>
    </div>

    <div class="py-5">
        <h4 class="text-center"><i class="fa-solid fa-masks-theater"></i> Performance {vote.performance}/10</h4>
        <form method="POST" action="?/vote" use:enhance bind:this={perfForm}>
            <input name="countrySlug" type="hidden" value={country.slug}>
            <input name="cat" type="hidden" value="performance">
            <input name="userId" type="hidden" value={$currentUser.id}>
            <div class="flex flex-row-reverse items-center mx-auto justify-between">
                {#each voteOptions as { key, label }}
                    <input id="performance-star-{label}" class="hidden peer" type="radio" bind:group={vote.performance} value={key} name="score" on:click={() => {
                        perfForm.requestSubmit()
                    }}/>
                    <label
                            class="text-yellow-400 ease-in-out duration-300 text-4xl
                            hover:text-yellow-500
                            before:content-['\2606']
                            peer-checked:before:content-['\2605']"
                            for="performance-star-{label}"></label>
                {/each}
            </div>
        </form>
    </div>

    <div class="py-5">
        <h4 class="text-center"><i class="fa-solid fa-user-ninja"></i> Costume {vote.costume}/10</h4>
        <form method="POST" action="?/vote" use:enhance bind:this={costForm}>
            <input name="countrySlug" type="hidden" value={country.slug}>
            <input name="cat" type="hidden" value="costume">
            <input name="userId" type="hidden" value={$currentUser.id}>
            <div class="flex flex-row-reverse items-center mx-auto justify-between">
                {#each voteOptions as { key, label }}
                    <input id="costume-star-{label}" class="hidden peer" type="radio" bind:group={vote.costume} value={key} name="score" on:click={() => {
                        costForm.requestSubmit()
                    }}/>
                    <label
                            class="text-yellow-400 ease-in-out duration-300 text-4xl
                            hover:text-yellow-500
                            before:content-['\2606']
                            peer-checked:before:content-['\2605']"
                            for="costume-star-{label}"></label>
                {/each}
            </div>
        </form>
    </div>

    <div class="py-5">
        <h4 class="text-center"><i class="fa-solid fa-burst"></i> Staging and Props {vote.props}/10</h4>
        <form method="POST" action="?/vote" use:enhance bind:this={propForm}>
            <input name="countrySlug" type="hidden" value={country.slug}>
            <input name="cat" type="hidden" value="props">
            <input name="userId" type="hidden" value={$currentUser.id}>
            <div class="flex flex-row-reverse items-center mx-auto justify-between">
                {#each voteOptions as { key, label }}
                    <input id="props-star-{label}" class="hidden peer" type="radio" bind:group={vote.props} value={key} name="score" on:click={() => {
                        propForm.requestSubmit()
                    }}/>
                    <label
                            class="text-yellow-400 ease-in-out duration-300 text-4xl
                            hover:text-yellow-500
                            before:content-['\2606']
                            peer-checked:before:content-['\2605']"
                            for="props-star-{label}"></label>
                {/each}
            </div>
        </form>
    </div>
</div>