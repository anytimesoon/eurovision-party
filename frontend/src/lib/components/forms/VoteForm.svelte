<script lang="ts">
    import {VoteModel} from "$lib/models/classes/vote.model";
    import {onMount} from "svelte";
    import {currentUser} from "$lib/stores/user.store";
    import {voteOptions} from "$lib/models/classes/voteOptions.model";
    import Spinner from "$lib/components/Spinner.svelte";
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import {VoteSingleModel} from "$lib/models/classes/voteSingle.model";
    import {voteEP} from "$lib/models/enums/endpoints.enum";
    import {put} from "$lib/utils/genericFetch";

    interface Props {
        vote: VoteModel;
        catName: string;
        countrySlug: string;
        updateVote: (vote: VoteModel) => void;
    }

    let {
        vote,
        catName,
        countrySlug,
        updateVote
    }: Props = $props();
    const localOptions = voteOptions.slice()
    let isFetching:boolean = false
    let cat:string = $state()
    let score:number = $state()
    let icon:string = $state()
    let formState = $state(formButtonState.ENABLED)

    const submit = async (newValue: number) => {
        formState = formButtonState.SENDING

        const newVote = await put<VoteModel, VoteSingleModel>(
            voteEP.UPDATE,
            new VoteSingleModel(
                $currentUser.id,
                countrySlug,
                cat,
                newValue,
            )
        ).then(res => VoteModel.deserialize(res))
        updateVote(newVote)
        formState = formButtonState.ENABLED
    }

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

    $effect(() => {
        if(vote) {
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
    });

    let shouldRotate = $derived(formState == formButtonState.SENDING ? "animate-spin" : "")
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

    <form>
        <div class="flex flex-row-reverse items-center mx-auto justify-between w-full">
            {#each localOptions as { key, label }}
                <input id="{cat}-star-{label}"
                       class="hidden peer"
                       type="radio"
                       bind:group={score}
                       value={key}
                       name="score"
                       onclick={() => {
                           submit(key)
                       }}/>
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