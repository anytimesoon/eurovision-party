<script lang="ts">
    import {commentStore} from "$lib/stores/comment.store";
    import {userStore} from "$lib/stores/user.store";
    import {socketStore} from "$lib/stores/socket.store";

    function sendMsg() {
        // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
        let input = document.getElementById("msg")! as HTMLInputElement;

        socketStore.send(input.value);
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
            <p>{ $userStore[comment.userId].name } at {comment.createdAt}</p>
            <p>{comment.text}</p>
        </div>
    {/each}
</div>

<input type="text" name="msg" id="msg" on:keyup={sendMsgWithKeyboard}/>
<input type="button" value="Send" on:click={sendMsg}/>