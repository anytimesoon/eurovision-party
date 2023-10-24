<script lang="ts">
    import type {VoteModel} from "$lib/models/classes/vote.model";
    import {onMount} from "svelte";
    import {currentUser} from "$lib/stores/user.store";
    import {voteOptions} from "$lib/models/classes/voteOptions.model";
    import {voteSvelteEP} from "$lib/models/enums/endpoints.enum";
    import type {VoteSingleModel} from "$lib/models/classes/voteSingle.model";
    import Spinner from "$lib/components/Spinner.svelte";

    export let vote:VoteModel
    export let catName:string
    export let countrySlug:string
    const localOptions = voteOptions.slice()
    let cat:string
    let score:number
    let icon:string
    let voteForm:HTMLFormElement

    onMount(() => {
        updateScore()
    })

    function updateScore() {
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
    }

    async function handleSubmit(e:SubmitEvent) {
        const form = e.target as HTMLFormElement
        const fd:FormData = new FormData(form)

        let res = await fetch(form.action, {
            method: form.method,
            body: fd
        })

        let voteRes:VoteSingleModel = await res.json()

        switch (voteRes.cat) {
            case "Song":
                console.log(voteRes.score)
                score = voteRes.score
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

    $: if(vote) {
        updateScore()
    }
</script>

<div class="py-5">
    <h3 class="text-center"><i class="fa-solid {icon} text-primary"></i> {catName}</h3>
    <div>
        {#if score === undefined}
            <Spinner size={4}/>
        {:else}
            <p class="text-center text-typography-grey text-sm">{score}/10</p>
        {/if}
    </div>
    <form method="POST" action="{voteSvelteEP.UPDATE}" on:submit|preventDefault={handleSubmit} bind:this={voteForm}>
        <input name="countrySlug" type="hidden" value={countrySlug}>
        <input name="cat" type="hidden" value={cat}>
        <input name="userId" type="hidden" value={$currentUser.id}>
        <div class="flex flex-row-reverse items-center mx-auto justify-between">
            {#each localOptions.reverse() as { key, label }}
                <input id="{cat}-star-{label}" class="hidden peer" type="radio" bind:group={score} value={key} name="score" on:click={() => {
                        voteForm.requestSubmit()
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