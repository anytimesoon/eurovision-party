<script>
    import {commentQueue} from "$lib/stores/commentQueue.store";
    import {userStore} from "$lib/stores/user.store";
    import ChatBubble from "$lib/components/chat/ChatBubble.svelte";
    import { fly } from 'svelte/transition';
    import { cubicInOut } from 'svelte/easing';
</script>

<div class="pt-3">
    {#each $commentQueue as chatMessage}
        <div transition:fly={{ duration: 150, y: 500, opacity: 0.5, easing: cubicInOut}}>
            <ChatBubble comment={chatMessage.body}
                        user={$userStore[chatMessage.body.userId]}
                        isCurrentUser={true}/>
        </div>
    {/each}
</div>
