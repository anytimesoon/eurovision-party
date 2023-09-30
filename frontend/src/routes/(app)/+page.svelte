<script lang="ts">
    import {commentStore} from "$lib/stores/comment.store";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import {ChatMessageModel} from "$lib/models/classes/chatMessage.model.js";
    import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
    import {socketStore} from "$lib/stores/socket.store";
    import ChatBubble from "$lib/components/chat/ChatBubble.svelte";

    function sendMsg() {
        const input = document.getElementById("msg")! as HTMLInputElement;

        if(input.value === "" || input.value === null) {
            return
        }

        const comment = new ChatMessageModel<CommentModel>(chatMsgCat.COMMENT, new CommentModel(input.value, $currentUser.id))
        $socketStore.send(JSON.stringify(comment))
        input.value = ""
        input.style.height = "40px"
    }

    function sendMsgWithKeyboard(e:KeyboardEvent){
        e.target.style.height = "1px";
        e.target.style.height = (24+e.target.scrollHeight)+"px";

        if(e.key == "Enter"){
            console.log("here")
            sendMsg()
        }
    }

</script>

<div class="h-[calc(100vh-6rem)] flex flex-col">
    <div id="chat-box" class="border-2 flex flex-col-reverse flex-auto border-purple-400 p-4 overflow-auto rounded mb-3">
        {#each $commentStore as comment}
            <ChatBubble comment={comment} user={$userStore[comment.userId]}/>
        {/each}
    </div>

    <div class="flex">
        <textarea class="w-full h-10 border-2 border-purple-400 rounded text-sm overflow-hidden" name="msg" id="msg" on:keyup={e => sendMsgWithKeyboard(e)}></textarea>
        <div class="flex flex-col-reverse">
            <button class="ml-1 px-3 h-10 bg-purple-500 text-white rounded" type="button" on:click={sendMsg}><i class="fa-solid fa-angles-right"></i></button>
        </div>
    </div>
</div>




