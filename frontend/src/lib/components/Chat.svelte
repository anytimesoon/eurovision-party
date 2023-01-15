<script lang="ts">
    import { onMount } from "svelte";
    import { connectToChat } from "$lib/stores/comment.store";
    import {commentStore} from "$lib/stores/comment.store.js";
    import {userStore} from "$lib/stores/user.store.js";
    let socket:WebSocket;
    
    onMount(async() => {
        socket = connectToChat()
    });

    function sendMsg() {
        // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
        let input = document.getElementById("msg")! as HTMLInputElement;
        if (!socket) {
            console.log("Your connection has been lost. Try reconnecting.");
        }
        if (!input.value) {
            console.log("Something went very, very wrong")
        }
        
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
            <p>{ $userStore[comment.userId].name } at {comment.createdAt}</p>
            <p>{comment.text}</p>
        </div>
    {/each}
</div>

<input type="text" name="msg" id="msg" on:keyup={sendMsgWithKeyboard}/>
<input type="button" value="Send" on:click={sendMsg}/>