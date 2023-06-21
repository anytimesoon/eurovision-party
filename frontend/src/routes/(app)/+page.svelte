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
        let input = document.getElementById("msg")! as HTMLInputElement;

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

<div>
    {#each $commentStore as comment}
        <div>

            <p>
                <span>
                    <img src={staticEP.IMG + users[comment.userId].icon} style="width:33px" alt={users[comment.userId].name + "'s avatar"}>
                    {users[comment.userId].name || ""}
                    {comment.createdAt.getHours()}:{comment.createdAt.getMinutes() < 10 ? "0" + comment.createdAt.getMinutes() : comment.createdAt.getMinutes()}
                </span>
            </p>
            <p>{comment.text}</p>
        </div>
    {/each}
</div>

<input type="text" name="msg" id="msg" on:keyup={sendMsgWithKeyboard}/>
<input type="button" value="Send" on:click={sendMsg}/>