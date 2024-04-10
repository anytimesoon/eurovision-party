<script lang="ts">
    import type {VoteModel} from "$lib/models/classes/vote.model";
    import {onMount} from "svelte";
    import {currentUser} from "$lib/stores/user.store";
    import {voteOptions} from "$lib/models/classes/voteOptions.model";
    import Spinner from "$lib/components/Spinner.svelte";
    import { enhance } from '$app/forms';
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";

    export let vote:VoteModel
    export let catName:string
    export let countrySlug:string
    const localOptions = voteOptions.slice()
    let isFetching:boolean = false
    let cat:string
    let score:number
    let icon:string
    let tempScore:number
    let formState = formButtonState.ENABLED
    let form:HTMLFormElement

    onMount(() => {
        switch (catName) {
            case "Song":
                score = vote.song
                icon = " 🎤"
                cat = "song"
                break
            case "Performance":
                score = vote.performance
                icon = "💃"
                cat = "performance"
                break
            case "Costumes":
                score = vote.costume
                icon = "👘"
                cat = "costume"
                break
            case "Staging and Props":
                score = vote.props
                icon = " \t🎆"
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

    $: shouldRotate = formState == formButtonState.SENDING ? "animate-spin" : ""
</script>

<div class="py-5">
    <h3 class="text-center">
        <span class="text-4xl">{icon}</span>{catName}
    </h3>
    <p class="text-center text-typography-grey text-sm">
        {#if isFetching}
            <Spinner size={"s"} thickness={"s"} isBlock={false}/>
        {:else}
            {score}
        {/if}/10
    </p>

    <form method="POST" action="?/vote" bind:this={form} use:enhance={() => {
        formState = formButtonState.SENDING

        return async ({ update }) => {
            await update()
            formState = formButtonState.ENABLED
        };
    }} >
        <input name="countrySlug" type="hidden" value={countrySlug}>
        <input name="cat" type="hidden" value={cat}>
        <input name="userId" type="hidden" value={$currentUser.id}>
        <div class="flex flex-row-reverse items-center mx-auto justify-between w-full">
            {#each localOptions as { key, label }}
                <input id="{cat}-star-{label}"
                       class="hidden peer"
                       type="radio"
                       bind:group={score}
                       value={key}
                       name="score"
                       on:click={() => {
                           tempScore = key
                           form.requestSubmit()}}/>
                <label
                        class="text-2xl
                               {shouldRotate}
                               cursor-pointer
                               before:content-['\2606']
                               peer-checked:before:content-['\2B50']"
                        for="{cat}-star-{label}"></label>
            {/each}
        </div>
    </form>
</div>