<script lang="ts">
    import {commentStore} from "$lib/stores/comment.store";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import {ChatMessageModel} from "$lib/models/classes/chatMessage.model.js";
    import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
    import {socketStore} from "$lib/stores/socket.store";
    import ChatBubble from "$lib/components/chat/ChatBubble.svelte";
    import Spinner from "$lib/components/Spinner.svelte";
    import Modal from "$lib/components/Modal.svelte";
    import {UserModel} from "$lib/models/classes/user.model";
    import {staticEP} from "$lib/models/enums/endpoints.enum";

    let chatButton:HTMLButtonElement
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

        const comment = new ChatMessageModel<CommentModel>(chatMsgCat.COMMENT, new CommentModel(input.value, $currentUser.id, replyComment))
        $socketStore.send(JSON.stringify(comment))
        closeReply()
        input.value = ""
        input.style.height = "40px"
    }

    function sendMsgWithKeyboard(e:KeyboardEvent){
        e.target.style.height = "1px"
        e.target.style.height = (4+e.target.scrollHeight)+"px"

        if(e.key == "Enter"){
            sendMsg()
        }
    }

    $: if ($socketStore) {
        if (chatButton != null) {
            chatButton.disabled = $socketStore.readyState != WebSocket.OPEN
        }
    }
</script>

<Modal bind:openModal={openModal} bind:closeModal={closeModal} isEasilyClosable={true}>
    {#if userWithActiveAvatar && userWithActiveAvatar.icon !== undefined}
        <img class="mx-auto" src={staticEP.IMG + userWithActiveAvatar.icon} alt={userWithActiveAvatar.name + "'s avatar"}/>
    {/if}
</Modal>

<div class="flex flex-col h-full">
    <div id="chat-box" class="border-2 flex flex-col-reverse flex-auto bg-canvas-secondary border-secondary p-4 overflow-y-auto overflow-x-hidden rounded mb-3">
        {#if $socketStore.readyState == WebSocket.CONNECTING}
            <div class="h-screen flex flex-col justify-center">
                <div>
                    <p class="text-center">Connecting to the chat</p>
                    <Spinner />
                </div>
            </div>
        {:else if $socketStore.readyState == WebSocket.OPEN}
            {#each $commentStore as comment}
                <ChatBubble comment={comment}
                            user={$userStore[comment.userId]}
                            isCurrentUser={($currentUser.id === comment.userId)}
                            openAvatarModal={openAvatarModal}
                            replyToComment={replyToComment}/>
            {/each}
        {:else}
            <div class="h-screen flex flex-col justify-center">
                <div class="text-center">
                    <p>Something went very wrong! 😬</p>
                    <p>Please refresh the page</p>
                </div>
            </div>
        {/if}
    </div>

    <div>
        {#if replyComment.text !== undefined}
            <div class="bg-canvas-secondary p-2 mb-1 rounded text-typography-main text-sm relative">
                <button class="bg-transparent absolute top-1 right-1"  on:click={closeReply}>
                    <i class="fa-regular fa-circle-xmark"></i>
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
                <button bind:this={chatButton} on:click={sendMsg}><i class="fa-solid fa-angles-right"></i></button>
            </div>
        </div>
    </div>

</div>




