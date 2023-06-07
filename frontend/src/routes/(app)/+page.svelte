<script lang="ts">
    import {commentStore} from "$lib/stores/comment.store";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {userStore} from "$lib/stores/user.store";

    export let data;
    let socket = data.socket

    socket.onmessage = function (event) {
        const split = event.data.split("\n")
        const newComments:CommentModel[] = split.map((c:string)=>{
            return new CommentModel().deserialize(c)
        })

        commentStore.update(comments => {
            return [...comments, ...newComments]
        });
    };

    function sendMsg() {
        // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
        let input = document.getElementById("msg")! as HTMLInputElement;

        socket.send(input.value);
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