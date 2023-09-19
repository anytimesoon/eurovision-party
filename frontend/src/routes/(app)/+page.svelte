<script lang="ts">
    import {commentStore} from "$lib/stores/comment.store";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import {ChatMessageModel} from "$lib/models/classes/chatMessage.model.js";
    import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
    import {socketStore} from "$lib/stores/socket.store";
    import ChatMessage from "$lib/components/chat/ChatMessage.svelte";

    function sendMsg() {
        const input = document.getElementById("msg")! as HTMLInputElement;

        if(input.value === "" || input.value === null) {
            return
        }

        const comment = new ChatMessageModel<CommentModel>(chatMsgCat.COMMENT, new CommentModel(input.value, $currentUser.id))
        $socketStore.send(JSON.stringify(comment))
        input.value = ""
    }

    function sendMsgWithKeyboard(e:KeyboardEvent){
        if(e.key == "Enter"){
            sendMsg()
        }
    }

</script>

<div class="h-[calc(100vh-6rem)] flex flex-col">
    <div id="chat-box" class="border-2 flex flex-col-reverse flex-auto border-purple-400 p-4 overflow-auto rounded mb-3">
        {#each $commentStore as comment}
            <ChatMessage comment={comment} user={$userStore[comment.userId]}/>
        {/each}
    </div>

    <div class="flex">
        <textarea class="w-full border-2 border-purple-400 rounded text-sm p-3" name="msg" id="msg" on:keyup={sendMsgWithKeyboard}></textarea>
        <input class="ml-4 p-4 bg-purple-500 text-white rounded" type="button" value="Send" on:click={sendMsg}/>
    </div>
</div>




