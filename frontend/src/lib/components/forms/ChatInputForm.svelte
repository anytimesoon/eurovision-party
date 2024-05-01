<script lang="ts">
    import {currentUser, userStore} from "$lib/stores/user.store";
    import Send from "svelte-material-icons/Send.svelte";
    import Image from "svelte-material-icons/Image.svelte";
    import {ChatMessageModel} from "$lib/models/classes/chatMessage.model";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
    import {commentQueue} from "$lib/stores/commentQueue.store";
    import {replyComment} from "$lib/stores/replyComment.store";
    import ReplyComment from "$lib/components/chat/ReplyComment.svelte";

    let textArea:HTMLTextAreaElement
    let message:string

    function sendMsg() {
        message.trim()
        if(message === "" || message.length === 0) {
            resetTextArea()
            return
        }

        const comment = new ChatMessageModel<CommentModel>(
            chatMsgCat.COMMENT,
                new CommentModel(
                    message,
                    $currentUser.id,
                    $replyComment.createdAt != null ? $replyComment : null,
                    null,
                    true
                )
        )

        commentQueue.addComment(comment)
        resetTextArea()
    }

    function resetTextArea() {
        replyComment.close()
        message = ""
        textArea.style.height = "1.25rem"
        textArea.focus()
    }

    function sendOrResize(e:KeyboardEvent){
        textArea.style.height = "1px"
        textArea.style.height = `${textArea.scrollHeight}px`

        if(e.key == "Enter"){
            message = message.slice(0, -1)
            sendMsg()
        }
    }

    $: if($replyComment) {
        if($replyComment.text !== undefined && textArea !== undefined) {
            textArea.focus()
        }
    }
</script>

<div>
    <div class="flex">
        <div class="flex flex-col-reverse">
            <button class="rounded-full py-3">
                <Image size="1.4em"/>
            </button>
        </div>
        <div class="flex-1
                    mx-3
                    border-solid
                    bg-canvas-secondary
                    border-2
                    border-gray-400
                    p-2
                    rounded-lg
                    shadow-sm">

            <ReplyComment />

            <textarea class="text-sm
                            h-5
                            p-0
                            overflow-hidden
                            border-0
                            focus:border-0
                            focus:outline-0"
                      name="msg"
                      bind:this={textArea}
                      bind:value={message}
                      on:keyup={e => sendOrResize(e)}></textarea>
        </div>

        <div class="flex flex-col-reverse">
            <button on:click={sendMsg} class="rounded-full py-3">
                <Send size="1.4em"/>
            </button>
        </div>
    </div>
</div>