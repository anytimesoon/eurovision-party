<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";
    import ReactionButton from "$lib/components/chat/ReactionButton.svelte";
    import {SvelteSet} from "svelte/reactivity";

    interface Props {
        comment: CommentModel
        isCurrentUser: boolean
    }

    let {comment, isCurrentUser}:Props = $props();
    let reactions = comment.reactions;



    let leftPadding = $derived(isCurrentUser ? "flex-row-reverse" : "ml-[2.7em]")
</script>

<div class="{leftPadding} flex -mt-1.5 overflow-auto w-[75%)">
    <ReactionButton comment={comment} isMenuButton={true} reaction={['+', new SvelteSet()]}/>

    {#each reactions.entries() as reaction}
        <ReactionButton comment={comment} reaction={reaction} isMenuButton={false}/>
    {/each}
</div>
