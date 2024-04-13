<script>
    import {commentQueue} from "$lib/stores/commentQueue.store";
    import {userStore} from "$lib/stores/user.store";
    import ChatBubble from "$lib/components/chat/ChatBubble.svelte";
</script>

<div class="pt-3">
    {#each $commentQueue as chatMessage}
            <ChatBubble comment={chatMessage.body}
                        user={$userStore[chatMessage.body.userId]}
                        isCurrentUser={true}/>
    {/each}
</div>
