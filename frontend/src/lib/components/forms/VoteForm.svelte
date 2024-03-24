<script lang="ts">
    import type {VoteModel} from "$lib/models/classes/vote.model";
    import {onMount} from "svelte";
    import {currentUser} from "$lib/stores/user.store";
    import {voteOptions} from "$lib/models/classes/voteOptions.model";
    import Spinner from "$lib/components/Spinner.svelte";
    import {VoteSingleModel} from "$lib/models/classes/voteSingle.model";
    import {voteSvelteEP} from "$lib/models/enums/endpoints.enum";

    export let vote:VoteModel
    export let catName:string
    export let countrySlug:string
    const localOptions = voteOptions.slice()
    let isFetching:boolean = false
    let cat:string
    let score:number
    let icon:string
    let tempScore:number
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

    function setScore(voteResult:VoteModel) {
        switch (catName) {
            case "Song":
                score = voteResult.song
                break
            case "Performance":
                score = voteResult.performance
                break
            case "Costumes":
                score = voteResult.costume
                break
            case "Staging and Props":
                score = voteResult.props
                break
        }
    }

    $: shouldRotate = isFetching ? "animate-spin" : ""

    const submitForm = async () => {
        isFetching = true
        let singleVote = new VoteSingleModel($currentUser.id, countrySlug, cat, tempScore)
        let voteResponse = await fetch(voteSvelteEP.UPDATE, {
            method: "PUT",
            body: JSON.stringify(singleVote),
            credentials: "same-origin",
            headers: {
                "Content-Type": "application/json"
            }
        }).then((res) => res.json())

        if (voteResponse.error != "") {
            alert(voteResponse.error)
        } else {
            console.log("here")
            setScore(voteResponse.body)
        }

        isFetching = false
    }

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

    <form bind:this={form} on:submit|preventDefault={() => submitForm()}>
        <input name="countrySlug" type="hidden" value={countrySlug}>
        <input name="cat" type="hidden" value={cat}>
        <input name="userId" type="hidden" value={$currentUser.id}>
        <div class="flex flex-row-reverse items-center mx-auto justify-between">
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
                        class="text-4xl
                               {shouldRotate}
                               cursor-pointer
                               before:content-['\2606']
                               peer-checked:before:content-['\2B50']"
                        for="{cat}-star-{label}"></label>
            {/each}
        </div>
    </form>
</div>