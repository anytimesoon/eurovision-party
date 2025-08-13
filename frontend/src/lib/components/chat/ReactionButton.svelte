<script lang="ts">
    import {postWithError} from "$lib/utils/genericFetch";
    import {CommentReactionModel} from "$lib/models/classes/commentReaction.model";
    import {commentEP} from "$lib/models/enums/endpoints.enum";
    import type {CommentModel} from "$lib/models/classes/comment.model";
    import {currentUser} from "$lib/stores/user.store";
    import {errorStore} from "$lib/stores/error.store";
    import type {SvelteSet} from "svelte/reactivity";
    import {emojiPickerState} from "$lib/stores/emojiPickerState.store";
    import {onMount} from "svelte";
    import {addClassList} from "$lib/utils/addClassList";
    import {MagicClassEnun} from "$lib/models/enums/magicClass.enun";

    interface Props {
        comment: CommentModel
        reaction: [string, SvelteSet<string>]
        isMenuButton: boolean
    }

    const {comment, reaction, isMenuButton}: Props = $props();
    const [emoji, users] = reaction
    let button: HTMLDivElement

    onMount(() => {
        addClassList(button, MagicClassEnun.EMOJI)
    })

    const handleClick = (e: MouseEvent, emoji: string) => {
        e.preventDefault();
        if (isMenuButton) {
            emojiPickerState.open(comment)
        } else {
            handleAddOrRemoveEmoji(emoji)
        }
    }

    const handleAddOrRemoveEmoji = (emoji:string) => {
        const action = comment.addOrRemoveReaction($currentUser.id, emoji)
        postWithError<CommentReactionModel, CommentReactionModel>(commentEP.REACTIONS, new CommentReactionModel(
            action,
            $currentUser.id,
            comment.id,
            emoji,
        )).then(resp => {
            if (resp.error) {
                $errorStore = resp.error
                comment.addOrRemoveReaction($currentUser.id, emoji)
            }
        })
    }

    let isSelected = $derived(users.has($currentUser.id) ? "border-secondary-light" : "border-canvas-secondary")
</script>

<div class="border-2 border-canvas-secondary z-10 max-h-min rounded-xl mx-1 bg-canvas-secondary" bind:this={button}>
    <button
            class=" px-2 py-0.5
                    bg-canvas-primary
                    rounded-xl
                    items-center
                    justify-center
                    flex
                    flex-auto
                    border
                    {isSelected}"
            onclick={(e) => handleClick(e, emoji)}
    >
        <span class="text-lg">{emoji}</span>

        {#if users.size > 1}
            <span class="text-xs pl-1">{users.size > 1 ? users.size : ""}</span>
        {/if}
    </button>
</div>