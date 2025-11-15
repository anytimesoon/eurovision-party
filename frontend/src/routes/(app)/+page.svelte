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
    import ImageLoader from "$lib/components/images/ImageLoader.svelte";
    import EmojiPicker from "$lib/components/chat/EmojiPicker.svelte";
    import {flip} from "svelte/animate";
    import {fade} from "svelte/transition";

    let openModal:VoidFunction = $state()
    let closeModal:VoidFunction = $state()
    let userWithActiveAvatar:UserModel = $state(UserModel.empty())
    let showCommentHistory = $state(false)

    const openAvatarModal = (user:UserModel) => {
        userWithActiveAvatar = user
        openModal()
    }
</script>

<ConnectionSpinner/>

<Modal bind:openModal={openModal} bind:closeModal={closeModal} isEasilyClosable={true}>
    {#if userWithActiveAvatar && userWithActiveAvatar.icon !== undefined}
        <ImageLoader src={staticEP.AVATAR_IMG + userWithActiveAvatar.icon} alt={userWithActiveAvatar.name + "'s avatar"} customClasses="mx-auto"/>
    {/if}
</Modal>

<div class="flex flex-col h-full">
    <div id="chat-box" class="scroll-smooth border-2 flex flex-col-reverse flex-auto bg-canvas-secondary border-secondary py-4 px-1 overflow-y-auto overflow-x-hidden rounded mb-3">
        {#if $commentQueue && $commentQueue.length > 0}
            <CommentQueue />
        {/if}

        {#each $recentComments as comment (comment.id)}
            <div animate:flip={{duration: 200}} in:fade={{delay: 100, duration: 100}}>
                <ChatBubble comment={comment}
                                user={$userStore.get(comment.userId)}
                                isCurrentUser={($currentUser.id === comment.userId)}
                                openAvatarModal={openAvatarModal}/>
            </div>
        {/each}

        {#if $olderComments.length > 0 && !showCommentHistory}
            <div class="p-3 mx-auto">
                <button onclick={() => showCommentHistory = true} class="drop-shadow-lg">
                    <span class="flex">
                        <Reload size="1.2em" class="pt-1 pr-0.5"/> Load older comments <CursorDefaultClick size="1.2em" class="pt-1 pl-0.5"/>
                    </span>
                </button>
            </div>
        {:else}
            {#each $olderComments as comment (comment.id)}
                <div animate:flip={{duration: 200}} in:fade>
                    <ChatBubble comment={comment}
                                user={$userStore.get(comment.userId)}
                                isCurrentUser={($currentUser.id === comment.userId)}
                                openAvatarModal={openAvatarModal}/>
                </div>
            {/each}
        {/if}
    </div>

    <ChatInputForm />

    <EmojiPicker />
</div>




