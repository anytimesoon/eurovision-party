<script lang="ts">
    import {olderComments, recentComments} from "$lib/stores/comment.store";
    import Reload from "svelte-material-icons/Reload.svelte";
    import CursorDefaultClick from "svelte-material-icons/CursorDefaultClick.svelte";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import ChatBubble from "$lib/components/chat/ChatBubble.svelte";
    import Modal from "$lib/components/Modal.svelte";
    import {UserModel} from "$lib/models/classes/user.model";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import {commentQueue} from "$lib/stores/commentQueue.store";
    import CommentQueue from "$lib/components/chat/CommentQueue.svelte";
    import ConnectionSpinner from "$lib/components/chat/ConnectionSpinner.svelte";
    import ChatInputForm from "$lib/components/forms/ChatInputForm.svelte";

    let openModal:VoidFunction
    let closeModal:VoidFunction
    let userWithActiveAvatar:UserModel = UserModel.empty()
    let showCommentHistory = false

    const openAvatarModal = (user:UserModel) => {
        userWithActiveAvatar = user
        openModal()
    }
</script>

<ConnectionSpinner/>

<Modal bind:openModal={openModal} bind:closeModal={closeModal} isEasilyClosable={true}>
    {#if userWithActiveAvatar && userWithActiveAvatar.icon !== undefined}
        <img class="mx-auto" src={staticEP.AVATAR_IMG + userWithActiveAvatar.icon} alt={userWithActiveAvatar.name + "'s avatar"}/>
    {/if}
</Modal>

<div class="flex flex-col h-full">
    <div id="chat-box" class="scroll-smooth border-2 flex flex-col-reverse flex-auto bg-canvas-secondary border-secondary py-4 px-1 overflow-y-auto overflow-x-hidden rounded mb-3">
        {#if $commentQueue && $commentQueue.length > 0}
            <CommentQueue />
        {/if}

        {#each {length: $recentComments.length} as _, index}
            {@const reverseIndex = $recentComments.length - 1 - index}
            {@const comment = $recentComments[reverseIndex]}
            <ChatBubble comment={comment}
                        user={$userStore[comment.userId]}
                        isCurrentUser={($currentUser.id === comment.userId)}
                        openAvatarModal={openAvatarModal}/>
        {/each}

        {#if $olderComments.length > 0 && !showCommentHistory}
            <div class="p-3 mx-auto">
                <button on:click={() => showCommentHistory = true} class="drop-shadow-lg">
                    <span class="flex">
                        <Reload size="1.2em" class="pt-1 pr-0.5"/> Load older comments <CursorDefaultClick size="1.2em" class="pt-1 pl-0.5"/>
                    </span>
                </button>
            </div>
        {:else}
            {#each {length: $olderComments.length} as _, index}
                {@const reverseIndex = $olderComments.length - 1 - index}
                {@const comment = $olderComments[reverseIndex]}
                <ChatBubble comment={comment}
                            user={$userStore[comment.userId]}
                            isCurrentUser={($currentUser.id === comment.userId)}
                            openAvatarModal={openAvatarModal}/>
            {/each}
        {/if}
    </div>

    <ChatInputForm />
</div>




