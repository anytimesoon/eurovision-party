<script lang="ts">
    import type {VoteModel} from "$lib/models/classes/vote.model";
    import {onMount} from "svelte";
    import {currentUser} from "$lib/stores/user.store";
    import {voteOptions} from "$lib/models/classes/voteOptions.model";
    import { enhance } from '$app/forms';

    export let vote:VoteModel
    export let catName:string
    export let countrySlug:string
    const localOptions = voteOptions.slice()
    let cat:string
    let score:number
    let icon:string
    let form:HTMLFormElement

    onMount(() => {
        switch (catName) {
            case "Song":
                score = vote.song
                icon = "fa-music"
                cat = "song"
                break
            case "Performance":
                score = vote.performance
                icon = "fa-masks-theater"
                cat = "performance"
                break
            case "Costumes":
                score = vote.costume
                icon = "fa-user-ninja"
                cat = "costume"
                break
            case "Staging and Props":
                score = vote.props
                icon = "fa-burst"
                cat = "props"
                break
        }
    })

    $: if(vote) {
        switch (catName) {
            case "Song":
                score = vote.song
                break
            case "Performance":
                score = vote.performance
                break
            case "Costumes":
                score = vote.costume
                break
            case "Staging and Props":
                score = vote.props
                break
        }
    }

</script>

<div class="py-5">
    <h3 class="text-center"><i class="fa-solid {icon} text-primary"></i> {catName}</h3>
    <p class="text-center text-typography-grey text-sm">{score}/10</p>
    <form method="POST" action="?/vote" use:enhance bind:this={form}>
        <input name="countrySlug" type="hidden" value={countrySlug}>
        <input name="cat" type="hidden" value={cat}>
        <input name="userId" type="hidden" value={$currentUser.id}>
        <div class="flex flex-row-reverse items-center mx-auto justify-between">
            {#each localOptions.reverse() as { key, label }}
                <input id="{cat}-star-{label}" class="hidden peer" type="radio" bind:group={score} value={key} name="score" on:click={() => {
                        form.requestSubmit()
                    }}/>
                <label
                        class="text-vote-star
                               text-4xl
                               before:content-['\2606']
                               peer-checked:before:content-['\2605']"
                        for="{cat}-star-{label}"></label>
            {/each}
        </div>
    </form>
</div>