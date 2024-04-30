<script lang="ts">
    import {quintInOut} from "svelte/easing";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import Send from "svelte-material-icons/Send.svelte";
    import CloseCircleOutline from "svelte-material-icons/CloseCircleOutline.svelte";
    import {ChatMessageModel} from "$lib/models/classes/chatMessage.model";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
    import {commentQueue} from "$lib/stores/commentQueue.store";
    import {replyComment} from "$lib/stores/replyComment.store";
    import { scale } from 'svelte/transition';

    function sendMsg() {
        const input = document.getElementById("msg")! as HTMLInputElement;

        if(input.value === "" || input.value === null) {
            return
        }

        const comment = new ChatMessageModel<CommentModel>(
            chatMsgCat.COMMENT,
                new CommentModel(
                    input.value,
                    $currentUser.id,
                    replyCommentOrNull(),
                    null,
                    true
                )
        )

        commentQueue.addComment(comment)
        replyComment.close()
        input.value = ""
        input.style.height = "40px"
        input.focus()
    }

    function replyCommentOrNull(){
        return $replyComment.createdAt != null ? $replyComment : null
    }

    function sendMsgWithKeyboard(e:KeyboardEvent){
        const input = e.target as HTMLInputElement
        input.style.height = "1px"
        input.style.height = (4+input.scrollHeight)+"px"

        if(e.key == "Enter"){
            sendMsg()
        }
    }
</script>
<div>
    {#if $replyComment.text !== undefined}
        <div transition:scale={{ duration: 500, opacity: 0.5, easing: quintInOut }} class="bg-canvas-secondary p-2 mb-1 rounded text-typography-main text-sm relative">
            <button class="bg-transparent absolute top-1 right-1"  on:click={() => replyComment.close()}>
                <CloseCircleOutline />
            </button>

            {#if $replyComment.userId !== undefined}
                <div class="pb-2">
                    {$userStore[$replyComment.userId].name}
                </div>
            {/if}
            {$replyComment.text}
        </div>
    {/if}
    <div class="flex">
        <textarea class="h-10 text-sm overflow-hidden" name="msg" id="msg" on:keyup={e => sendMsgWithKeyboard(e)}></textarea>
        <div class="flex flex-col-reverse ml-2">
            <button on:click={sendMsg}>
                <Send size="1.4em"/>
            </button>
        </div>
    </div>
</div>