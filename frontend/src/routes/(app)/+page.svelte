<script lang="ts">
    import {commentStore} from "$lib/stores/comment.store";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {currentUser} from "$lib/stores/user.store";
    import {ChatMessageModel} from "$lib/models/classes/chatMessage.model.js";
    import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
    import type {UserModel} from "$lib/models/classes/user.model";
    import {staticEP} from "$lib/models/enums/endpoints.enum.js";

    export let data;
    let socket = data.socket
    let users:Map<string, UserModel> = data.users

    function sendMsg() {
        const input = document.getElementById("msg")! as HTMLInputElement;

        if(input.value === "" || input.value === null) {
            return
        }

        const comment = new ChatMessageModel<CommentModel>(chatMsgCat.COMMENT, new CommentModel(input.value, $currentUser.id))
        socket.send(JSON.stringify(comment))
        input.value = ""
    }

    function sendMsgWithKeyboard(e:KeyboardEvent){
        if(e.key == "Enter"){
            sendMsg()
        }
    }

</script>

<div class="flex flex-col">
    <div id="chat-box" class="border-2 flex-auto border-purple-400 h-0 p-4 overflow-auto rounded mb-3">
        {#each $commentStore as comment}
            {#if $currentUser.id === comment.userId}
                <div class="flex w-full mt-2 space-x-3 max-w-xs ml-auto justify-end">
                    <div>
                        <p class="text-sm text-right">{users[comment.userId].name || ""}</p>
                        <div class="bg-purple-800 text-white p-3 rounded-l-lg rounded-br-lg">

                            <p class="text-sm">{comment.text}</p>
                        </div>
                        <span class="text-xs text-gray-500 leading-none">{comment.createdAt.getHours()}:{comment.createdAt.getMinutes() < 10 ? "0" + comment.createdAt.getMinutes() : comment.createdAt.getMinutes()}</span>
                    </div>
                    <img class="flex-shrink-0 h-10 w-10 rounded-full" src={staticEP.IMG + users[comment.userId].icon} alt={users[comment.userId].name + "'s avatar"}>
                </div>
            {:else}
                <div class="flex w-full mt-2 space-x-3 max-w-xs">
                    <img class="flex-shrink-0 h-10 w-10 rounded-full" src={staticEP.IMG + users[comment.userId].icon} alt={users[comment.userId].name + "'s avatar"}>
                    <div>
                        <p class="text-sm">{users[comment.userId].name || ""}</p>
                        <div class="bg-gray-300 p-3 rounded-r-lg rounded-bl-lg">

                            <p class="text-sm">{comment.text}</p>
                        </div>
                        <span class="text-xs text-gray-500 leading-none">{comment.createdAt.getHours()}:{comment.createdAt.getMinutes() < 10 ? "0" + comment.createdAt.getMinutes() : comment.createdAt.getMinutes()}</span>
                    </div>
                </div>
            {/if}


        {/each}
    </div>

    <div class="flex">
        <textarea class="flex w-full border-2 border-purple-400 rounded px-3 text-sm p-4" name="msg" id="msg" on:keyup={sendMsgWithKeyboard}></textarea>
        <input class="m-4 p-4 bg-purple-500 text-white rounded" type="button" value="Send" on:click={sendMsg}/>
    </div>
</div>




