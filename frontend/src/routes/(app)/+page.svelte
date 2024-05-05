<script lang="ts">
    import {commentStore} from "$lib/stores/comment.store";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import ChatBubble from "$lib/components/chat/ChatBubble.svelte";
    import Modal from "$lib/components/Modal.svelte";
    import {UserModel} from "$lib/models/classes/user.model";
    import {staticSvelteEP} from "$lib/models/enums/endpoints.enum";
    import {commentQueue} from "$lib/stores/commentQueue.store";
    import CommentQueue from "$lib/components/chat/CommentQueue.svelte";
    import ConnectionSpinner from "$lib/components/chat/ConnectionSpinner.svelte";
    import ChatInputForm from "$lib/components/forms/ChatInputForm.svelte";

    let openModal:VoidFunction
    let closeModal:VoidFunction
    let userWithActiveAvatar:UserModel = new UserModel()

    const openAvatarModal = (user:UserModel) => {
        userWithActiveAvatar = user
        openModal()
    }
</script>

<ConnectionSpinner/>

<Modal bind:openModal={openModal} bind:closeModal={closeModal} isEasilyClosable={true}>
    {#if userWithActiveAvatar && userWithActiveAvatar.icon !== undefined}
        <img class="mx-auto" src={staticSvelteEP.AVATAR_IMG + userWithActiveAvatar.icon} alt={userWithActiveAvatar.name + "'s avatar"}/>
    {/if}
</Modal>

<div class="flex flex-col h-full">
    <div id="chat-box" class="scroll-smooth border-2 flex flex-col-reverse flex-auto bg-canvas-secondary border-secondary p-4 overflow-y-auto overflow-x-hidden rounded mb-3">
        {#if $commentQueue && $commentQueue.length > 0}
            <CommentQueue />
        {/if}

        {#each $commentStore as comment}
            <ChatBubble comment={comment}
                        user={$userStore[comment.userId]}
                        isCurrentUser={($currentUser.id === comment.userId)}
                        openAvatarModal={openAvatarModal}/>
        {/each}
    </div>

    <ChatInputForm />
</div>




