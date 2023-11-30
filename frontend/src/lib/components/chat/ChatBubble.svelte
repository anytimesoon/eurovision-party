<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import type {UserModel} from "$lib/models/classes/user.model";
    import ChatContent from "$lib/components/chat/ChatContent.svelte";

    export let comment:CommentModel
    export let user:UserModel
    export let isCurrentUser:boolean

    $: userNameStyle = isCurrentUser ? "text-right" : ""
    $: currentUserBubbleContainer = isCurrentUser ? "ml-auto justify-end" : ""
    $: currentUserImage = isCurrentUser ? "order-last ml-3" : "mr-3"
    $: currentUserBubble = isCurrentUser ? "bg-chat-bubble-me" : "bg-chat-bubble-you"
    $: roundedCorners = isCurrentUser ? "rounded-l-md rounded-r-sm" : "rounded-r-md rounded-l-sm"
    $: compactBubble = comment.isCompact ? "mt-1" : "mt-2"
</script>

{#if user.authLvl === authLvl.BOT}
    <div class="text-center mt-2 text-s p-3">
        <p class="text-sm">{comment.text}</p>
    </div>
{:else}


    <div class="flex w-full max-w-[22rem] {currentUserBubbleContainer} {compactBubble}">
        {#if !comment.isCompact}
            <img class="flex-shrink-0 h-10 w-10 rounded-full {currentUserImage}" src={staticEP.IMG + user.icon} alt={user.name + "'s avatar"}>
        {:else}
            <div class="h-10 w-10 {currentUserImage}"></div>
        {/if}

        <div>
            {#if !comment.isCompact}
                <p class="block {userNameStyle}">{user.name}</p>
            {/if}

            <div class="px-3 py-2 {roundedCorners} {currentUserBubble}">
                <ChatContent comment={comment}
                             isCurrentUser={isCurrentUser}/>
            </div>
        </div>



    </div>
{/if}