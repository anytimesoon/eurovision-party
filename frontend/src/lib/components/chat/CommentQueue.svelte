<script>
    import {commentQueue} from "$lib/stores/commentQueue.store";
    import {userStore} from "$lib/stores/user.store";
    import ChatBubble from "$lib/components/chat/ChatBubble.svelte";
    import {flip} from "svelte/animate";
    import {fade} from "svelte/transition";

    let isHidden = $state(true)

    $effect(() => {
        if ($commentQueue.length > 0) {
            setTimeout(() => {
                isHidden = false
            }, 1000)
        } else {
            isHidden = true
        }
    })
</script>

<div class="pt-3">
    {#each $commentQueue as chatMessage (chatMessage.body.id)}
        <div animate:flip={{duration: 200}} in:fade={{delay: 100, duration: 200}}>
            {#if !isHidden}
                <ChatBubble comment={chatMessage.body}
                            user={$userStore.get(chatMessage.body.userId)}
                            isCurrentUser={true}/>
            {/if}
        </div>
    {/each}
</div>
