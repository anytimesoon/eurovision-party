<script lang="ts">
    import {commentStore} from "$lib/stores/comment.store";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import {ChatMessageModel} from "$lib/models/classes/chatMessage.model.js";

    export let data;
    let socket = data.socket

    socket.onmessage = function (event) {
        const split = event.data.split("\n")
        split.map((c:string)=>{
            const chatMessage = JSON.parse(c)
            switch (chatMessage.category) {
                case "comment":
                    let comment:CommentModel = chatMessage.body
                    comment.createdAt = new Date(chatMessage.body.createdAt)
                    commentStore.update(comments => {
                        return [...comments, comment]
                    });
                    break
                default:
                    console.log(chatMessage.category)
                    console.log("bad message: " + c)
            }
        })


    };

    function sendMsg() {
        // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
        let input = document.getElementById("msg")! as HTMLInputElement;

        const comment = new ChatMessageModel<CommentModel>("comment", new CommentModel(input.value, $currentUser.id))
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
            <p>{$userStore[comment.userId].name || ""} {comment.createdAt.getHours()}:{comment.createdAt.getMinutes()}</p>
            <p>{comment.text}</p>
        </div>
    {/each}
</div>

<input type="text" name="msg" id="msg" on:keyup={sendMsgWithKeyboard}/>
<input type="button" value="Send" on:click={sendMsg}/>