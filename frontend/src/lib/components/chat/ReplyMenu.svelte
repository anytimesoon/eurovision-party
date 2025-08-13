<script lang="ts">
    import Reply from "svelte-material-icons/Reply.svelte";
    import EmoticonHappyOutline from "svelte-material-icons/EmoticonHappyOutline.svelte";
    import { emojiPickerState } from "$lib/stores/emojiPickerState.store";
    import type {CommentModel} from "$lib/models/classes/comment.model";
    import {MagicClassEnun} from "$lib/models/enums/magicClass.enun";
    import {addClassList} from "$lib/utils/addClassList";
    import {onMount} from "svelte";

    interface Props {
        replyButtonHandler: Function
        parentComment: CommentModel
    }

    let { replyButtonHandler, parentComment }: Props = $props()
    let emojiPickerButton: HTMLButtonElement;

    onMount(() => {
        addClassList(emojiPickerButton, MagicClassEnun.EMOJI)
    })

    const emojiButtonHandler = () => {
        emojiPickerState.open(parentComment)
    }
</script>

<div class="absolute
    -top-5
    right-1
    bg-primary
    rounded
    border
    border-white
    cursor-pointer
    flex">

    <button onclick={e => replyButtonHandler(e)} class="px-2" >
        <Reply />
    </button>

    <div class="border"></div>

    <button onclick={emojiButtonHandler} class="px-2 {MagicClassEnun.EMOJI}" bind:this={emojiPickerButton}>
        <EmoticonHappyOutline/>
    </button>
</div>
