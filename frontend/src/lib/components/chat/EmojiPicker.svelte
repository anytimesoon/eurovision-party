<script lang="ts">
    import {emojiPickerState} from "$lib/stores/emojiPickerState.store";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import 'emoji-picker-element';
    import {onMount} from "svelte";
    import {Picker} from "emoji-picker-element";
    import {addClassList} from "$lib/utils/addClassList";
    import {MagicClassEnun} from "$lib/models/enums/magicClass.enun";
    import {postWithError} from "$lib/utils/genericFetch";
    import {commentEP} from "$lib/models/enums/endpoints.enum";
    import {CommentReactionModel} from "$lib/models/classes/commentReaction.model";
    import {errorStore} from "$lib/stores/error.store";

    let bottomMenu: HTMLDivElement
    let emojiPickerContainer: HTMLDivElement
    const emojiPicker = new Picker({
        emojiVersion: 14
    })



    onMount(() => {
        emojiPickerContainer.appendChild(emojiPicker)
        addClassList(emojiPickerContainer, MagicClassEnun.EMOJI)
        document.querySelector('emoji-picker')
            .addEventListener('emoji-click', event => {
                const action = $emojiPickerState.comment.addOrRemoveReaction($currentUser.id, event.detail.unicode)
                const comment = $emojiPickerState.comment
                emojiPickerState.close()
                postWithError<CommentReactionModel, CommentReactionModel>(commentEP.REACTIONS, new CommentReactionModel(
                    action,
                    $currentUser.id,
                    comment.id,
                    event.detail.unicode,
                )).then(resp => {
                    if (resp.error) {
                        $errorStore = resp.error
                        comment.addOrRemoveReaction($currentUser.id, event.detail.unicode)
                    }
                })
            });
    })

    $effect(() => {
        if ($emojiPickerState.isVisible) {
            bottomMenu.classList.remove("-bottom-[475px]")
            bottomMenu.classList.add("bottom-0")
        } else {
            bottomMenu.classList.remove("bottom-0")
            bottomMenu.classList.add("-bottom-[475px]")
        }
    })
</script>

<div bind:this={bottomMenu} class="fixed rounded bg-canvas-primary -bottom-[475px] max-w-[600px] mx-auto left-0 right-0 h-[475px] z-30 duration-500 overflow-hidden border border-secondary">

    <div class="flex flex-col h-full">
        <div class="p-3 h-[75px] overflow-hidden relative before:absolute before:bottom-0 before:left-0 before:right-0 before:h-8 before:bg-gradient-to-t before:from-canvas-primary before:to-transparent">
            {$userStore.get($emojiPickerState.comment.userId)?.name || ""}
            <p>{$emojiPickerState.comment.text}</p>
        </div>
        <div class="w-full flex-1" bind:this={emojiPickerContainer}>
        </div>
    </div>
</div>