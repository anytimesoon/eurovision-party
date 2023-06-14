<script lang="ts">
    import {commentStore} from "$lib/stores/comment.store";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import {ChatMessageModel} from "$lib/models/classes/chatMessage.model.js";
    import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
    import type {UserModel} from "$lib/models/classes/user.model";

    export let data;
    let socket = data.socket
    let users:Map<string, UserModel> = data.users

    socket.onmessage = function (event) {
        const split = event.data.split("\n")
        split.map((c:string)=>{
            const chatMessage = JSON.parse(c)
            switch (chatMessage.category) {
                case chatMsgCat.COMMENT:
                    let comment:CommentModel = chatMessage.body
                    comment.createdAt = new Date(chatMessage.body.createdAt)
                    commentStore.update(comments => {
                        return [...comments, comment]
                    });
                    break
                default:
                    console.log("bad message: " + c)
            }
        })


    };

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
            <p>{users[comment.userId].name || ""} {comment.createdAt.getHours()}:{comment.createdAt.getMinutes()}</p>
            <p>{comment.text}</p>
        </div>
    {/each}
</div>

<input type="text" name="msg" id="msg" on:keyup={sendMsgWithKeyboard}/>
<input type="button" value="Send" on:click={sendMsg}/>