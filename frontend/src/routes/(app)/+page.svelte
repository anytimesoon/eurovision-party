<script lang="ts">
    import {commentStore} from "$lib/stores/comment.store";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import {ChatMessageModel} from "$lib/models/classes/chatMessage.model.js";
    import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
    import ChatBubble from "$lib/components/chat/ChatBubble.svelte";
    import Modal from "$lib/components/Modal.svelte";
    import {UserModel} from "$lib/models/classes/user.model";
    import {staticSvelteEP} from "$lib/models/enums/endpoints.enum";
    import {commentQueue} from "$lib/stores/commentQueue.store";
    import CommentQueue from "$lib/components/chat/CommentQueue.svelte";
    import ConnectionSpinner from "$lib/components/chat/ConnectionSpinner.svelte";
    import { scale } from 'svelte/transition';
    import {quintInOut} from 'svelte/easing';
    import Send from "svelte-material-icons/Send.svelte";
    import CloseCircleOutline from "svelte-material-icons/CloseCircleOutline.svelte";

    let openModal:VoidFunction
    let closeModal:VoidFunction
    let userWithActiveAvatar:UserModel = new UserModel()
    let replyComment:CommentModel = new CommentModel()

    const openAvatarModal = (user:UserModel) => {
        userWithActiveAvatar = user
        openModal()
    }

    const replyToComment = (comment:CommentModel) => {
        replyComment = comment
    }

    const closeReply = () => {
        replyComment = new CommentModel()
    }

    function sendMsg() {
        const input = document.getElementById("msg")! as HTMLInputElement;

        if(input.value === "" || input.value === null) {
            return
        }

        const comment = new ChatMessageModel<CommentModel>(
            chatMsgCat.COMMENT,
            new CommentModel(
                input.value,
                $currentUser.id,
                replyCommentOrNull(),
                null,
                true
            )
        )

        commentQueue.addComment(comment)
        closeReply()
        input.value = ""
        input.style.height = "40px"
        input.focus()
    }

    function replyCommentOrNull(){
        return replyComment.createdAt != null ? replyComment : null
    }

    function sendMsgWithKeyboard(e:KeyboardEvent){
        const input = e.target as HTMLInputElement
        input.style.height = "1px"
        input.style.height = (4+input.scrollHeight)+"px"

        if(e.key == "Enter"){
            sendMsg()
        }
    }
</script>

<ConnectionSpinner/>

<Modal bind:openModal={openModal} bind:closeModal={closeModal} isEasilyClosable={true}>
    {#if userWithActiveAvatar && userWithActiveAvatar.icon !== undefined}
        <img class="mx-auto" src={staticSvelteEP.IMG + userWithActiveAvatar.icon} alt={userWithActiveAvatar.name + "'s avatar"}/>
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
                        openAvatarModal={openAvatarModal}
                        replyToComment={replyToComment}/>
        {/each}
    </div>

    <div>
        {#if replyComment.text !== undefined}
            <div transition:scale={{ duration: 500, opacity: 0.5, easing: quintInOut }} class="bg-canvas-secondary p-2 mb-1 rounded text-typography-main text-sm relative">
                <button class="bg-transparent absolute top-1 right-1"  on:click={closeReply}>
                    <CloseCircleOutline />
                </button>

                {#if replyComment.userId !== undefined}
                    <div class="pb-2">
                        {$userStore[replyComment.userId].name}
                    </div>
                {/if}
                {replyComment.text}
            </div>
        {/if}
        <div class="flex">
            <textarea class="h-10 text-sm overflow-hidden" name="msg" id="msg" on:keyup={e => sendMsgWithKeyboard(e)}></textarea>
            <div class="flex flex-col-reverse ml-2">
                <button on:click={sendMsg}>
                    <Send size="1.4em"/>
                </button>
            </div>
        </div>
    </div>

</div>




