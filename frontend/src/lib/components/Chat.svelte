<script lang="ts">
    import { onMount } from "svelte";
    import { connectToChat } from "$lib/stores/comment.store";
	import CommentLog from "$lib/components/CommentLog.svelte";
    let socket:any;
    
    onMount(async() => {
        socket = connectToChat()
    });

    function sendMsg() {
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

<CommentLog />
<input type="text" name="msg" id="msg" on:keyup={sendMsgWithKeyboard}/>
<input type="button" value="Send" on:click={sendMsg}/>